// Copyright 2018 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"fmt"
	"net/http"
	"time"

	apimetrics "k8s.io/apiserver/pkg/endpoints/metrics"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"

	"sigs.k8s.io/metrics-server/pkg/api"
	"sigs.k8s.io/metrics-server/pkg/scraper"
	"sigs.k8s.io/metrics-server/pkg/storage"
)

type Config struct {
	Apiserver        *genericapiserver.Config
	Rest             *rest.Config
	Kubelet          *scraper.KubeletClientConfig
	MetricResolution time.Duration
	ScrapeTimeout    time.Duration
}

func (c Config) Complete() (*server, error) {
	informer, err := c.informer()
	if err != nil {
		return nil, err
	}
	kubeletClient, err := c.Kubelet.Complete()
	if err != nil {
		return nil, fmt.Errorf("unable to construct a client to connect to the kubelets: %v", err)
	}
	nodes := informer.Core().V1().Nodes()
	scrape := scraper.NewScraper(nodes.Lister(), kubeletClient, c.ScrapeTimeout)

	// Disable default metrics handler and create custom one
	c.Apiserver.EnableMetrics = false
	metricsHandler, err := c.metricsHandler()
	if err != nil {
		return nil, err
	}

	genericServer, err := c.Apiserver.Complete(informer).New("metrics-server", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}
	genericServer.Handler.NonGoRestfulMux.HandleFunc("/metrics", metricsHandler)

	store := storage.NewStorage()
	if err := api.Install(store, informer.Core().V1(), genericServer); err != nil {
		return nil, err
	}
	return NewServer(
		nodes.Informer().HasSynced,
		informer,
		genericServer,
		store,
		scrape,
		c.MetricResolution,
	), nil
}

func (c Config) metricsHandler() (http.HandlerFunc, error) {
	// Create registry for Metrics Server metrics
	registry := metrics.NewKubeRegistry()
	err := RegisterMetrics(registry, c.MetricResolution)
	if err != nil {
		return nil, err
	}
	// Register apiserver metrics in legacy registry
	apimetrics.Register()

	// Return handler that serves metrics from both legacy and Metrics Server registry
	return func(w http.ResponseWriter, req *http.Request) {
		legacyregistry.Handler().ServeHTTP(w, req)
		metrics.HandlerFor(registry, metrics.HandlerOpts{}).ServeHTTP(w, req)
	}, nil
}

func (c Config) informer() (informers.SharedInformerFactory, error) {
	// set up the informers
	kubeClient, err := kubernetes.NewForConfig(c.Rest)
	if err != nil {
		return nil, fmt.Errorf("unable to construct lister client: %v", err)
	}
	// we should never need to resync, since we're not worried about missing events,
	// and resync is actually for regular interval-based reconciliation these days,
	// so set the default resync interval to 0
	return informers.NewSharedInformerFactory(kubeClient, 0), nil
}

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

package scraper

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/mailru/easyjson"

	corev1 "k8s.io/api/core/v1"

	"sigs.k8s.io/metrics-server/pkg/utils"
)

// KubeletInterface knows how to fetch metrics from the Kubelet
type KubeletInterface interface {
	// GetSummary fetches summary metrics from the given Kubelet
	GetSummary(ctx context.Context, node *corev1.Node) (*Summary, error)
}

type kubeletClient struct {
	defaultPort       int
	useNodeStatusPort bool
	client            *http.Client
	scheme            string
	addrResolver      utils.NodeAddressResolver
	buffers           sync.Pool
}

var _ KubeletInterface = (*kubeletClient)(nil)

type ErrNotFound struct {
	endpoint string
}

func (err *ErrNotFound) Error() string {
	return fmt.Sprintf("%q not found", err.endpoint)
}

func (kc *kubeletClient) makeRequestAndGetValue(client *http.Client, req *http.Request, value easyjson.Unmarshaler) error {
	// TODO(directxman12): support validating certs by hostname
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	b := kc.getBuffer()
	defer kc.returnBuffer(b)
	_, err = io.Copy(b, response.Body)
	if err != nil {
		return err
	}
	body := b.Bytes()
	if response.StatusCode == http.StatusNotFound {
		return &ErrNotFound{req.URL.String()}
	} else if response.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed - %q.", response.Status)
	}

	err = easyjson.Unmarshal(body, value)
	if err != nil {
		return fmt.Errorf("failed to parse output. Error: %v", err)
	}
	return nil
}

func (kc *kubeletClient) GetSummary(ctx context.Context, node *corev1.Node) (*Summary, error) {
	port := kc.defaultPort
	nodeStatusPort := int(node.Status.DaemonEndpoints.KubeletEndpoint.Port)
	if kc.useNodeStatusPort && nodeStatusPort != 0 {
		port = nodeStatusPort
	}
	addr, err := kc.addrResolver.NodeAddress(node)
	if err != nil {
		return nil, fmt.Errorf("unable to extract connection information for node %q: %v", node.Name, err)
	}
	url := url.URL{
		Scheme:   kc.scheme,
		Host:     net.JoinHostPort(addr, strconv.Itoa(port)),
		Path:     "/stats/summary",
		RawQuery: "only_cpu_and_memory=true",
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	summary := &Summary{}
	client := kc.client
	if client == nil {
		client = http.DefaultClient
	}
	err = kc.makeRequestAndGetValue(client, req.WithContext(ctx), summary)
	return summary, err
}

func (kc *kubeletClient) getBuffer() *bytes.Buffer {
	return kc.buffers.Get().(*bytes.Buffer)
}

func (kc *kubeletClient) returnBuffer(b *bytes.Buffer) {
	b.Reset()
	kc.buffers.Put(b)
}

To add metrics to our Kubernetes environment, we'll use Helm to install Prometheus.

First we need helm installed as a client on our workstation, and then we can install the Kubernetes side component in our EKS system.  Get the helm binary for your environment here:

MacOSX:
https://storage.googleapis.com/kubernetes-helm/helm-v2.11.0-darwin-amd64.tar.gz

Linux:
https://storage.googleapis.com/kubernetes-helm/helm-v2.11.0-linux-amd64.tar.gz

Windows:
https://storage.googleapis.com/kubernetes-helm/helm-v2.11.0-windows-amd64.zip

Or use a package manager.

Then we can install the RBAC configuration for tiller so that it has the appropriate access, and lastly we can initialize our helm system:

kubectl create -f helm-rbac.yaml
helm init --service-account=tiller

Once Helm is installed, launching Prometheus is a simple command, though note that we are defining the storage class that Prometheus should use to store it's metrics:
helm install --name promeks --set server.persistentVolume.storageClass=gp2 stable/prometheus

And lastly, we want to expose the Prometheus UI so that we can have a look at some of the Pod/Container level metrics:


kubectl --namespace default port-forward $(kubectl get pods --namespace default -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}") 9090 &

Once the portforward is working, we can point a web browser at:

http://localhost:9090

look to see what metrics are being gathered.

container_cpu_usage_seconds_total

And we can also generate a little load if we'd like:

kubectl run --image rstarmer/curl:v1 curl
kubectl exec -it $(kubectl get pod -l run=curl -o jsonpath={.items..metadata.name}) -- \
sh -c 'while [[ true ]]; do curl -o - http://hostname-v1/version/ ; done'


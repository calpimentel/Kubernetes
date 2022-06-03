# Adding ingress is a Kubernetes function

We'll add the Traefik load balancer as an ingress function, and make use of the EKS integration with Amazon ELB to enable external access.

As ingress can route based on DNS, we can also do a little DNS manipulation to get traffic routed to our resources.

1) Since we're using 1.10.0 Kubernetes (or newer) we'll need to make sure we have a cluster role binding for the services to use:

kubectl apply -f https://raw.githubusercontent.com/containous/traefik/master/examples/k8s/traefik-rbac.yaml

2) We'll leverage the deployment model for our ingress controller, as we don't necessarily want to bind host address, and would rather have the ingress transit through the normal kube-proxy functions (note that we're changing the default "NodePort" type to "LoadBalancer"):

kubectl apply -f <(curl -so - https://raw.githubusercontent.com/containous/traefik/master/examples/k8s/traefik-deployment.yaml | sed -e 's/NodePort/LoadBalancer/')

3) We can now expose our hostname app as an ingress resource:

kubectl create -f hostname-ingress.yaml

Note that it assumes a hostname of traefik-ui.minikube, so we can confirm access as follows:

1) Get the loadbalancer service address:

export INGRESS=`kubectl get svc -n kube-system traefik-ingress-service -o jsonpath={.status.loadBalancer.ingress[0].hostname}`

Capture the actual IP of one of the loadbalancers in the set:

export INGRESS_ADDR=`host $INGRESS | head -1 | cut -d' ' -f 4`

Verify that it's responding to web requests:

curl -sLo /dev/null -Hhost:hostname-v1.local http://${INGRESS_ADDR}/ -w "%{http_code}\n"

Add an entry to the local /etc/hosts file to point to our resource:

echo "$INGRESS_ADDR hostname-v1.local" | sudo tee -a /etc/hosts

Now try pointing a web browser at that hostname:

http://hostname-v1.local

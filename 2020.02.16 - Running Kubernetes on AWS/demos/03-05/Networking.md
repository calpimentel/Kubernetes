Networking in EKS uses the VPC-CNI project to use the AWS VPC network model to provide connectivity across the cluster.  This is more efficient than having another layer of networking (e.g. Flannel, Calico, Weave, etc.) deployed as an overlay on top of the system, and maps perfectly into the VPC environment, using the VPC network management and IPAM services to support address management further improving the efficiency of the overall Kubernetes deployment.

We can see this in action by looking at the network information for any pod in the system:

kubectl run --image alpine alpine sleep 3600
IPs=`kubectl get pod $(kubectl get pod -l run=alpine -o jsonpath={.items..metadata.name}) -o yaml | awk '/IP/ {print $2}'`
for n in $IPs; do kubectl exec -it $(kubectl get pod -l run=alpine -o jsonpath={.items..metadata.name})  traceroute $n ; done

In order to enable policy, a CNI network needs to be in place, and by default the VPC based networking in EKS is already configured appropriately.  Policy however is not part of the VPC networking provided by Amazon, and instead, an integration with the Calico policy manager has been integrated with the VPC CNI service.

Enabling this function is as simple as launching the calico CNI policy manifest from the Amazon VPC CNI project:

kubectl apply -f https://raw.githubusercontent.com/aws/amazon-vpc-cni-k8s/release-1.1/config/v1.1/calico.yaml

This will create a daemonset running the calico policy engine on each configured node.

Let's run a container with curl enabled to test our target system (hostname-v1 from the initial install):

kubectl run --image rstarmer/curl:v1 curl

And let's verify that we can communicate to the http://hostname-v1 service endpoint:

kubectl exec -it $(kubectl get pod -l run=curl -o jsonpath={.items..metadata.name})  -- curl --connect-timeout 5 http://hostname-v1

Now, we can first disable network access by installing the baseline "Default Deny" policy, which will break our application access:

kubectl apply -f default-deny.yaml

And check to see that we can't communicate:

kubectl exec -it $(kubectl get pod -l run=curl -o jsonpath={.items..metadata.name})  -- curl --connect-timeout 5 http://hostname-v1

And then we can add back in a rule allowing access again:

kubectl apply -f allow-hostname.yaml

And one more check:

kubectl exec -it $(kubectl get pod -l run=curl -o jsonpath={.items..metadata.name})  -- curl --connect-timeout 5 http://hostname-v1


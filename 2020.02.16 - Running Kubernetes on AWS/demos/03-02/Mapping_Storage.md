Once storage classes are defined, mapping uses the standard Kubernetes models, and as we defined "Retain" as the reclaim policy, we have storage
that maintains persistence even when we delete the PersistentVolumeClaim and the Pod that claimed the storage.

Let's create a simple app that gets a 10G volume and mounts it into the web directory:

kubectl create -f hostname-volume.yaml

We can see that there is now a PV that is created with a PVC that claims it, and if we check our pod, we can see that our www directory is empty (because it has a 1G volume mounted to it).

kubectl get pv
kubectl get pvc -o wide
kubectl exec -it $(kubectl get pod -l app=hostname-volume -o jsonpath={.items..metadata.name}) -- df -h /www

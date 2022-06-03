The solution is simple:

First remove the "stack" we created earlier:

kubectl delete -f hostname-volume.yaml

Then clean up the PV itself.  Note, you do _not_ want to do this if there are more than one PVs created, as this will grab _all_ of the volumes and delete them!

kubectl delete pv $(kubectl get pv -o jsonpath={.items..metadata.name})



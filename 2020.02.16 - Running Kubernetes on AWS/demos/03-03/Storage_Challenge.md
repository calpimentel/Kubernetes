A simple challenge.  We can clean up our hostname-volume "stack" with:

kubectl delete -f hostname-volume.yaml

But this leaves a resource behind. If we really want to clean up _all_ of our storage, how would we go about cleaning up the volumes themselves.

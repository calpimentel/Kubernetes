By default EKS doesn't have any storage classes defined, and we need to have a storage class model in order to be able to create persistent storage.

Luckily the 'plumbing' is already there, and we simply have to enable our storage class connection to the underlying EBS service.

kubectl apply -f gp-storage.yaml

This creates a "standard" EBS based volume when a persistent volume request is created.  Size is determined by the persistent volume request.

In addition, we have configured this resource as our default, so if an application asks for storage without defining a class, we'll get this class configured.

Creating some Fast (100iops/GB) SSD Storage is also straightforward:

kubectl apply -f fast-storage.yaml

Additional parameters are available to tune for different classes of storage, and they are defined here:
https://kubernetes.io/docs/concepts/storage/storage-classes/#aws-ebs

Users within the EKS environment are authenticated against AWS IAM, which provides enhanced security.  If we add our 'clusterUser' credentials to the local aws client:

cat ~/Downloads/credentials-2.csv

aws configure --profile=clusterUser

export AWS_PROFILE=clusterUser

We will see that kubernetes will still try to talk to the API, but wil fail:

kubectl get pods

Adding additional users to the kubernetes cluster in EKS is done by adding new
users to the "system:masters" group which maps to the equivalent of the ClusterAdmin role in Kubernetes RBAC rules.

The key parameter we need is the User's IAM ARN, which can be pulled from the User IAM page in the AWS console:
data:
  mapUsers: |
    - userarn: USER-ARN
      username: admin
      groups:
        - system:masters

We need to add the mapUsers: section to the aws-auth-cm.yaml document, and we can do that either locally and "apply" the changes, or we can edit the document in place in the kubernetes service.

We will edit the file in place, as we don't want to have to recreate the worker node role mappings which are part of the same auth structure:

export AWS_PROFILE=clusterAdmin
kubectl edit configmap aws-auth -n kube-system

Once we're done with the edit, we can switch back to our clusterUser and we should have access to the system:

export AWS_PROFILE=clusterUser
kubectl get pods

Add the mapUsers: section right after the data: key, above the mapRoles: | line. It should look similar to the aws-auth-cm.yaml document.

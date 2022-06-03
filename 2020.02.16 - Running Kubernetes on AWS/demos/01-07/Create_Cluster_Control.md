# Create Cluster

Once we have created our IAM roles policies, and generated our VPC we're ready to launch our EKS control plane.

We can launch via the EKS console: https://console.aws.amazon.com/eks

Much of this information is available via pull-down items in the EKS console, but it's good to confirm that the IDs match which is why we captured the data into our outputs.txt file.

Once the cluster is "up" (which may take 10-15 minutes!) we can configure out kubectl command with the aws cli:
First we need to log into the CLI using the Key and Secret we downloaded for the clusterAdmin (credentials.csv):

aws configure --profile=clusterAdmin

export AWS_PROFILE=clusterAdmin

aws eks update-kubeconfig --name classCluster

# Verify Kubernetes access
And lastly, we should be able to confirm our access

kubectl get pods
kubectl get nodes

And can even launch a manifest like a simple hostname response app:

kubectl apply -f hostname.yaml

But there's one thing missing, no worker nodes!  We'll fix that next.


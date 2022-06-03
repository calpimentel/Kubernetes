1) Proper IAM configuration
 Create a Role defined with EKS permissions
 EKS:* policy applied to user/group
2) VPC created for EKS use
 Create EKS VPC with cloudformation
3) Create cluster control plane
 https://us-west-2.console.aws.amazon.com/eks
4) Establish Kubectl credentials (install aws-iam-authenticator and kubectl)
 aws eks update-kubeconfig --name cluster_name
5) Create worker nodes
 Create an autoscaling group of nodes with cloudformation
6) Create an aws-auth configmap to allow the nodes to register
 kubectl apply -f aws-auth-cm.yaml
7) Wait for the nodes to register
 kubectl get nodes -w
8) Ensure you can create an ELB
aws iam create-service-linked-role --aws-service-name elasticloadbalancing.amazonaws.com
9) Use your Kubernetes environment!

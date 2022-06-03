# Create worker nodes
https://console.aws.amazon.com/cloudformation/

From S3 template
https://amazon-eks.s3-us-west-2.amazonaws.com/cloudformation/2018-08-30/amazon-eks-nodegroup.yaml
Region	                     Amazon EKS-optimized   AMI	with GPU support
US West (Oregon) (us-west-2)  ami-0a54c984b9f908c81	ami-0731694d53ef9604b
US East (N. Va)  (us-east-1)	ami-0440e4f6b9713faf6	ami-058bfb8c236caae89
EU (Ireland)     (eu-west-1)	ami-0c7a4976cb6fafd3a	ami-0706dc8a5eed2eed9

Capture the worker ARN (NodeInstanceRole) from the CloudFormation output.
(store it in our outputs file)

We'll also be updating the aws-auth-cm.yaml document with that ARN so that the newly created workers have authorization to connect to our EKS control plane, and then installing the configmap to enable the authentication:

kubectl apply -f aws-auth-cm.yaml

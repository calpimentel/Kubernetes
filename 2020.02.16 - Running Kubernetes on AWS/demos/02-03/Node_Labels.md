In order to label the nodes, we need to create another set of workers, and we can then include an "extra args" parameter to add our unique label.  We could label nodes with:

kubectl label nodes <node-name> <label-key>=<label-value>

and then select the node group with a mapping like the following (which would be in the pod spec template):
```
  containers:
  - name: nginx
    image: nginx
    imagePullPolicy: IfNotPresent
  nodeSelector:
    <label-key>: <label-value>
```

But if the node is replaced or scaled into our out of an autoscalegroup, then we would loose that labeling.

To add labels by default when we create the ASG definition via cloudformation, we add:

--kubelet-extra-args '--node-labels "nodetype=generalpurpose"'

to the *'BootstrapArguments'* section when following the create worker process as before:
https://console.aws.amazon.com/cloudformation/

From S3 template
https://amazon-eks.s3-us-west-2.amazonaws.com/cloudformation/2018-08-30/amazon-eks-nodegroup.yaml
Region	                     Amazon EKS-optimized   AMI	with GPU support
US West (Oregon) (us-west-2)  ami-0a54c984b9f908c81	ami-0731694d53ef9604b
US East (N. Va)  (us-east-1)	ami-0440e4f6b9713faf6	ami-058bfb8c236caae89
EU (Ireland)     (eu-west-1)	ami-0c7a4976cb6fafd3a	ami-0706dc8a5eed2eed9

Capture the worker ARN (NodeInstanceRole) from the CloudFormation output.
(store it in our outputs file) and use it to update the aws-auth-cm.yaml document (note the version here needs the ARN from the original group of workers as well)

Once those nodes come online we can launch another version of our hostname that will only deploy onto the new nodes.

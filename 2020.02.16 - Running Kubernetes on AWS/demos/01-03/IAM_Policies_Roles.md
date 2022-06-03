In order to access the EKS cluster services (CRUD operations, and even just being able to connect to the cluster), an IAM policy is needed.  The following, while "open" only allows eks service access, and is appropriate for most users who need to manipulate the EKS service.

https://console.aws.amazon.com/iam/
Create a Policy with the following JSON object, and name it AdminEKSPolicy

iam_eks_policy.json
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "eks:*"
            ],
            "Resource": "*"
        }
    ]
}
```

We also need CloudFormation access which we can call AdminCloudFormationPolicy
iam_cf_policy.json
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "*",
            "Resource": "*"
        }
    ]
}
```

# Create IAM role
Create a Role for EKS:
https://console.aws.amazon.com/iam/
Capture the Role ARN for use in your EKS deployment. there is an output.txt file that is a template for capturing these IDs that are needed for subsequent stages.

We can use the pre-defined EKS role "model" for this role, which we will call AdminEKSRole

This will have two Amazon defined policies:

AmazonEKSServicePolicy 
AmazonEKSClusterPolicy 

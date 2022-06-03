While we can add policy to our normal users in AWS, for this class we will create
two users, an admin user and a second eks system user.

User 1:
 clusterAdmin
 - eks admin policy
 - k8s admin "system:master" group
   the following policies:
   AdminEKSPolicy
   AdminCloudFormationPolicy
   AmazonEKSServicePolicy 
   AmazonEKSClusterPolicy 

User 2:
 clusterUser
 - no IAM policies
 - k8s admin "system:master" group

For both users, create programmatic credentials, and for the admin user, create a password credential as well.

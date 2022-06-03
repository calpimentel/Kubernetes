In order to make use of the ASG that the default node cloudformation template creates, we need to add a policy to the group via the EC2 console.

Select Autoscaling Groups
Create a Policy
 - simple policy defines a target CPU utilization.  I find 70% works well
 for many environments.

In our unused cluster, the number of instances should start to shrink.

## Install kubectl as normal from the instructions found here:
https://kubernetes.io/docs/tasks/tools/install-kubectl/

Linux:
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl

MacOS:
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/darwin/amd64/kubectl

Windows:
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.12.0/bin/windows/amd64/kubectl.exe

## We also need the aws-iam-authenticator binary:
https://docs.aws.amazon.com/eks/latest/userguide/configure-kubectl.html

You need the binary appropriate to your OS:
Linux:
curl -sLO https://amazon-eks.s3-us-west-2.amazonaws.com/1.10.3/2018-07-26/bin/linux/amd64/aws-iam-authenticator

MacOS:
curl -sLO https://amazon-eks.s3-us-west-2.amazonaws.com/1.10.3/2018-07-26/bin/darwin/amd64/aws-iam-authenticator

Windows:
curl -sLO https://amazon-eks.s3-us-west-2.amazonaws.com/1.10.3/2018-07-26/bin/windows/amd64/aws-iam-authenticator.exe

In both cases, make the binary executable if necessary (chmod +x), and copy it to a directory in the command PATH (/usr/local/bin or %system32%/)

# Install the AWS CLI if you don't already have it
https://docs.aws.amazon.com/cli/latest/userguide/installing.html
(note, if you are on a Mac and don't have pip installed: sudo easy_install pip)

Install the awscli tool via python pip:
pip install awscli --upgrade --user


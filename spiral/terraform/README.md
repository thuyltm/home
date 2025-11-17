# What is Terraform?
Infrastructure as Code (IaC) tools allow you to manage infrastructure with configuration files rather than through a graphical user interface

Terraform is an infrastructure as code tool that lets you build, change and version infrastructure safely, consistently and repeatable way

Terraform lets you define resources and infrastructure in human-readable, declarative configuration files, and manages your infrastructure's lifecycle

There are many benefits:
- Terraform can manage infrastructure on **multiple cloud platforms**
- The **human-readable configuration** language helps you write infrastrucrure code quickly
- Terraform's state allows you to **track resouce changes** throughout your deployments
- You can commit your configuratios to **version control** to safely collaborate on infrastructure

# Terraform CommanLines
```sh
#Initialize the directory
terraform init
#Format and validate the configuration
terraform fmt
terraform validate
#Generate a plan, this will read your configuration files, compare them to the current state 
#(if any resources are alerady deployed and manageed by Terraform), and display the proposed changes
terraform plan -out=example.tfplan
#Crate or Update infrastructe
terraform apply "example.tfplan"
terraform destroy
```



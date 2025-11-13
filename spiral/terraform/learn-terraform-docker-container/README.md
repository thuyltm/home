# Review the file configuration **main.tf**
The required **providers** Terraform will use to provision your infrastructure. Terraform installs providers from the Terraform Registry by default

The **source** attribute defines an optional hostname, a namespace, and the provider type

## What is Providers
The provider is a plugin that Terraform uses to create and manage your resources. You can use multiple provider blocks in your Terraform configuration to manage resouces from different providers. You can even use different providers together. For example, you could pass the Docker Image ID to a Kubernetes service

## What is Resources
Resource blocks define components of your infrastructure. A resource might be a physical or virtual component such as a Docker container, or a Heroku application

Resource blocks have two strings before the block: the **resource type** and **resource name**. The prefix of the type maps to the name of the provider. Together, the resource type and resource name form a unique ID for the resource

Resource blocks contain arguments which you use to configure the resource. Our [providers reference documents](https://registry.terraform.io/browse/providers) the required and optional arguments for each resource
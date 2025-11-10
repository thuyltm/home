# Theory
- Ansible uses playbook to describe automation jobs
- Ansible is designed for multi-tier deployment
- Ansible is completely agentless which means Ansible works by connecting your nodes through ssh (by default)
- After connecting to your nodes, Ansible pushes small programs called as Ansible Modules. Ansible runs that modules on your 
 nodes and removes them when finished
- Ansible manages your inventory in simple text files which are the hosts file. Ansible uses the hosts file where one can group the 
 hosts and can controll the actions on a specific group in the playbooks

# What is configuration management
If you want to install the new version of WebLogic/WebSphere server on all of the machines present in your enterprise, it is not feasible for you to manually go and update each and every machine

You can install WebLogic/WebSphere in one go on all of your machines with Ansible playbooks and inventory written in the most simple way. All you have to do is list out the IP addresses of your nodes in the inventory and write a playbook to install WebLogic/WebSphere. Run the playbook from your control machine & it will be installed on all your nodes

[<img src="https://www.tutorialspoint.com/ansible/images/ansible_works.jpg">](https://www.tutorialspoint.com/ansible/ansible_introduction.html)

- The management node controls the entire execution of the playbook
- The inventory file provides the list of hosts where the Ansible modules needs to be run
- The management node does a ssh connection and exectures the small modules on the hosts machine and installs the product/software

# What is Ansible
Ansible does not add any database. It does not require any daemons to start or keep it running. While manging remote machines, Ansible does not leave any software installed or running on them

You can practive Ansible for free using a variety of online platform
- DiveInto.com Ansible Lab: This fully-featured lab environment runs instantly in Google Cloud Shell. You only need to sign in with a Google account, and no setup is required on your end

# TUTORIAL
https://github.com/labex-labs/ansible-free-tutorials?tab=readme-ov-file
https://kodekloud.com/free-courses

# VIRTUAL ENVIRONMENT
https://diveinto.com/playgrounds/ansible-lab

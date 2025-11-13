```sh
#!/bin/bash
ansible-playbook -i hosts ping.yml
ansible-playbook -i hosts ping.yml -v

ansible-playbook -i hosts ping_all.yml
ansible-playbook -i hosts ping_data.yml -v
#running the ping module directly from the commandline without creating a playbook
ansible local -i hosts -m ping
#you can also specify multiple groups or patterns
ansible 'webservers:local' -i hosts -m ping
ansible all -i hosts -m ping -a data=Command line test
ansible-inventory -i hosts --list
ansible-inventory -i hosts ping_comprehensive.yml
```

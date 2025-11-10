#!/bin/bash

# Update system packages
echo "🔄 Updating system packages..."
sudo apt update
echo

# Install sshpass if not already installed
# sshpass allows users to provide an SSH password in a script or commandline,
# rather than having to manually type it when prompted by the ssh client
echo "🔄 Installing sshpass..."
sudo apt install -y sshpass
echo

# Create SSH keys, overwriting existing keys without prompting
echo "🔑 Generating new SSH keys and overwriting old ones..."
#/dev/null is used to suppress unwanted output from commands or scripts, making the terminal output cleaner
# 2>&1 redirects file descriptor 2 (stderr) to file descriptor 1 (stdout)
ssh-keygen -t rsa -b 2048 -f ~/.ssh/id_rsa -q -N "" <<< y >/dev/null 2>&1
echo

# Remove known_hosts file to avoid issues with reused IP addresses
echo "🔍 Removing known_hosts file..."
rm -f ~/.ssh/known_hosts
echo

# Define the password for ssh-copy-id
PASSWORD='password'

# Define base hosts and number of instances
HOSTS=(ubuntu centos)
INSTANCES=(1 2 3)
USERS=(ansible root)

# Loop through each host and instance combination to copy the SSH key
for HOST in "${HOSTS[@]}"; do
    for INSTANCE in "${INSTANCES[@]}"; do
        for USER in "${USERS[@]}"; do
            TARGET="$HOST$INSTANCE"
            echo "🚀 Copying SSH key to $USER@$TARGET..."
            # Use sshpass to pass the password and ssh-copy-id to copy the SSH key
            sshpass -p "$PASSWORD" ssh-copy-id -o StrictHostKeyChecking=no "$USER@$TARGET"
            if [ $? -eq 0 ]; then
                echo "✅ Successfully copied SSH key to $USER@$TARGET"
            else
                echo "❌ Failed to copy SSH key to $USER@$TARGET"
            fi
            echo
        done
    done
done
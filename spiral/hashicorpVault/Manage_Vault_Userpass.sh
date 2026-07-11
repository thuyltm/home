#! /bin/sh
vault login
vault auth enable userpass
vault write auth/userpass/users/test password=test policies=admins
vault login -method=userpass username=test password=test
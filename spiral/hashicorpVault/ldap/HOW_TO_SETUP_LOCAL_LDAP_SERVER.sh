#! /bin/bash
sudo apt-get install ldap-utils
docker run --name vault-openldap \
  --env LDAP_ORGANISATION="learn" \
  --env LDAP_DOMAIN="learn.example" \
  --env LDAP_ADMIN_PASSWORD="2LearnedVault" \
  -p 389:389 \
  -p 636:636 \
  --detach \
  --rm osixia/openldap:1.4.0
ldapadd -cxWD "cn=admin,dc=learn,dc=example" -f ldap.ldif
# Enter LDAP Password: 2LearnedVault
#adding new entry "ou=groups,dc=learn,dc=example"
#adding new entry "ou=users,dc=learn,dc=example"
#adding new entry "cn=dev,ou=groups,dc=learn,dc=example"
#adding new entry "cn=alice,ou=users,dc=learn,dc=example"

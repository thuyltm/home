Roles in Vault add many configuration settings to an __auth__ method or __secrets__ engine.

A __role__ is collection of parameters that you group together to simplify __plugin configuration__. Authentication requests only need to pass the role name to Vault. Vault will read the role configuration and issue a token based on the settings configured for that role.

The user authentication plugin selected for the POC is the _userpass_ auth method. This auth method does not support the use of roles.

The _kubernetes_ auth method that works with most types of Kubernetes deployments use roles
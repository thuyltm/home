terraform {
    required_providers {
        vault = {
            source = "hashicorp/vault"
            version = "5.0.0"
        }
    }
}

provider "vault" {
    auth_login {
        path = "auth/userpass/login/${var.login_username}"
        parameters = {
            password = var.login_password
        }
    }
}


#! /bin/sh
minikube start
###################################################################################
## Generate Vault Certificate
###################################################################################
####################################
# 1. Generate the Vault private key
####################################
openssl genrsa -out tls.key 2048
###################################################
# 2. Generate the Certification Signing Request
###################################################
openssl req -new -key tls.key -out tls.csr -config tls-config.conf
################################################################################################################
# 3. Generate Vault Certificate by Sending the Certification Signing Request to Kubernetes Certificate Authority
################################################################################################################
# request data
cat tls.csr|base64|tr -d '\n'
kubectl create -f tls-sign-request.yaml
# certificatesigningrequest.certificates.k8s.io/tls.svc created
#kubectl get CertificateSigningRequest tls.svc
#NAME        AGE     SIGNERNAME                                    REQUESTOR              REQUESTEDDURATION   CONDITION
#tls.svc   2m25s   kubernetes.io/kubelet-serving                 minikube-user          100d                Pending
#kubectl describe csr tls.svc
####################################################
# APPROVE THE CERTIFICATION SIGNING REQUEST IN K8S
####################################################
kubectl certificate approve tls.svc
# certificatesigningrequest.certificates.k8s.io/tls.svc approved
# Confirm the certificate was issued
# kubectl get CertificateSigningRequest tls.svc
#NAME        AGE   SIGNERNAME                      REQUESTOR       REQUESTEDDURATION   CONDITION
#tls.svc   16m   kubernetes.io/kubelet-serving   minikube-user   100d                Approved,Issued
##########################################################
# Retrieve the tls certification
############################################################
kubectl get csr tls.svc -o jsonpath='{.status.certificate}' | openssl base64 -d -A -out tls.crt
# openssl x509 -in tls.crt -text -noout
###################################################
# retrieve kubernetes CA certification
###################################################
kubectl config view --raw --minify --flatten -o jsonpath='{.clusters[].cluster.certificate-authority-data}' \
 | base64 -d k8s.ca

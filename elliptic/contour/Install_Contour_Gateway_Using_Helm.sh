#! /bin/bash

helm repo add contour https://projectcontour.github.io/helm-charts/
helm install contour contour/contour --namespace contour --create-namespace
kubectl -n contour get po,svc
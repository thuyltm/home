In OpenSearch, __kibanaserver__ is a default internal system user used by OpenSearch Dashboards to communicate with the OpenSearch cluster.

OpenSearch Dashboards connects to the backend OpenSearch nodes using the __opensearch.username__ and __opensearch.password__ parameters in your __opensearch_dashboards.yml__ config file.

The password for username __kibanaserver__ defined in your OpenSearch Security plugin's __internal_users.yml__ file

The kibanaserver use is assigned __the _kibana_server_ role__ to ensure it has administrative priviledge over __.kibana or .opensearch_dashboards index patterns__.
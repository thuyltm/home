```sh
# Starts a master instance on the machine the script is executed on.
# By default, Spark might only bind to localhost. 
# To accept connections from anywhere, set the SPARK_MASTER_HOST environment variable to 0.0.0.0 when starting the master. 
% SPARK_MASTER_HOST=0.0.0.0 $SPARK_HOME/sbin/start-master.sh
# INFO Master: Starting Spark master at spark://thuy-Vivobook-Go-E1504FA-E1504FA:7077
# INFO MasterWebUI: Bound MasterWebUI to 0.0.0.0, and started at http://localhost:8080
% $SPARK_HOME/sbin/start-worker.sh spark://thuy-Vivobook-Go-E1504FA-E1504FA:7077
# To run an interactive Spark sell against the cluster
% $SPARK_HOME/bin/spark-shell --master spark://thuy-Vivobook-Go-E1504FA-E1504FA:7077
% curl -XPOST http://thuy-Vivobook-Go-E1504FA-E1504FA:6066/v1/submissions/create \
--header "Content-Type:application/json;charset=UTF-8" \
--data '{
  "appResource": "",
  "sparkProperties": {
    "spark.master": "spark://thuy-Vivobook-Go-E1504FA-E1504FA:7077",
    "spark.app.name": "Spark Pi",
    "spark.driver.memory": "1g",
    "spark.driver.cores": "1",
    "spark.jars": ""
  },
  "clientSparkVersion": "",
  "mainClass": "org.apache.spark.deploy.SparkSubmit",
  "environmentVariables": { },
  "action": "CreateSubmissionRequest",
  "appArgs": [ "/home/thuy/Downloads/spark-4.1.1-bin-hadoop3/examples/src/main/python/pi.py", "10" ]
}'
```
1> To develop code, you load a specific the sbt project named hello-spark into IntelliJ idea . After reloading the sbt project, IntelliJ idea will auto-import 
dependencies and classpath into the sbt project

2> Run __sbt package__ to build a jar file

3> Start Master/Slave Local Spark Server
```sh
% $SPARK_HOME/sbin/start-master.sh
# You look at the master's web UI (http://localhost:8080 by default) to find spark://HOST:PORT URL which pass as the "master" argument to Spark Context
% $SPARK_HOME/sbin/start-worker.sh spark://thuy-Vivobook-Go-E1504FA-E1504FA:7077
```
4> Submit task to local process
```sh
spark-submit \
  --class "SimpleApp" \
  --master "local[4]" \
  target/scala-2.13/hello-spark_2.13-1.0.jar
```

Breakdown of the argument:
- __--master__: This is a command line option used by the spark-submit, pyspark or spark-shell scripts to define the cluster manager or master URL that the Spark application should connect to
- __local[4]__: This is a specific master URL, which instruct Spark to:
  - local: Running in "local mode" is meaning all Spark components (driver, master and executors) operator wihton a single JVM process on the machine whre the command is executed
  - [4]: Use exactly four worker threads (or CPU cores) for parrel task execution within that single JVM process

5> Submit task to Spark Standalone
```sh
spark-submit \
--class "EstimatorTransformerParamExample" \
--master spark://thuy-Vivobook-Go-E1504FA-E1504FA:7077 \
--deploy-mode client \
--name "Spark Pi Example" \
target/scala-2.13/hello-spark_2.13-1.0.jar
```

6> Submit task to Hadoop Yarn
```sh
% export HADOOP_HOME=/usr/local/hadoop
% export HADOOP_CONF_DIR=$HADOOP_HOME/etc/hadoop
% export YARN_CONF_DIR=$HADOOP_HOME/etc/hadoop
% spark-submit \
%  --class "PipelineExample" \
%  --master yarn \
%  --deploy-mode cluster \
%  --name "PipelineExample" \
%  target/scala-2.13/hello-spark_2.13-1.0.jar
INFO Client: Uploading resource file:/tmp/spark-5523bef9-30f5-4cc5-a615-093f24d1fe67/__spark_libs__2959459217800340450.zip -> hdfs://localhost:9000/user/thuy/.sparkStaging/application_1771901797707_0002/__spark_libs__2959459217800340450.zip
INFO Client: Uploading resource file:/home/thuy/Documents/Learn/home/perimeter/spark/hello-spark/target/scala-2.13/hello-spark_2.13-1.0.jar -> hdfs://localhost:9000/user/thuy/.sparkStaging/application_1771901797707_0002/hello-spark_2.13-1.0.jar
INFO Client: Uploading resource file:/tmp/spark-5523bef9-30f5-4cc5-a615-093f24d1fe67/__spark_conf__13090918525502185352.zip -> hdfs://localhost:9000/user/thuy/.sparkStaging/application_1771901797707_0002/__spark_conf__.zip
INFO Client: Submitting application application_1771901797707_0002 to ResourceManager
INFO YarnClientImpl: Submitted application application_1771901797707_0002
INFO Client: Application report for application_1771901797707_0002 (state: ACCEPTED)
INFO Client: Application report for application_1771901797707_0002 (state: FINISHED)
INFO Client: 
  client token: N/A
  diagnostics: N/A
  ApplicationMaster host: 172.19.0.1
  ApplicationMaster RPC port: 45115
  queue: root.default
  start time: 1771902010316
  final status: SUCCEEDED
  tracking URL: http://thuy-Vivobook-Go-E1504FA-E1504FA:8088/proxy/application_1771901797707_0002/
  user: thuy
```


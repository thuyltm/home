#### What is Apache Beam
Apache Beam is the easiest way to do batch and streaming data processing. Write once, run anywhere data processing for mission-critical production workloads.

1. Direct Runner
```sh
go run perimeter/beam/wordcount.go --input perimeter/beam/kinglear.txt --output perimeter/beam/counts.txt
```
2. Run with Spark Job Server in the Beam source
```sh
#INFO SparkContext: Running Spark version 3.5.0
#INFO SparkContext: OS info Linux, 6.17.0-14-generic, amd64
#INFO SparkContext: Java version 21.0.10
% export JDK_JAVA_OPTIONS="--add-exports=java.base/sun.nio.ch=ALL-UNNAMED"
% ./gradlew :runners:spark:3:job-server:runShadow -PsparkMasterUrl=spark://thuy-Vivobook-Go-E1504FA-E1504FA:7077
% go run wordcount.go --input kinglear.txt --output counts --runner spark --endpoint localhost:8099
```sh
#### Install Hadopp: 
1. Extract the hadoop extracted files to /usr/local/hadoop
```sh
tar -xf hadoop-3.4.2.tar.gz -C /usr/local/hadoop
```
2. Edit ~/.bashrc to add HADOOP_HOME and JAVA_HOME environment variables, then add $HADOOP_HOME/bin and $HADOOP_HOME/bin to the Path
3. Edit $HADOOP_HOME/etc/hadoop/hadoop-env.sh to set the JAVA_HOME path
#### Pseudo-distributed on the Single Node
By default, Hadoop is configured to run in a non-distributed mode, as a single Java process.

Hadoop can also be run on a single-node in a pseudo-distributed mode where each Hadoop daemon runs in a separate Java process

1. Start NameNode daemon and DataNode daemon
```sh
# Formating namenode is seen as preliminary to the start-dfs
% hdfs namenode -format
% start-dfs.sh
Starting namenodes on [http://localhost:9870]
Starting datanodes
Starting secondary namenodes [thuy-Vivobook-Go-E1504FA-E1504FA]
# The hadoop daemon log output is written to the $HADOOP_HOME/logs
# Browse the web interface to the NameNode http://localhost:9870/
```
2. Start ResouceManager daemon and NodeManager daemon
```sh
% start-yarn.sh
Starting resourcemanager
Starting nodemanagers
```
#### Cluster Setup
To start a Hadoop cluster, you will need to start both the HDFS and YARN cluster
1. Format a new distributed filesystem as hdfs
```sh
hdfs namenode -format
```
2. Start the HDFS NameNode with the following command on the designated node as hdfs
```sh
hdfs --daemon start namenode
```
3. Start the HDFS DataNode with the following command on the designated node as hdfs
```sh
yarn --daemon start resourcemanager
```
4. Start a NodeManager on each designated host as yarn
```sh
yarn --daemon start nodemanager
```
5. Start the MapReduce JobHistory Server with the following command, run on the designed server as mapred
```sh
mapred --daemon start historyserver
```
#### Extra step: Start the MapReduce JobHistory Server with the following command
```sh
mapred --daemon start historyserver
```
To configure the Hadoop cluster, you will need to configure the _environment_ in which the Hadoop daemons execute as well as the _configuration parameters_ for the Hadoop daemons

#### Configure Environment
- $HADOOP_HOME/etc/hadoop/hadoop-env.sh         
- $HADOOP_HOME/etc/hadoop/yarn-env.sh
- $HADOOP_HOME/etc/hadoop/mapred-env.sh

For example, To configure Namenode to use parallelGC and a 4GB Java Heap, the following statement should be added in hadoop-env.sh
```sh
export HDFS_NAMENODE_OPTS="-XX:+UseParallelGC -Xmx4g"
```

#### Configure the Hadoop Daemons
- $HADOOP_HOME/etc/hadoop/core-site.xml
- $HADOOP_HOME/etc/hadoop/hdfs-site.xml
- $HADOOP_HOME/etc/hadoop/yarn-site.xml
- $HADOOP_HOME/etc/hadoop/mapred-site.xml

List all worker hostnames or IP addresses in your _$HADOOP_HOME/etc/hadoop/workers_ file, one per line. Some scripts like 
```sh
% $HADOOP_HOME/sbin/start-dfs.sh
# OR
% $HADOOP_HOME/sbin/start-yarn.sh
```
will use the _$HADOOP_HOME/etc/hadoop/workers_ file to run commands on many hosts at once. In order to use this functionality, ssh trusts (via either passphraseless ssh) must be established for the accounts used to run Hadoop

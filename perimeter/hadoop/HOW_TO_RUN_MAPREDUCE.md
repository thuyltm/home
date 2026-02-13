1. Start HDFS and Yarn in Single Cluster
2. Load Hadoop Libs and Configuration for Yarn
```sh
% gedit $HADOOP_HOME/etc/hadoop/yarn-site.xml

<configuration>
	<property>
	    <name>yarn.application.classpath</name>
	    <value>
		$HADOOP_HOME/etc/hadoop/*,
		$HADOOP_HOME/share/hadoop/common/*,
		$HADOOP_HOME/share/hadoop/common/lib/*,
		$HADOOP_HOME/share/hadoop/hdfs/*,
		$HADOOP_HOME/share/hadoop/hdfs/lib/*,
		$HADOOP_HOME/share/hadoop/mapreduce/*,
		$HADOOP_HOME/share/hadoop/mapreduce/lib/*,
		$HADOOP_HOME/share/hadoop/yarn/*,
		$HADOOP_HOME/share/hadoop/yarn/lib/*
	    </value>
	</property>
	<property>
		<name>yarn.nodemanager.aux-services</name>
		<value>mapreduce_shuffle</value>
	</property>
	<property>
		<name>yarn.nodemanager.env-whitelist</name>
		<value>JAVA_HOME,HADOOP_COMMON_HOME,HADOOP_HDFS_HOME,HADOOP_CONF_DIR,CLASSPATH_PREPEND_DISTCACHE,
		       HADOOP_YARN_HOME,HADOOP_HOME,PATH,LANG,TZ,HADOOP_MAPRED_HOME</value>
	</property>
</configuration>
```
3. Prepare input data
```sh
hdfs dfs -mkdir -p /user/thuyltm
hdfs dfs -ls /user/thuyltm/
hdfs dfs -rm -r /user/thuyltm/output
hdfs dfs -ls hdfs://localhost:9000/user/
hdfs dfs -mkdir /user/thuyltm/input
hdfs dfs -put perimeter/hadoop/input/*.txt /user/thuyltm/input
```
4. Build WordCount_deploy.jar

Use the bazel build command with the _deploy.jar suffix appended to your target name
```sh
% bazel build //perimeter/hadoop:WordCount_deploy.jar
# Check the content of the WordCount_deploy.jar jar file whether there have class and inner classes inside this jar file
% jar tvf bazel-bin/perimeter/hadoop/WordCount_deploy.jar | grep "WordCount"
WordCount$IntSumReducer.class
WordCount$TokenizeMapper.class
WordCount.class
# Check if the jar file is built correctly
% java -jar bazel-bin/perimeter/hadoop/WordCount_deploy.jar 
```
_deploy.jar is an implicit output target that creates a single, self-container JAR suitable for deployment

5. Run
```sh
% hadoop jar bazel-bin/perimeter/hadoop/WordCount_deploy.jar /user/thuyltm/input /user/thuyltm/output
INFO mapreduce.Job: The url to track the job: http://thuy-Vivobook-Go-E1504FA-E1504FA:8088/proxy/application_1770886467729_0002/
```
6. Check the result
```sh
% hdfs dfs -cat /user/thuyltm/output/*
Bye     1
Goodbye 1
Hadoop  2
Hello   2
World   2
```

7. Run WordCount2_deploy.jar and check result
```sh
% hadoop jar bazel-bin/perimeter/hadoop/WordCount2_deploy.jar -Dwordcount.case.sensitive=true /user/thuyltm/input /user/thuyltm/output -skip /user/thuyltm/input/patterns.txt
% hadoop fs -cat /user/thuyltm/output/*
Bye     1
Goodbye 1
Hadoop  1
Hello   1
Hello   1
World   2
\       3
hadoop  1
```
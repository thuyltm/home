![HDFS Architecture](https://hadoop.apache.org/docs/r1.2.1/images/hdfsarchitecture.gif)

HDFS (Hadoop Distributed File System) is the primary data storage system for Apache Hadoop, designed to store massive datasets across clusters of inexpensive, commodity hardware.

HDFS has a master/slave architecture where the **NameNode** acts as the master managing file system metadata, while DataNodes store the actual data blocks

HDFS automatically replicates data blocks (default is 3x) across different nodes to ensure no data loss if a node fails
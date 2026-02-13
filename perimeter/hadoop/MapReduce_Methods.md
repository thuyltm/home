#### Mapper
Mapper maps input key/value pairs to a set of intermediate key/value pairs.

Mapper implementatios are passed to the job via _Job.setMapperClass(Class)_ method. All intermediate values associated with a given output key are subsequently grouped by the framework, and passed to the _Reducer(s)_ to determine the final output. Users can control the grouping by specifying a _Comparator_ via _Job.setGroupingComparatorClass(Class)_

The Mapper outputs are sorted and then partitioned per Reducer. Users can control which keys go to which Reducer by implementing a custom Partitioner.

The intermediate outputs are to be compressed and the CompressionCodec to be used via the Configuration

#### Reducer
Reducer implementations are passed the Job for the job via the Job.setReducerClass(Class) method. 

Reducer has 3 primary phases: shuffle, sort and reduce.
1. Shuffle, Sort: Input to the Reducer is the sorted output of the mappers. Therefore, the framework fetches the relevant partition of the output of all the mappers. The suffle and sort phases occurs simultaneously, while map-outputs are being fetched they are merged
2. Reduce: The framework groups Reducer inputs by keys
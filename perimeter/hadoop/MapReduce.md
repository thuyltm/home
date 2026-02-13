A MapReduce job usually splits the input dataset into independent chunks which are processed by the __map__ tasks in a completely parallel manner. The framework __sorts the outputs of the maps__, which are then input to the __reduce__ tasks

The Hadoop job client __submits__ the job (jar/executable, etc.) and configuration __to the ReourceManager__ which __distribute__ the software/configuration __to the workders__, schedule tasks and monitor them, provide status and diagnostic information to the job-client

The MapReduce framework views the input to the job as a set of <key, value> pairs and produces a set of <key, value> pairs as the output of the job
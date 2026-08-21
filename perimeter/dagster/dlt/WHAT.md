__Data load tool (dlt)__ is an open-source, lightweight Python library designed to simplify data loading. It takes care of many of the more tedious aspects of ETL, including schema management, data type handling, and normalization, so you can focus on what matters most

dlt helps you move data between systems without having to build and maintain all the supporting infrastucure yourself.

### Why Dagster combine with dlt
The simplicity and flexibility of lightweight ETL dlt framework integrate well with Dagster's robust orchestration capabilities. With Dagster, those responsibilities like setting up separate infrastructure - such as databases and task queues for state tracking and orchestration - are already built-in.
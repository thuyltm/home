### Why use dbt and Dagster together?
dbt Core can __only transform data that is alreay in a data warehouse__. The data transformed using dbt come from somewhere and go somewhere. Dbt can't extract from a source, load it into its final destination or automate either of these operation. This is __where Dagster comes in to orchestrate the entire workflow__.

Dagster separates data assets from the execution that produces them and gives you the ability to monitor and debug each dbt model individuallys
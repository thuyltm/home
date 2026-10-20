from pathlib import Path

from dagster_dbt import DbtProject

dbt_project_dir_value = Path(__file__).joinpath("../..", "analytics").resolve()
dbt_project = DbtProject(
    project_dir=dbt_project_dir_value if dbt_project_dir_value is not None else "src/dagster_dbt_quickstart/analytics",
    profiles_dir=dbt_project_dir_value if dbt_project_dir_value is not None else "src/dagster_dbt_quickstart/analytics",
    profile="dagster_dbt_university",   # Must match the top-level key in profiles.yml
    target="dev", 
)

# Compiles the manifest.json so Dagster can see the models
#####################################################################################################
# The prepare_if_dev() method automatically prepares your dbt project at run time during development,
# meaning you no longer have to run dbt parse
# The preparation process works by pulling the dbt project's dependencies 
# and reloading the manifest file to detect any changes
#####################################################################################################
dbt_project.prepare_if_dev() 

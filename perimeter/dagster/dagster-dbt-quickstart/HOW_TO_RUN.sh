#! /bin/sh
#####################################################################################
# scaffold a new dagster project
#####################################################################################
uvx create-dagster@latest project dagster-dbt
# activate the virtual environment
cd dagster-dbt
source .venv/bin/activate
######################################################################################
# Generate the dbt project's hidden target folder referenced by Dagster project
######################################################################################
dbt deps
dg check defs #-> produce the dbt  project's hidden target folder
dbt parse
######################################################################################
# Run Dagster project
######################################################################################
dg dev
######################################################################################
# Follow the task lineage diagram to run dependent tasks in sequential order
#####################################################################################
######################################################################################
# Dagster trigger @dbt_assets to build dbt model
######################################################################################
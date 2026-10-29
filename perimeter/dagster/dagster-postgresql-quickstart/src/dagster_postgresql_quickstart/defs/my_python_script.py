from dagster_pipes import open_dagster_pipes

if __name__ == "__main__":
    # Open the pipes session to establish the connection with Dagster
    with open_dagster_pipes() as context:
        
        # 1. Read information passed down from Dagster
        extra_param = context.get_extra("some_parameter")
        
        # ... Your data processing or heavy compute logic here ...
        
        # 2. Report the final asset materialization along with metadata
        context.report_asset_materialization(
            metadata={
                "row_count": {"raw_value": 1500, "type": "int"},
                "status": "success"
            }
        )

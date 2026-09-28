from google.cloud import bigquery
import os
os.environ["GOOGLE_API_USE_CLIENT_CERTIFICATE"] = "false"
import pandas as pd


def model(dbt, session):
    client = bigquery.Client()
    results = client.query_and_wait(
        """
        SELECT
            CONCAT(
                'https://stackoverflow.com/questions/',
                CAST(id as STRING)) as url, view_count
        FROM `bigquery-public-data.stackoverflow.posts_questions`
        WHERE tags like '%google-bigquery%'
        ORDER BY view_count DESC
        LIMIT 10
        """
    )
    new_list = [[row.url, row.view_count] for row in results]
    # Create DataFrame
    df = pd.DataFrame(new_list, columns=["URL", "View Count"])

    return df
        

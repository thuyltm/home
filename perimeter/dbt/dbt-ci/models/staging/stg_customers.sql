with source as (
    select * from {{ ref('raw_customers') }}
)

select
    customer_id,
    lower(first_name) as first_name,
    lower(last_name) as last_name,
    signup_date
from source
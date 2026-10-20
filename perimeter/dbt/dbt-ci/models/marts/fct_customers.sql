with customers as (
    select * from {{ ref('stg_customers') }}
)

select
    customer_id,
    first_name || ' ' || last_name as full_name,
    signup_date
from customers
{{ config(materialized='table') }} -- Explicitly creates this model as a table
with raw_trips as (
    select *
    from {{ source('raw_taxis', 'trips') }}
)
select
    {{
        dbt_utils.generate_surrogate_key([
            'PULocationID',
            'DOLocationID',      
            'VendorID',
            'tpep_pickup_datetime',
            'tpep_dropoff_datetime',
        ])
    }} as trip_id,
    date_diff('minutes', tpep_pickup_datetime, tpep_dropoff_datetime) as duration,
    case payment_type
        when 0 then 'Unknown'
        when 1 then 'Credit Card'
        when 2 then 'Cash'
        when 3 then 'No Charge'
        when 4 then 'Dispute'
        when 5 then 'Unknown'
        when 6 then 'Voided Trip'
    end as payment_type,
    case RatecodeID
        when 1 then 'Standard rate'
        when 2 then 'JFK'
        when 3 then 'Newark'
        when 4 then 'Nassau or Westchester'
        when 5 then 'Negotiated fare'
        when 6 then 'Group ride'
        when 99 then 'Unknown'
        else 'Unknown'
    end as rate_code,
    case VendorID
        when '1' then 'Creative Mobile Technologies, LLC'
        when '2' then 'VeriFone Inc.'
        else 'Unknown'
    end as vendor_name,
    *
    exclude (
        payment_type,
        RatecodeID,
        VendorID
    )
from raw_trips
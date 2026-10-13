with raw_zones as (
    select *
    from {{ source('raw_taxis', 'zones') }}
)
select
    objectid,
    zone as zone_name,
    borough,
    zone_name like '%Airport' as is_airport,
    LocationID
from raw_zones
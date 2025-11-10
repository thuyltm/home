#!/bin/bash

## System health check script
echo "HEALTH_CHECK_START: $(date)"

## CPU load
cpu_load=$(uptime | awk -F'load average: ' '{print $2}' | cut -d, -f1)
echo "CPU_LOAD: $cpu_load"
if (($(echo "$cpu_load > 1.0" | bc -l))); then
    echo "CPU_STATUS: WARNING"
else
    echo "CPU_STATUS: OK"
fi

## Memory usage
mem_total=$(free | grep Mem | awk '{print $2}')
mem_used=$(free | grep Mem | awk '{print $3}')
mem_pct=$(echo "scale=2; $mem_used / $mem_total * 100" | bc)
echo "MEM_USAGE_PCT: $mem_pct"
if (($(echo "$mem_pct > 80" | bc -l))); then
    echo "MEM_STATUS: WARNING"
else
    echo "MEM_STATUS: OK"
fi

## Disk usage
disk_usage=$(df -h / | tail -1 | awk '{print $5}' | tr -d '%d')
echo "DISK_USAGE_PCT: $disk_usage"
if [ "$disk_usage" -gt 80 ]; then
    echo "DISK_STATUS: WARNING"
else
    echo "DISK_STATUS: OK"
fi

## Check for zombie processes
zoombie_count=$(ps aux | grep -c Z)
echo "ZOOMBIE_PROCESSES: $zoombie_count"
if [ "$zoombie_count" -gt 0 ]; then
    echo "ZOOMBIE_STATUS: WARNING"
else
    echo "ZOOMBIE_STATUS: OK"
fi
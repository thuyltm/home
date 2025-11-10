#!/bin/bash
## Print system information
echo "=== System Information ==="
echo "Hostname: $(hostname)"
echo "Date: $(date)"
echo "Kernel: $(uname -r)"
echo "Memory:"
free -h

## Display CPU info
echo "CPU_MODEL: $(grep 'mode name' /proc/cpuinfo | head -1 | cut -d ':' -f2 | xargs)"
echo "CPU_CORES: $(grep -c 'processor' /proc/cpuinfo)"
mem_total=

# Generate some standard output
echo "=== Standard Output ==="
echo "This is standard output"
echo "Hello from the script!"

## Generate some standard error
## in a shell script, >&2 is a redirection operator that sends the sdout of a command
## to the stderr stream
echo "=== Standard Output ===" >&2
echo "This is standard error" >&2
echo "An example error message" >&2

## Exit with a specific code
exit 0

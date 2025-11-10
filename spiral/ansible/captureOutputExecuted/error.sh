#!/bin/bash

## Print some standard output
echo "Starting error demonstration script"
echo "This will appear in stdout"

## Print some standard error
echo "This will appear in stderr" >&2
echo "Error: Something went wrong!" >&2

## Exit with a non-zero code to indicate error
exit 1
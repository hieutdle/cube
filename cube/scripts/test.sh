#!/bin/bash

# Exit script on any error

set -e 

echo "Running tests for all folders ..."

go test ./...

echo "All tests completed successfully!"
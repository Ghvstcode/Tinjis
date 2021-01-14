#!/bin/bash

set -e

cd K8s-specs
echo "Creating Resources"
kubectl create -f antaeus.yml
echo "Created, Connect on port 30475"

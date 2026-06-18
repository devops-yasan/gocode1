#!/usr/bin/env bash
set -euo pipefail

ENVIRONMENT="${ENVIRONMENT:-staging}"

echo "Deploying application to ${ENVIRONMENT}"
kubectl apply -f k8s/ --namespace "${ENVIRONMENT}"
echo "Deployment completed"

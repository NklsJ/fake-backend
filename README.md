# Fake Backend for load testing

## Includes
- Docker image for the fake backend server
- Kubernetes deployment and external load balancer to expose the image (you need to setup a cluster separately)

## Server endpoints
- `/json-testing-endpoint`
- `/file-upload-testing-endpoint`

## Configuration
- You will need a kubernetes cluster.

1. Submit the docker image in `./docker` folder to whatever registry you prefer
2. Replace `[IMAGE_SOURCE]` in the file ./kubernetes/load-test-fake-backend-deployment.yaml with your image source
3. Replace `[REPLICAS]` in the file ./kubernetes/load-test-fake-backend-deployment.yaml with the amount of replicas you wish to have
4. Run `kubectl apply -f kubernetes/load-test-fake-backend-deployment.yaml` to create the deployment
5. Run `kubectl apply -f kubernetes/load-test-fake-backend-service.yaml` to expose the deployment via external load balancer

# Horizontal Pod Autoscaler

This is an example of a HorizontalPodAutoscaler. The directory structure is as follows:

* `/api_service` - simple API which can be built as a Docker image
* `/manifests` - manifests for deploying the API to Kubernetes

## Requirements

1. Kubernetes cluster must have [Metrics Server](https://github.com/kubernetes-sigs/metrics-server) installed.

    If using Docker Desktop to run Kubernetes, you can install this with the below command.

    ```bash
    kubectl apply -f ./manifests/components.yaml
    ```

3. Build Docker image for API service.

    ```bash
    docker build ./api_service -t hello-api:latest
    ```

3. Make API Docker image accessible in Kubernetes cluster.

    If using Docker Desktop, the locally-built image should already be available inside your cluster.

    If using [kind](https://kind.sigs.k8s.io/), run the below command:

    ```bash
    kind load docker-image hello-api:latest
    ```

## Deploy application

Run the below command:

```bash
kubectl apply -k ./manifests/app
```

Alternatively, you can uninstall the application with the below command:

```bash
kubectl delete -k ./manifests/app
```

## Run workload & monitor pods

Port forward (or otherwise make accessible) the service `hello-api-service` to your host machine.

Then run a persistent workload with the below script:

```bash
name="world" # feel free to change this to your name :)
while; do
    curl -k "localhost:8080/hello/$name"
    echo ""
    sleep 0.01
done
```

After some time (1-2 minutes), check the number of pods you have. You can also play around with the `sleep` value in the above script along with the different scaling values in `./manifests/app/api_hpa.yaml`.

You can also check the events from the HorizontalPodAutoscaler. It will indicate when and why it happened to scale the deployment.

```bash
kubectl describe hpa/hello-api-hpa
```

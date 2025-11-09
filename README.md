# Micro health check

This is dummy http client that is used for container health checks by docker daemon.
You don't need this thing in k8s as kubelet can do http checks, where docker can only do _exec_.
Also, you don't need this if you already have http client (such as `curl`, `wget`, ...) in your container.

**Known limitations (or rather design decisions)**

- only `GET` method is supported
- no retry mechanism

### Usage

- In Dockerfile

```dockerfile
HEALTHCHECK CMD microhc --url http://127.0.0.1:8080/health
```

- In `docker-compose.yaml`

```yaml
services:
  my-service:
    image: my-oci-reg/my-image:v1.0.0
    healthcheck:
      test: microhc --url http://127.0.0.1:8080/health
```

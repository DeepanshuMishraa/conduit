# Conduit

A lightweight reverse proxy and load balancer built with Go.

Conduit is a learning project focused on understanding how production API gateways work. Instead of relying on managed solutions, the goal is to build the core pieces from scratch and gradually evolve the project into a cloud-native gateway.

## Features

* Reverse proxy with Nginx
* Load balancing across multiple Go instances
* Structured request logging
* Prometheus metrics
* Grafana dashboards
* Docker Compose for local orchestration
* Kubernetes deployment (planned)

## Architecture

```text
                Client
                   │
                   ▼
                 Nginx
                   │
         ┌─────────┼─────────┐
         ▼         ▼         ▼
      API-1     API-2     API-3
                   │
             Prometheus
                   │
                Grafana
```

## Running locally

```bash
docker compose up --build
```

The following services will be available:

| Service     | URL                   |
| ----------- | --------------------- |
| API Gateway | http://localhost      |
| Prometheus  | http://localhost:9090 |
| Grafana     | http://localhost:3000 |

## Project Status

This project is being built incrementally. Current and planned milestones include:

* [x] Multiple Go API instances
* [x] Nginx reverse proxy
* [x] Docker Compose
* [x] Prometheus integration
* [x] Grafana dashboards
* [x] Custom middleware
* [ ] Rate limiting
* [ ] Service discovery
* [ ] Health checks
* [ ] Circuit breaker
* [ ] Kubernetes deployment
* [ ] Horizontal Pod Autoscaler
* [ ] Graceful shutdown


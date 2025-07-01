# Nginx Reverse Proxy with Docker Microservices

A containerized microservices architecture using Nginx as a reverse proxy for Go and Python backend services, orchestrated with Docker Compose.

## Overview

This project demonstrates how to set up a reverse proxy using Nginx to route requests between multiple microservices. The architecture includes two backend services - one built with Go and another with Python - all containerized using Docker.

## Architecture

```
Client Request → Nginx (Port 8080) → Backend Services
                     ├── /service1/* → Go Service (Port 8001)
                     └── /service2/* → Python Service (Port 8002)
```

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Installation & Setup

### For Ubuntu/Linux

```bash
# Install Docker & Docker Compose
sudo apt update
sudo apt install docker.io docker-compose -y
sudo systemctl enable docker
sudo systemctl start docker

# Clone the repository
git clone -b feature/nginx-docker-assignment https://github.com/niranjan-46/nginx-docker-assignment.git
cd nginx-docker-assignment

# Build and run the containers
docker-compose up --build
```

### For Windows (Docker Desktop)

1. Install Docker Desktop and ensure it's running
2. Open PowerShell or WSL and run:

```bash
git clone -b feature/nginx-docker-assignment https://github.com/niranjan-46/nginx-docker-assignment.git
cd nginx-docker-assignment
docker-compose up --build
```

## Usage

Once the services are running, access the application at `http://localhost:8080`.

### Testing the Services

Use the following curl commands to verify routing and health checks:

**Go Service (/service1)**
```bash
curl http://localhost:8080/service1/ping
curl http://localhost:8080/service1/hello
curl http://localhost:8080/service1/health
```

**Python Service (/service2)**
```bash
curl http://localhost:8080/service2/ping
curl http://localhost:8080/service2/hello
curl http://localhost:8080/service2/health
```

## Project Structure

```
.
├── docker-compose.yml         # Docker Compose configuration
├── nginx/
│   ├── nginx.conf             # Nginx reverse proxy rules
│   └── Dockerfile             # Nginx container build
├── service_1/                 # Go microservice
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── service_2/                 # Python microservice
│   ├── app.py
│   ├── requirements.txt
│   └── Dockerfile
└── README.md                  # Project documentation
```

## How It Works

**Nginx Reverse Proxy**: Nginx listens on port 8080 and routes incoming requests to the appropriate backend services based on URL paths.

**Service Routing**:
- `/service1/*` → Go service running on port 8001
- `/service2/*` → Python service running on port 8002

**Service Endpoints**: Each service provides three endpoints:
- `/ping` — Connectivity check
- `/hello` — Test response
- `/health` — Health check endpoint

**Containerization**: All services are containerized using Docker and orchestrated via Docker Compose for easy deployment and scaling.

## Managing Services

**Start services**:
```bash
docker-compose up --build
```

**Start services in background**:
```bash
docker-compose up --build -d
```

**Stop services**:
```bash
docker-compose down
```

**View logs**:
```bash
docker-compose logs
```

## Technology Stack

- **Nginx**: Reverse proxy and load balancer
- **Go**: High-performance backend service
- **Python**: Backend service with Flask/FastAPI
- **Docker**: Containerization platform
- **Docker Compose**: Multi-container orchestration


Troubleshooting Build Issues: 

🩹 Pro tip: Clean up old or broken containers & images
Sometimes Docker Compose tries to reuse corrupted containers or outdated images. To start fresh, run:
```bash
docker-compose down -v --remove-orphans
```
```bash
docker system prune -f
```
Then rebuild and restart the services with:

```bash
docker-compose build --no-cache
```
```bash
docker-compose up
```

## Author

**Niranjan**  
GitHub: [@niranjan-46](https://github.com/niranjan-46)

---

*This project was completed as part of an internship assignment to demonstrate containerized microservices architecture with Nginx reverse proxy.*

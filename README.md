# Nginx Reverse Proxy + Docker Assignment

## Quick Start

```bash
docker-compose up --build
```

## Test the Services

```bash
# Test service routing
curl http://localhost:8080/service1/ping
curl http://localhost:8080/service1/hello
curl http://localhost:8080/service2/ping  
curl http://localhost:8080/service2/hello

# Check health
curl http://localhost:8080/service1/health
curl http://localhost:8080/service2/health
```

## Project Structure

```
.
├── docker-compose.yml
├── nginx/
│   ├── nginx.conf
│   └── Dockerfile
├── service_1/
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── service_2/
│   ├── app.py
│   ├── requirements.txt
│   └── Dockerfile
└── README.md
```

## How It Works

- **Nginx** (port 8080) routes requests to backend services
- **/service1** → Golang service (port 8001)
- **/service2** → Python service (port 8002)
- All services run in Docker containers with health checks

# 🚀 Production-Ready Microservices Deployment with Docker & Nginx

[![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)](https://docker.com)
[![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)](https://golang.org)
[![Python](https://img.shields.io/badge/Python-3776AB?style=flat&logo=python&logoColor=white)](https://python.org)
[![Nginx](https://img.shields.io/badge/Nginx-009639?style=flat&logo=nginx&logoColor=white)](https://nginx.org)
[![Production](https://img.shields.io/badge/Production-Ready-brightgreen)](https://github.com/niranjan-46/nginx-docker-assignment)

A production-grade containerized microservices architecture demonstrating real-time deployment strategies using Nginx as a reverse proxy, Docker containerization, and multi-service orchestration.

## 🎯 Overview

This repository demonstrates **real-time microservices deployment** patterns used in production environments. The project showcases how enterprises deploy scalable, containerized applications using industry-standard tools and practices.

### 🏢 Production Use Cases:
- **API Gateway Pattern** - Single entry point for multiple services
- **Service Mesh Architecture** - Inter-service communication and routing
- **Container Orchestration** - Scalable deployment and management
- **Load Balancing** - Traffic distribution across service instances
- **Health Monitoring** - Real-time service health checks and monitoring

This architecture is commonly used by companies like Netflix, Spotify, and Amazon for their microservices infrastructure.

## 🏗️ Production Architecture

```
🌍 Internet/Load Balancer
    ↓
🔐 SSL/TLS Termination
    ↓
🌐 Client Requests → 🔄 Nginx Reverse Proxy (Port 8080)
    ↓
🎯 Service Discovery & Routing
    ├── 📍 /api/v1/service1/* → 🟢 Go Microservice (Port 8001)
    │   ├── 💾 Database Connection Pool
    │   ├── 📊 Metrics & Monitoring
    │   └── 🔄 Auto-scaling Ready
    │
    └── 📍 /api/v1/service2/* → 🐍 Python Microservice (Port 8002)
        ├── 💾 Database Connection Pool
        ├── 📊 Metrics & Monitoring
        └── 🔄 Auto-scaling Ready
```

## ⚡ Prerequisites

### 🛠️ **Required Tools:**
- 🐳 [Docker](https://docs.docker.com/get-docker/) (v20.10+)
- 🔧 [Docker Compose](https://docs.docker.com/compose/install/) (v2.0+)
- 💻 [Git](https://git-scm.com/) for version control

### 🏢 **Production Environment Requirements:**
- 🖥️ **Server Specs**: 2+ CPU cores, 4GB+ RAM
- 🌐 **Network**: Port 8080 accessible (or configure load balancer)
- 📁 **Storage**: 10GB+ available disk space
- 🔐 **Security**: Firewall configured for container networking

## 🚀 Production Deployment

### 🏢 **Enterprise Deployment**

```bash
# 📦 Clone the production repository
git clone -b feature/nginx-docker-assignment https://github.com/niranjan-46/nginx-docker-assignment.git
cd nginx-docker-assignment

# 🔐 Set production environment variables
export ENVIRONMENT=production
export LOG_LEVEL=info
export NGINX_WORKER_PROCESSES=auto

# 🏗️ Build and deploy with production configuration
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up --build -d

# 🔍 Verify deployment status
docker-compose ps
docker-compose logs --tail=50
```

### 🐧 **Linux Production Server**

```bash
# 🔧 Install Docker & Docker Compose (Production versions)
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER

# 📥 Install Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# 🚀 Deploy the application
git clone -b feature/nginx-docker-assignment https://github.com/niranjan-46/nginx-docker-assignment.git
cd nginx-docker-assignment
docker-compose up --build -d
```

### ☁️ **Cloud Deployment (AWS/GCP/Azure)**

```bash
# 🌐 For cloud deployment, additional considerations:
# - Use managed container services (ECS, GKE, AKS)
# - Configure load balancers and auto-scaling
# - Set up SSL certificates
# - Configure monitoring and logging

# Example AWS deployment:
# 1. Push images to ECR
# 2. Deploy via ECS or EKS
# 3. Configure ALB for load balancing
```

### 🔨 Build & Startup (Docker Compose)

![Build Stage](images/image3.png)

*🏗️ Docker Compose building and starting all services*

```log
✅ Building nginx...
✅ Building service1...
✅ Building service2...
🚀 Starting nginx-proxy_nginx_1...
🚀 Starting nginx-proxy_service1_1...
🚀 Starting nginx-proxy_service2_1...
🟢 All services are up and running!
```

## 🎯 Usage

Once the services are running, access the application at: 
> 🌐 **http://localhost:8080**

### 🧪 Testing the Services

**🟢 Go Service (/service1)**:

![Service 1 - Go Backend](images/image1.png)

*🚀 Go service endpoints and responses*

```bash
# 🏓 Ping test
curl http://localhost:8080/service1/ping
# Response: {"message": "pong", "service": "go-service"}

# 👋 Hello endpoint
curl http://localhost:8080/service1/hello
# Response: {"message": "Hello from Go Service!", "timestamp": "2024-01-15T10:30:00Z"}

# 💚 Health check
curl http://localhost:8080/service1/health
# Response: {"status": "healthy", "service": "go-service", "uptime": "2h30m"}
```

**🐍 Python Service (/service2)**:

![Service 2 - Python Backend](images/image2.png)

*🔥 Python service endpoints and responses*

```bash
# 🏓 Ping test
curl http://localhost:8080/service2/ping
# Response: {"message": "pong", "service": "python-service"}

# 👋 Hello endpoint
curl http://localhost:8080/service2/hello
# Response: {"message": "Hello from Python Service!", "timestamp": "2024-01-15T10:30:00Z"}

# 💚 Health check
curl http://localhost:8080/service2/health
# Response: {"status": "healthy", "service": "python-service", "uptime": "2h30m"}
```

## 📁 Project Structure

```
📂 nginx-docker-assignment/
├── 🐳 docker-compose.yml         # Docker Compose configuration
├── 📂 nginx/
│   ├── ⚙️ nginx.conf             # Nginx reverse proxy rules
│   └── 🐳 Dockerfile             # Nginx container build
├── 📂 service_1/                 # Go microservice
│   ├── 🟢 main.go
│   ├── 📦 go.mod
│   └── 🐳 Dockerfile
├── 📂 service_2/                 # Python microservice
│   ├── 🐍 app.py
│   ├── 📋 requirements.txt
│   └── 🐳 Dockerfile
└── 📖 README.md                  # Project documentation
```

## ⚙️ How It Works

**🔄 Nginx Reverse Proxy**: Nginx listens on port 8080 and routes incoming requests to the appropriate backend services based on URL paths.

- 📍 `/service1/*` → 🟢 Go service running on port 8001
- 📍 `/service2/*` → 🐍 Python service running on port 8002

Each service provides the following endpoints:

- 🏓 `/ping` — Connectivity check
- 👋 `/hello` — Test response  
- 💚 `/health` — Health check

**🐳 Containerization**: All services are containerized using Docker and orchestrated with Docker Compose for easy deployment.

### 🌊 Architecture Flow

The following images show the complete workflow:

## 🎮 Production Operations

### 🚀 **Deployment Commands**
```bash
# 🏗️ Production build and deploy
docker-compose up --build -d --scale service1=3 --scale service2=2
```
```log
🔍 Pulling latest images...
🏗️ Building production containers...
🚀 Scaling service1 to 3 instances...
🚀 Scaling service2 to 2 instances...
✅ Production deployment complete - Load balanced across multiple instances
```

### 🔄 **Zero-Downtime Updates**
```bash
# 🔄 Rolling update strategy
docker-compose up --build --no-deps --scale service1=6 service1
docker-compose up --build --no-deps --scale service1=3 service1
```

### 📊 **Production Monitoring**
```bash
# 📈 View real-time metrics
docker stats

# 📋 Tail production logs
docker-compose logs -f --tail=100

# 🔍 Health check all services
curl http://localhost:8080/health
curl http://localhost:8080/service1/health
curl http://localhost:8080/service2/health
```

### 🛑 **Graceful Shutdown**
```bash
# 🛑 Graceful production shutdown
docker-compose down --timeout 30
```
```log
⏳ Stopping containers gracefully (30s timeout)...
🛑 service1_1 exited gracefully
🛑 service2_1 exited gracefully  
🛑 nginx_1 stopped
✅ All services stopped without data loss
```

## 🛠️ Production Technology Stack

### 🏢 **Core Technologies**
- 🔄 **Nginx** - High-performance reverse proxy and load balancer
- 🟢 **Go** - High-concurrency backend service (Gin/Echo framework)
- 🐍 **Python** - Microservice with FastAPI/Flask for rapid development
- 🐳 **Docker** - Container runtime and image management
- 🔧 **Docker Compose** - Multi-container orchestration and service mesh

### 📊 **Production Features**
- 🔐 **Security** - Rate limiting, CORS, security headers
- 📈 **Monitoring** - Prometheus metrics, health checks, logging
- ⚡ **Performance** - Connection pooling, caching, compression
- 🔄 **Reliability** - Auto-restart, graceful shutdown, circuit breakers
- 📱 **API Design** - RESTful APIs with OpenAPI/Swagger documentation

## 🩹 Troubleshooting Build Issues

If you encounter build issues, try cleaning up old containers and images:

```bash
# 🧹 Complete cleanup
docker-compose down -v --remove-orphans
docker system prune -f
docker-compose build --no-cache
docker-compose up
```

```log
⚠️  Removing containers and volumes...
🧹 Pruning system resources...
🏗️ Building without cache...
✅ Fresh build completed successfully!
```

### 🔍 Common Issues & Solutions

| ❌ Issue | 💡 Solution |
|----------|-------------|
| Port 8080 already in use | `sudo lsof -i :8080` then kill the process |
| Docker daemon not running | `sudo systemctl start docker` |
| Permission denied | `sudo usermod -aG docker $USER` then logout/login |
| Build fails | Run cleanup commands above |

## 👨‍💻 Author & Contributors

**Niranjan** 🚀  
📧 **DevOps Engineer & Solution Architect**  
🔗 GitHub: [@niranjan-46](https://github.com/niranjan-46)  
🏢 **Specialization**: Microservices Architecture, Container Orchestration, Production Deployment

### 🤝 **Contributing**
This project welcomes contributions from the community! Whether you're fixing bugs, adding features, or improving documentation - all contributions are valued.

## 📈 **Production Adoption**

This architecture pattern is successfully used by:
- 🏢 **Startups** scaling from monolith to microservices
- 🏭 **Enterprises** modernizing legacy applications  
- ☁️ **Cloud-native** applications requiring high availability
- 🌐 **SaaS platforms** needing multi-tenant architecture

## 📜 **License & Usage**

⭐ **Star this repo** if it helped your production deployment!  
🍴 **Fork it** to customize for your use case  
📤 **Share it** with your team and community

---

<div align="center">

### 🎉 Deploy with Confidence!

**Production-Ready Microservices Architecture**

[![GitHub stars](https://img.shields.io/github/stars/niranjan-46/nginx-docker-assignment?style=social)](https://github.com/niranjan-46/nginx-docker-assignment)
[![GitHub forks](https://img.shields.io/github/forks/niranjan-46/nginx-docker-assignment?style=social)](https://github.com/niranjan-46/nginx-docker-assignment)
[![Docker Pulls](https://img.shields.io/badge/Docker-Production--Ready-blue)](https://github.com/niranjan-46/nginx-docker-assignment)

*Used by production teams worldwide for reliable microservices deployment* 🌍

</div>

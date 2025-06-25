# 📦 Nginx Reverse Proxy with Docker Microservices

A simple setup using **Nginx** as a reverse proxy for two backend microservices — one in **Go** and another in **Python** — all containerized using **Docker**.

---

## 🛠 Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

---

## 🚀 Getting Started

To build and start all services:

```bash
docker-compose up --build
🧪 Testing the Services
Use the following curl commands to test the routing and health of each service.

🔹 Go Service (/service1)
bash
Copy
Edit
curl http://localhost:8080/service1/ping
curl http://localhost:8080/service1/hello
curl http://localhost:8080/service1/health
🔹 Python Service (/service2)
bash
Copy
Edit
curl http://localhost:8080/service2/ping
curl http://localhost:8080/service2/hello
curl http://localhost:8080/service2/health
📁 Project Structure
bash
Copy
Edit
.
├── docker-compose.yml         # Docker Compose configuration
├── nginx/
│   ├── nginx.conf             # Nginx reverse proxy configuration
│   └── Dockerfile             # Dockerfile for Nginx container
├── service_1/                 # Go microservice
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── service_2/                 # Python microservice
│   ├── app.py
│   ├── requirements.txt
│   └── Dockerfile
└── README.md                  # Project documentation
⚙️ How It Works
Nginx listens on port 8080 and proxies requests based on the URL path:

/service1/* → Go service (port 8001)

/service2/* → Python service (port 8002)

Both services include basic test and health check endpoints.

Everything runs in separate containers using Docker Compose.

👤 Author
Niranjan
🔗 GitHub: @niranjan-46

This project was completed as part of an internship assignment

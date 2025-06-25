🔁 Nginx Reverse Proxy + Docker Assignment
This project demonstrates the use of Nginx as a reverse proxy to route traffic to two backend microservices running in Docker containers — one written in Go, the other in Python.

🚀 Quick Start
bash
Copy
Edit
docker-compose up --build
🔍 Test the Services
Use the following commands to verify routing and service health:

bash
Copy
Edit
# Service 1 (Go)
curl http://localhost:8080/service1/ping
curl http://localhost:8080/service1/hello
curl http://localhost:8080/service1/health

# Service 2 (Python)
curl http://localhost:8080/service2/ping
curl http://localhost:8080/service2/hello
curl http://localhost:8080/service2/health
📁 Project Structure
bash
Copy
Edit
.
├── docker-compose.yml         # Docker Compose setup
├── nginx/
│   ├── nginx.conf             # Nginx configuration
│   └── Dockerfile            # Nginx Docker image
├── service_1/                 # Go service
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── service_2/                 # Python service
│   ├── app.py
│   ├── requirements.txt
│   └── Dockerfile
└── README.md
⚙️ How It Works
Nginx listens on port 8080 and proxies traffic to:

/service1/* → Go service (port 8001)

/service2/* → Python service (port 8002)

Each service container includes basic endpoints (/ping, /hello, /health) for testing.

Health checks are defined and handled within each container.

All services are containerized and orchestrated using Docker Compose.

🙋‍♂️ Author
Made with 💻 by Niranjan

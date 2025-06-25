📦 Nginx Reverse Proxy with Docker Microservices
A simple setup using Nginx as a reverse proxy for two backend services — one in Go, the other in Python — all running in Docker containers.

🏁 Prerequisites
Docker

Docker Compose

🚀 Getting Started
To build and run all services, execute:

bash
Copy
Edit
docker-compose up --build
🧪 Testing the Services
You can test routing and health endpoints using curl:

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
🗂️ Project Structure
bash
Copy
Edit
.
├── docker-compose.yml         # Docker Compose configuration
├── nginx/
│   ├── nginx.conf             # Nginx reverse proxy rules
│   └── Dockerfile             # Nginx Dockerfile
├── service_1/                 # Go (Golang) microservice
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── service_2/                 # Python microservice
│   ├── app.py
│   ├── requirements.txt
│   └── Dockerfile
└── README.md                  # Project documentation
⚙️ How It Works
Nginx listens on port 8080 and proxies requests to:

/service1/* → Go service (port 8001)

/service2/* → Python service (port 8002)

Each microservice runs in its own container.

Health checks and basic test routes are included in both services.

All components are orchestrated using Docker Compose.

👤 Author
Niranjan
GitHub: @niranjan-46
This project was completed as part of an internship assignment.

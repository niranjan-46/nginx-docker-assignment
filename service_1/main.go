package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	printBanner()

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	serviceVersion := os.Getenv("SERVICE_VERSION")
	if serviceVersion == "" {
		serviceVersion = "v1.0.0"
	}

	log.Printf("📦 Container Hostname: %s", hostname)
	log.Printf("🌐 Service Version: %s", serviceVersion)

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/info", infoHandler)
	mux.HandleFunc("/nginx-test", nginxTestHandler)

	server := &http.Server{
		Addr:    ":8001",
		Handler: mux,
	}

	go func() {
		log.Println("🚀 Service 1 listening on port 8001...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server failed: %v", err)
		}
	}()

	// Graceful shutdown on SIGINT/SIGTERM
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("⚙️ Shutting down gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}
	log.Println("✅ Server exited cleanly")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>🚀 Cloud Service 1</title>
			<style>
				body {
					margin: 0;
					padding: 0;
					height: 100vh;
					background: linear-gradient(135deg, #1e3c72, #2a5298);
					overflow: hidden;
					display: flex;
					justify-content: center;
					align-items: center;
					flex-direction: column;
					font-family: 'Segoe UI', sans-serif;
					color: #fff;
				}
				.animated-text {
					font-size: 3em;
					background: linear-gradient(90deg, #00c6ff, #0072ff);
					background-size: 200% 200%;
					-webkit-background-clip: text;
					-webkit-text-fill-color: transparent;
					animation: gradientMove 5s ease infinite;
					text-align: center;
					margin: 0 20px;
				}
				.cloud {
					position: absolute;
					width: 200px;
					height: 60px;
					background: rgba(255, 255, 255, 0.2);
					border-radius: 100px;
					top: 30%;
					left: -200px;
					animation: moveCloud 60s linear infinite;
				}
				.cloud::before, .cloud::after {
					content: "";
					position: absolute;
					background: inherit;
					width: 100px;
					height: 100px;
					border-radius: 50%;
				}
				.cloud::before {
					top: -30px;
					left: 10px;
				}
				.cloud::after {
					top: -20px;
					right: 10px;
				}
				@keyframes moveCloud {
					0% { transform: translateX(0); }
					100% { transform: translateX(150vw); }
				}
				@keyframes gradientMove {
					0% {background-position: 0% 50%;}
					50% {background-position: 100% 50%;}
					100% {background-position: 0% 50%;}
				}
				p {
					font-size: 1.2em;
					margin-top: 30px;
				}
				code {
					background: rgba(255, 255, 255, 0.1);
					padding: 5px 10px;
					border-radius: 5px;
				}
			</style>
		</head>
		<body>
			<div class="cloud"></div>
			<div class="animated-text">🚀 Welcome to Cloud Service 1</div>
			<p>Explore endpoints: <code>/ping</code>, <code>/hello</code>, <code>/health</code>, <code>/info</code>, <code>/nginx-test</code></p>
		</body>
		</html>
	`)
	logRequest(r)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	jsonResponse(w, map[string]string{
		"status":  "ok",
		"service": "1",
	})
	logRequest(r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	jsonResponse(w, map[string]string{
		"message": "Hello from Service 1",
	})
	logRequest(r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	jsonResponse(w, map[string]string{
		"status": "healthy",
	})
	logRequest(r)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	hostname, _ := os.Hostname()
	serviceVersion := os.Getenv("SERVICE_VERSION")
	if serviceVersion == "" {
		serviceVersion = "v1.0.0"
	}
	jsonResponse(w, map[string]string{
		"hostname":        hostname,
		"service_version": serviceVersion,
		"message":         "Cloud & Docker info for debugging",
	})
	logRequest(r)
}

func nginxTestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	jsonResponse(w, map[string]string{
		"message": "This response verifies Nginx is proxying correctly!",
	})
	logRequest(r)
}

func jsonResponse(w http.ResponseWriter, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("❌ Failed to encode JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func logRequest(r *http.Request) {
	log.Printf("📥 [%s] %s from %s (UA: %s)", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
}

func printBanner() {
	fmt.Println(`
  ____                 _          ____  
 / ___| ___   ___   __| | ___    / ___| 
| |  _ / _ \ / _ \ / _  |/ _ \  | |     
| |_| | (_) | (_) | (_| |  __/  | |___  
 \____|\___/ \___/ \__,_|\___|   \____| 
                                       
🚀 Service 1 starting up...
`)
}


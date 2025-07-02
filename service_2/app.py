from flask import Flask, jsonify, request, Response
import logging
import os

app = Flask(__name__)

# --- Colored log formatting ---
class ColorFormatter(logging.Formatter):
    COLORS = {
        'DEBUG': "\033[0;36m",
        'INFO': "\033[0;32m",
        'WARNING': "\033[0;33m",
        'ERROR': "\033[0;31m",
        'CRITICAL': "\033[1;41m",
    }
    RESET = "\033[0m"

    def format(self, record):
        color = self.COLORS.get(record.levelname, "")
        message = super().format(record)
        return f"{color}{message}{self.RESET}"

handler = logging.StreamHandler()
handler.setFormatter(ColorFormatter("[%(asctime)s] [%(levelname)s] %(message)s"))
app.logger.addHandler(handler)
app.logger.setLevel(logging.INFO)

# --- Endpoints ---
@app.route("/")
def index():
    app.logger.info("🌐 / accessed from %s", request.remote_addr)
    html_content = """
    <!DOCTYPE html>
    <html>
    <head>
        <title>🚀 Welcome to Service 2</title>
        <style>
            body {
                margin: 0;
                height: 100vh;
                display: flex;
                justify-content: center;
                align-items: center;
                flex-direction: column;
                font-family: sans-serif;
                background: linear-gradient(135deg, #141e30, #243b55);
                color: white;
                overflow: hidden;
            }
            h1 {
                font-size: 3em;
                background: linear-gradient(90deg, #ff5f6d, #ffc371);
                -webkit-background-clip: text;
                -webkit-text-fill-color: transparent;
                animation: gradient 4s ease infinite;
            }
            p {
                font-size: 1.2em;
                margin-top: 20px;
            }
            @keyframes gradient {
                0% {background-position: 0%;}
                50% {background-position: 100%;}
                100% {background-position: 0%;}
            }
            .cloud {
                position: absolute;
                width: 180px;
                height: 50px;
                background: rgba(255, 255, 255, 0.2);
                border-radius: 100px;
                top: 40%;
                left: -200px;
                animation: moveCloud 50s linear infinite;
            }
            .cloud::before, .cloud::after {
                content: "";
                position: absolute;
                background: inherit;
                border-radius: 50%;
            }
            .cloud::before { width: 80px; height: 80px; top: -20px; left: 20px; }
            .cloud::after { width: 60px; height: 60px; top: -15px; right: 20px; }
            @keyframes moveCloud {
                0% { transform: translateX(0); }
                100% { transform: translateX(140vw); }
            }
        </style>
    </head>
    <body>
        <div class="cloud"></div>
        <h1>☁️ Service 2 is Alive!</h1>
        <p>Try <code>/ping</code>, <code>/hello</code>, or <code>/health</code> endpoints!</p>
    </body>
    </html>
    """
    return Response(html_content, mimetype="text/html")

@app.route("/ping")
def ping():
    app.logger.info("📡 /ping accessed from %s", request.remote_addr)
    return jsonify(status="ok", service="2")

@app.route("/hello")
def hello():
    app.logger.info("👋 /hello accessed from %s", request.remote_addr)
    return jsonify(message="Hello from Service 2")

@app.route("/health")
def health():
    app.logger.info("❤️ /health accessed from %s", request.remote_addr)
    return jsonify(status="healthy")

if __name__ == "__main__":
    host = "0.0.0.0"
    port = int(os.getenv("PORT", 8002))
    app.logger.info("🚀 Starting Service 2 on http://%s:%d", host, port)
    app.run(host=host, port=port)


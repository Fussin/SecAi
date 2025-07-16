#!/bin/bash
# AUTONOMOUSPEN AI - Complete Setup Script
# Save this as setup_autonomouspen.sh and run it

echo "🚀 Setting up AUTONOMOUSPEN AI Project..."

# Create all directories
mkdir -p autonomouspen-ai/{cmd/{scanner,api,worker},internal/{core,recon,scanner,validator,ai,reporter},pkg/{bugbounty,dedup,evidence},python/{ai_models,payloads,learning},web/{dashboard,api},docker/{scanner,ai,infrastructure},configs/{programs,payloads,rules},prompts/{recon,exploit,report},tests/{unit,integration,targets}}

cd autonomouspen-ai

# Create README.md
cat > README.md << 'EOF'
# AUTONOMOUSPEN AI

Fully autonomous penetration testing platform with zero human intervention.

## Features
- 100% Autonomous operation
- AI-driven vulnerability discovery
- Automated PoC generation
- Bug bounty platform integration
- Scales to 1000+ concurrent targets

## Quick Start
```bash
docker-compose up -d
EOF

# Create .gitignore
cat > .gitignore << 'EOF'
# Binaries
*.exe
*.dll
*.so
*.dylib
bin/

# Test binary
*.test

# Output
*.out

# Go vendor
vendor/

# Python
pycache/
*.py[cod]
*$py.class
venv/
env/

# Environment
.env
.env.local

# IDE
.vscode/
.idea/
*.swp
*.swo

# Logs
*.log
logs/

# Docker
docker-compose.override.yml
EOF

# Create go.mod
cat > go.mod << 'EOF'
module github.com/autonomouspen/scanner

go 1.21

require (
github.com/gin-gonic/gin v1.9.1
github.com/go-redis/redis/v8 v8.11.5
github.com/lib/pq v1.10.9
)
EOF

# Create requirements.txt
cat > requirements.txt << 'EOF'
langchain==0.1.0
openai==1.0.0
anthropic==0.8.0
playwright==1.40.0
celery==5.3.0
redis==5.0.0
fastapi==0.104.0
uvicorn==0.24.0
transformers==4.35.0
sqlalchemy==2.0.0
psycopg2-binary==2.9.9
pydantic==2.5.0
httpx==0.25.0
beautifulsoup4==4.12.0
Pillow==10.1.0
imagehash==4.3.1
simhash==2.1.2
EOF

# Create docker-compose.yml
cat > docker-compose.yml << 'EOF'
version: '3.8'

services:
  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_DB: autonomouspen
      POSTGRES_USER: autonomouspen
      POSTGRES_PASSWORD: ${DB_PASSWORD:-secretpassword}
    volumes:
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U autonomouspen"]
      interval: 10s
      timeout: 5s
      retries: 5

  redis:
    image: redis:7-alpine
    command: redis-server --appendonly yes --requirepass ${REDIS_PASSWORD:-redispassword}
    volumes:
      - redis_data:/data
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 10s
      timeout: 5s
      retries: 5

  scanner:
    build:
      context: .
      dockerfile: docker/scanner/Dockerfile
    deploy:
      replicas: 3
    environment:
      - DATABASE_URL=postgresql://autonomouspen:${DB_PASSWORD:-secretpassword}@postgres/autonomouspen
      - REDIS_URL=redis://:${REDIS_PASSWORD:-redispassword}@redis:6379
    depends_on:
      postgres:
        condition: service_healthy
      redis:
        condition: service_healthy

  ai_engine:
    build:
      context: .
      dockerfile: docker/ai/Dockerfile
    environment:
      - OPENAI_API_KEY=${OPENAI_API_KEY}
      - ANTHROPIC_API_KEY=${ANTHROPIC_API_KEY}
      - DATABASE_URL=postgresql://autonomouspen:${DB_PASSWORD:-secretpassword}@postgres/autonomouspen
      - REDIS_URL=redis://:${REDIS_PASSWORD:-redispassword}@redis:6379
    depends_on:
      - postgres
      - redis

volumes:
  postgres_data:
  redis_data:
EOF

# Create main scanner Go file
mkdir -p cmd/scanner
cat > cmd/scanner/main.go << 'EOF'
package main

import (
"fmt"
"log"
"os"
)

func main() {
fmt.Println("🚀 AUTONOMOUSPEN AI Scanner Starting...")

// Initialize configuration
config := loadConfig()

// Start autonomous scanning loop
log.Println("Scanner initialized. Beginning autonomous operation...")

// This will run forever
select {}
}

func loadConfig() map[string]string {
return map[string]string{
"mode": "autonomous",
"targets_per_hour": "1000",
}
}
EOF

# Create Python AI Decision Engine
mkdir -p python/ai_models
cat > python/ai_models/decision_engine.py << 'EOF'
"""
AUTONOMOUSPEN AI - Decision Engine
Fully autonomous penetration testing brain
"""

from typing import List, Dict, Any
import json
from langchain import OpenAI, LLMChain
from langchain.prompts import PromptTemplate

class AutonomousDecisionEngine:
  """Core AI brain for autonomous penetration testing"""

  def __init__(self):
      self.llm = OpenAI(temperature=0.7)
      self.scan_planner = self._create_scan_planner()
      self.payload_generator = self._create_payload_generator()

  def _create_scan_planner(self) -> LLMChain:
      """Creates the AI chain for planning scans"""
      prompt = PromptTemplate(
          input_variables=["recon_data", "technology_stack"],
          template="""
You are AUTONOMOUSPEN AI, an expert autonomous penetration tester.

Reconnaissance Data:
{recon_data}

Technology Stack:
{technology_stack}

Create a prioritized testing plan with:

* vulnerability_priorities (ordered by impact)
* test_sequences (specific attack chains)
* time_allocation (minutes per test)
* bypass_strategies (for detected protections)
Output as JSON.
"""
      )
      return LLMChain(llm=self.llm, prompt=prompt)

  def _create_payload_generator(self) -> LLMChain:
      """Creates the AI chain for generating payloads"""
      prompt = PromptTemplate(
          input_variables=["vulnerability_type", "context", "waf_detected"],
          template="""
Generate advanced payloads for {vulnerability_type}.

Context: {context}
WAF Detected: {waf_detected}

Create 10 payloads with increasing sophistication and evasion techniques.
Include polyglots and encoding variations.

Output as JSON array.
"""
      )
      return LLMChain(llm=self.llm, prompt=prompt)

  def plan_scan(self, target: Dict[str, Any]) -> Dict[str, Any]:
      """Plans the scanning strategy for a target"""
      result = self.scan_planner.run(
          recon_data=json.dumps(target.get("recon", {})),
          technology_stack=target.get("tech_stack", "Unknown")
      )
      return json.loads(result)

  def generate_payloads(self, vuln_type: str, context: str, waf: bool = False) -> List[str]:
      """Generates context-aware payloads"""
      result = self.payload_generator.run(
          vulnerability_type=vuln_type,
          context=context,
          waf_detected=str(waf)
      )
      return json.loads(result)

if __name__ == "__main__":
  # Test the engine
  engine = AutonomousDecisionEngine()
  print("✅ AUTONOMOUSPEN AI Decision Engine initialized")
EOF

# Create XSS Scanner Module
mkdir -p internal/scanner
cat > internal/scanner/xss.go << 'EOF'
package scanner

import (
"fmt"
"log"
)

type XSSScanner struct {
Name string
}

func NewXSSScanner() *XSSScanner {
return &XSSScanner{
Name: "XSS Scanner Module",
}
}

func (x *XSSScanner) Scan(target string) {
log.Printf("Scanning %s for XSS vulnerabilities...", target)
// Implementation will go here
}
EOF

# Create Makefile
cat > Makefile << 'EOF'
.PHONY: help build run test clean

help:
@echo "AUTONOMOUSPEN AI - Makefile Commands"
@echo "  make build - Build all containers"
@echo "  make run   - Run the complete system"
@echo "  make test  - Run all tests"
@echo "  make clean - Clean build artifacts"

build:
docker-compose build

run:
docker-compose up -d

test:
go test ./...
python -m pytest python/

clean:
docker-compose down -v
rm -rf bin/
EOF

# Create Scanner Dockerfile
mkdir -p docker/scanner
cat > docker/scanner/Dockerfile << 'EOF'
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o scanner cmd/scanner/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/scanner .
CMD ["./scanner"]
EOF

# Create AI Dockerfile
mkdir -p docker/ai
cat > docker/ai/Dockerfile << 'EOF'
FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY python/ ./python/
ENV PYTHONPATH=/app

CMD ["python", "-m", "uvicorn", "python.api.main:app", "--host", "0.0.0.0", "--port", "8000"]
EOF

# Create .env.example
cat > .env.example << 'EOF'
# AUTONOMOUSPEN AI Configuration
# AI API Keys
OPENAI_API_KEY=your_openai_key_here
ANTHROPIC_API_KEY=your_anthropic_key_here

# Bug Bounty Platform Keys
HACKERONE_API_KEY=your_hackerone_key_here
BUGCROWD_API_KEY=your_bugcrowd_key_here
INTIGRITI_API_KEY=your_intigriti_key_here

# Database & Redis
DB_PASSWORD=secretpassword
REDIS_PASSWORD=redispassword
EOF

echo "✅ AUTONOMOUSPEN AI Project setup complete!"
echo "➡️ Next steps:"
echo "1. cd autonomouspen-ai"
echo "2. Edit .env.example and save as .env with your API keys"
echo "3. Run 'make build' then 'make run'"

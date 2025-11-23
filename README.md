# Feed Management

A high-performance Go microservice designed to deliver a dynamic and personalized social media feed. It processes user relationships and content interactions in real-time using an event-driven pipeline, ensuring fast and scalable feed generation.

## Tech Stack

- **Language:** Go 1.21+
- **Database:** Mongodb
- **Cache:** Redis
- **Message Broker:** LavinMQ / RabbitMQ
- **Containerization:** Docker & Docker Compose
- **Orchestration:** Kubernetes

## Prerequisites

- Go 1.21 or higher
- Docker & Docker Compose
- Mongodb
- Redis
- LavinMQ or RabbitMQ

## Quick Start

### 1. Clone the Repository

```bash
git https://github.com/nurmanhadi/go-feed-management.git
cd feed-management
```

### 2. Configure Environment Variables

```bash
cp .env.example .env
```

Edit `.env` with your configuration:

```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=27017
DB_USERNAME=root
DB_PASSWORD=root
DB_NAME=feed-management

# Cache Configuration
CACHE_HOST=localhost
CACHE_PORT=6379

# Message Broker Configuration
BROKER_HOST=localhost
BROKER_PORT=5672
BROKER_USERNAME=guest
BROKER_PASSWORD=guest
BROKER_VHOST=someone

# api user
API_USER=http://localhost:3001
```

### 3. Start Services with Docker Compose

```bash
docker-compose up -d
```

### 4. Run Database Migrations

```bash
make migrate-up
```

## Development

### Run Locally

```bash
go run cmd/main.go
```

### Build Binary

```bash
go build -o bin/feed-service cmd/main.go
```

## Api Documentation
[Here!](docs/api.md)

## Security Considerations

- Enable TLS/SSL for production deployments
- Use HTTPS for all API endpoints
- Validate and sanitize all user inputs
- Implement rate limiting on authentication endpoints
- Use environment variables for sensitive configuration (API keys, database credentials)
- Implement proper authentication and authorization mechanisms
- Enable CORS only for trusted domains in production
- Keep dependencies updated regularly

## License

This project is licensed under the MIT License.

## Author

**Nurman Hadi**  
Backend Developer (Golang, Microservices)  
GitHub: [@nurmanhadi](https://github.com/nurmanhadi)
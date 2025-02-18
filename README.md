
<div align="center">
<picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://github.com/user-attachments/assets/86f56636-9836-44d7-b3ab-b6a43fb1b4ec">
    <img alt="Permify logo" src="https://github.com/user-attachments/assets/86f56636-9836-44d7-b3ab-b6a43fb1b4ec" width="200">
</picture>

<h1 align="center">
   GoSkeleton
</h1>
</div>
<p align="center">
GoSkeleton is a structured and scalable Go project template with built-in support for <b>gRPC, HTTP, Swagger, telemetry,
circuit breakers, rate limiting</b>, and <b>export utilities</b>. It provides a clean foundation for developing
production-ready applications.
</p>

<p align="center">
    <a href="https://github.com/tolgaOzen/go-skeleton" target="_blank"><img src="https://img.shields.io/github/go-mod/go-version/tolgaOzen/go-skeleton?style=for-the-badge&logo=go" alt="Skeleton Go Version" /></a>&nbsp;
    <a href="https://goreportcard.com/report/github.com/tolgaOzen/go-skeleton" target="_blank"><img src="https://goreportcard.com/badge/github.com/tolgaOzen/go-skeleton?style=for-the-badge&logo=go" alt="Skeleton Go Report Card" /></a>&nbsp;
    <a href="https://github.com/tolgaOzen/go-skeleton" target="_blank"><img src="https://img.shields.io/github/license/tolgaOzen/go-skeleton?style=for-the-badge" alt="Skeleton Licence" /></a>&nbsp;
    <a href="https://img.shields.io/github/commit-activity/m/tolgaOzen/go-skeleton?style=for-the-badge" target="_blank"><img src="https://img.shields.io/github/commit-activity/m/tolgaOzen/go-skeleton?style=for-the-badge" alt="Skeleton Commit Activity" /></a>&nbsp;
</p>  

## Features

- **gRPC & HTTP Support** – Fully integrated gRPC and RESTful HTTP APIs.
- **Swagger API Documentation** – Automatically generated API docs with Swagger UI.
- **Telemetry (Tracing & Metrics)** – OpenTelemetry support for monitoring and observability.
- **Circuit Breaker** – Built-in resilience patterns to handle failures gracefully.
- **Rate Limiting** – Prevents abuse with request throttling mechanisms.
- **Export Utilities** – Includes Makefile commands for running and managing the application.
- **Docker Support** – Easily deployable using `docker-compose`.
- **Security & Code Quality** – Integrated linters and security scanners.

---

## Getting Started

### Prerequisites

- **Go 1.16+**
- **`buf`** (for Protocol Buffers)
- **Docker** (for containerized deployments)
- **Make** (for running commands efficiently)

---

## Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/tolgaOzen/go-skeleton.git
   cd go-skeleton
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Run the application**

   ```bash
   go run cmd/skeleton/skeleton.go
   ```

4. **Run with Docker**

   ```bash
   docker-compose up --build
   ```

---

## Project Structure

```
go-skeleton/
├── cmd/skeleton       # Main application entry point
├── config             # Configuration files
├── internal           # Private application logic and business rules
├── pkg               # Reusable utilities and helpers
├── proto             # gRPC Protocol Buffers definitions
├── docs              # Swagger documentation & OpenAPI specs
├── tools             # Developer tools (linting, security scans, etc.)
└── Makefile          # Helpful automation commands
```

---

## API Documentation

Swagger UI is available at:

```
http://localhost:8080/swagger/index.html
```

---

## Observability

- **Tracing** – OpenTelemetry integration for distributed tracing.
- **Metrics** – Prometheus-compatible metrics for monitoring.

---

## Makefile Commands

GoSkeleton includes **export utilities** via Makefile to streamline development.

### Usage

```bash
make help
```

### Available Commands

| Command              | Description                                     |
|----------------------|-------------------------------------------------|
| `make build`         | Build the Go application                        |
| `make format`        | Format code using `gofumpt`                     |
| `make lint-all`      | Run all linters (`golangci-lint`, `hadolint`)   |
| `make security-scan` | Scan for security vulnerabilities using `gosec` |
| `make coverage`      | Generate global code coverage report            |
| `make clean`         | Remove temporary and generated files            |
| `make release`       | Prepare for release (format, scan, clean)       |
| `make serve`         | Run the compiled Go application                 |
| `make compose-up`    | Start the application using `docker-compose`    |
| `make compose-down`  | Stop all `docker-compose` services              |

---

## License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

For more details, visit the **[GoSkeleton GitHub repository](https://github.com/tolgaOzen/go-skeleton)**.
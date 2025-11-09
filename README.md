# 🚀 GraphQL Server

## ⚙️ `.gqlgen.yml` Configuration

```yml
schema:
  - internal/graphql/schema/**/*.graphql

exec:
  filename: internal/graphql/generated/generated.go
  package: generated

model:
  filename: internal/graphql/generated/models_gen.go
  package: generated

resolver:
  layout: follow-schema
  dir: internal/graphql/resolvers
  package: resolvers

autobind: []
```

## 📦 Local Setup

1. Copy the example config file

```bash
cp dev.example.toml dev.toml
```

### These are the variables that we use

```toml
[server]
server_port = ":5000"
origins = ["http://localhost:3000", "https://example.com"]

[database]
db_host = "localhost"
db_port = "5432"
db_user = "postgres"
db_password = "sahil"
db_name = "graphql"

[jwt]
signing_key = "K3#v@9$1!pZ^mL2&uQ7*rF4)gT8_W+oB"
encryption_key = "8x@R!5#N0$kQ7_vT3&Wp2+Z^fC6*bM1h"
```

2. Download dependencies

```bash
go mod download
```

3. Run the server locally

```bash
go run cmd/server/main.go
```

---

## 🐳 Run with Docker (Standard Server)

1. Build the Docker image

```bash
docker build -f Docker/Dockerfile.server -t graphql-server .
```

2. Run the container

```bash
docker run --name graphql-server -p 5000:5000 graphql-server
```

---

## ☁️ Run with Docker (Serverless - AWS Lambda Compatible)

1. Build the Lambda-style image

```bash
docker build -f Docker/Dockerfile.serverless -t graphql-serverless .
```

2. Run the container locally with Lambda Runtime Interface Emulator (RIE)

```bash
docker run --rm -p 9000:8080 graphql-serverless
```

---

## 🛠️ Generate GraphQL Schema and Code

1. Install gqlgen CLI

```bash
go install github.com/99designs/gqlgen@latest
```

2. Generate GraphQL schema and queries

```bash
gqlgen generate
```

---

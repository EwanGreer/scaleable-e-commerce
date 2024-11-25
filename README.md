# README

A scale-able e-commerce API

## Running

This project required `Docker` to run.

To run all services in docker, including Kafka etc.

```bash
make all
```

To run an individual service, and its deps

```bash
docker compose up -d <service_name>
```

## Services

- User Service
- Product Catalog Service
- Shopping Cart Service
- Order Service
- Payment Service
- Notification Service

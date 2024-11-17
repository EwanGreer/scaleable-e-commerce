# README

A scaleable e-commerce API

## Running

To run all services in docker, including kafka etc.

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

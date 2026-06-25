# Financial Planner Go

Cleaned and secured Financial Planner service.

## What is corrected

- Removed Policy Planner external API integration.
- Removed Policy Planner configuration and Resty usage.
- Fixed controller-service-repository imports.
- Added request validation.
- Added security headers.
- Added 1 MB request body limit.
- Added server timeouts.
- Added DB connection pool settings.
- Masked email and phone in API response.
- Hidden raw JSON payload and error message from API response.
- Added `/health` endpoint.

## Run

```bash
go mod tidy
go run main.go
```

## Migration

Migration runs automatically through `AutoMigrate`.

Manual migration:

```bash
go run migrate/migrate.go
```

## API

```http
POST /finance/financial-planner/report
GET /finance/financial-planner/report/:id
GET /finance/financial-planner/report/download/:id
GET /health
```

## Optional User Header

```http
X-User-ID: 1
```

If passed, report fetch will be scoped to that user ID.

## Important

Use `.env.example` as reference. Do not commit real production passwords in `.env`.

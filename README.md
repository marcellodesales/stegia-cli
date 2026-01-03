# stegia

A Go CLI for TOTVS Datasul REST integration prototyping.

## What this repo includes

- `stegia totvs companies list`
  - Prints the HTTP request and a mocked HTTP response for `GET /api/btb/v1/companies` when `TOTVS_HOSTNAME=example.com`.
- `stegia totvs suppliers add -f <file.toon> [--company-id <id>]`
  - Parses a TOON file, selects a company (explicit or auto-match by city/state), prints the HTTP request, and prints a mocked create response when `TOTVS_HOSTNAME=example.com`.

The codebase is organized using:
- **Controller / Service / Builder** per feature (companies, suppliers)
- **Factory** for client/service creation
- **util** for env loading, TOON parsing, HTTP request/response printing, JSON helpers
- **logger** with debug/info/error levels

## Requirements

- Go 1.22+

## Quick start

```bash
# optional: create local env file
cat > local.env <<'EOF'
TOTVS_HOSTNAME=example.com
TOTVS_USERNAME=admin
TOTVS_PASSWORD=admin
LOG_LEVEL=debug
EOF

go mod tidy
go build -o stegia .

./stegia totvs companies list
./stegia totvs suppliers add -f ./examples/coca-cola.toon
```

## Environment selection

- If `ENV=prd`, the CLI loads `prd.env`
- Otherwise it loads `local.env`

If env files or vars are missing, it defaults to `admin:admin` and `example.com`.

## Notes

- This repo intentionally avoids real HTTP calls by default (prototype mode).
- Replace endpoints and payload field names to match your installed Datasul Swagger (`apipublic*.json`).

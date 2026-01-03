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

# CLI Examples

## totvs companies list

* Lists the companies

```console
$ ./stegia totvs companies list
time=2026-01-03T14:52:55.386-08:00 level=INFO msg="loaded env" envFile=/Users/marcellodesales/dev/github.com/marcellodesales/stegia-cli/local.env hostname=example.com

=== HTTP REQUEST ===
GET https://example.com/api/btb/v1/companies
Authorization: Basic YWRtaW46YWRtaW4=
Accept: application/json


=== HTTP RESPONSE ===
HTTP/1.1 200
Content-Type: application/json

{
  "items": [
    {
      "companyId": "01",
      "companyCode": "TOTVS-BR",
      "companyName": "TOTVS BRASIL MATRIZ",
      "country": "BR",
      "state": "SP",
      "city": "São Paulo",
      "status": "ACTIVE"
    },
    {
      "companyId": "02",
      "companyCode": "TOTVS-GO",
      "companyName": "TOTVS GOIÁS",
      "country": "BR",
      "state": "GO",
      "city": "Goiânia",
      "status": "ACTIVE"
    },
    {
      "companyId": "03",
      "companyCode": "TOTVS-RJ",
      "companyName": "TOTVS RIO DE JANEIRO",
      "country": "BR",
      "state": "RJ",
      "city": "Rio de Janeiro",
      "status": "INACTIVE"
    }
  ],
  "count": 3
}

=== COMPANIES (ACTIVE) ===
- companyId=01 code=TOTVS-BR name=TOTVS BRASIL MATRIZ (São Paulo/SP)
- companyId=02 code=TOTVS-GO name=TOTVS GOIÁS (Goiânia/GO)
```

## Command Help

* Some commands help 

```console
~/dev/github.com/marcellodesales/stegia-cli ⌚ 14:52:55
$ ./stegia totvs
TOTVS integration commands

Usage:
  stegia totvs [command]

Available Commands:
  companies   Company operations
  suppliers   Supplier (fornecedor) operations

Flags:
  -h, --help   help for totvs

Use "stegia totvs [command] --help" for more information about a command.

~/dev/github.com/marcellodesales/stegia-cli ⌚ 14:53:22
$ ./stegia totvs suss
TOTVS integration commands

Usage:
  stegia totvs [command]

Available Commands:
  companies   Company operations
  suppliers   Supplier (fornecedor) operations

Flags:
  -h, --help   help for totvs

Use "stegia totvs [command] --help" for more information about a command.

~/dev/github.com/marcellodesales/stegia-cli ⌚ 14:53:28
$ ./stegia totvs suppliers
Supplier (fornecedor) operations

Usage:
  stegia totvs suppliers [command]

Available Commands:
  add         Create a supplier from a TOON file (lists companies first)

Flags:
  -h, --help   help for suppliers

Use "stegia totvs suppliers [command] --help" for more information about a command.

~/dev/github.com/marcellodesales/stegia-cli ⌚ 14:53:32
$ ./stegia totvs suppliers add
Error: required flag(s) "file" not set
Usage:
  stegia totvs suppliers add [flags]

Flags:
      --company-id string   CompanyId header value (optional; auto-selected if omitted)
  -f, --file string         Path to .toon file (TOON format)
  -h, --help                help for add

required flag(s) "file" not set

~/dev/github.com/marcellodesales/stegia-cli ⌚ 14:53:34
$ ./stegia totvs suppliers add  --company-id 2 -f
```

## Supplier Example 

```yaml
supplierType: JURIDICAL
supplierName: COCA-COLA INDUSTRIAS LTDA
tradeName: COCA-COLA BRASIL
taxId:
  cnpj: "45997000000104"
status: ACTIVE
country: BR
address:
  street: Av. Anhanguera
  number: "5000"
  district: Setor Central
  city: Goiânia
  state: GO
  zipCode: "74043010"
contact:
  email: financeiro@cocacola.com.br
  phone: +556230000000
classification:
  supplierGroup: NACIONAL
integration:
  externalId: "toon:coca-cola-br-go"
  sourceSystem: stegia%
```

## Companies Add

* Add a company based on an existing ACTIVE company

```console
$ ./stegia totvs suppliers add  --company-id 2 -f examples/coca-cola.toon
time=2026-01-03T14:55:55.397-08:00 level=INFO msg="loaded env" envFile=/Users/marcellodesales/dev/github.com/marcellodesales/stegia-cli/local.env hostname=example.com
time=2026-01-03T14:55:55.398-08:00 level=INFO msg="parsed TOON file" file=examples/coca-cola.toon
time=2026-01-03T14:55:55.398-08:00 level=INFO msg="selected company" companyId=2 reason="explicit --company-id (not found/ACTIVE in list)"

=== HTTP REQUEST ===
POST https://example.com/api/cdp/v1/suppliers
Authorization: Basic YWRtaW46YWRtaW4=
Companyid: 2
Content-Type: application/json
Accept: application/json

{
  "address": {
    "city": "Goiânia",
    "complement": "",
    "district": "Setor Central",
    "number": "5000",
    "state": "GO",
    "street": "Av. Anhanguera",
    "zipCode": "74043010"
  },
  "country": "BR",
  "integration": {
    "externalId": "toon:coca-cola-br-go",
    "sourceSystem": "stegia"
  },
  "source": {
    "address": {
      "city": "Goiânia",
      "district": "Setor Central",
      "number": "5000",
      "state": "GO",
      "street": "Av. Anhanguera",
      "zipCode": "74043010"
    },
    "classification": {
      "supplierGroup": "NACIONAL"
    },
    "contact": {
      "email": "financeiro@cocacola.com.br",
      "phone": "+556230000000"
    },
    "country": "BR",
    "integration": {
      "externalId": "toon:coca-cola-br-go",
      "sourceSystem": "stegia"
    },
    "status": "ACTIVE",
    "supplierName": "COCA-COLA INDUSTRIAS LTDA",
    "supplierType": "JURIDICAL",
    "taxId": {
      "cnpj": "45997000000104"
    },
    "tradeName": "COCA-COLA BRASIL"
  },
  "status": "ACTIVE",
  "supplierName": "COCA-COLA INDUSTRIAS LTDA",
  "supplierType": "JURIDICAL",
  "taxId": {
    "cnpj": "00000000000000"
  },
  "tradeName": "COCA-COLA BRASIL"
}

=== HTTP RESPONSE ===
HTTP/1.1 201
Content-Type: application/json

{
  "companyId": "2",
  "createdAt": "2026-01-03T22:55:55Z",
  "echoRequest": {
    "address": {
      "city": "Goiânia",
      "complement": "",
      "district": "Setor Central",
      "number": "5000",
      "state": "GO",
      "street": "Av. Anhanguera",
      "zipCode": "74043010"
    },
    "country": "BR",
    "integration": {
      "externalId": "toon:coca-cola-br-go",
      "sourceSystem": "stegia"
    },
    "source": {
      "address": {
        "city": "Goiânia",
        "district": "Setor Central",
        "number": "5000",
        "state": "GO",
        "street": "Av. Anhanguera",
        "zipCode": "74043010"
      },
      "classification": {
        "supplierGroup": "NACIONAL"
      },
      "contact": {
        "email": "financeiro@cocacola.com.br",
        "phone": "+556230000000"
      },
      "country": "BR",
      "integration": {
        "externalId": "toon:coca-cola-br-go",
        "sourceSystem": "stegia"
      },
      "status": "ACTIVE",
      "supplierName": "COCA-COLA INDUSTRIAS LTDA",
      "supplierType": "JURIDICAL",
      "taxId": {
        "cnpj": "45997000000104"
      },
      "tradeName": "COCA-COLA BRASIL"
    },
    "status": "ACTIVE",
    "supplierName": "COCA-COLA INDUSTRIAS LTDA",
    "supplierType": "JURIDICAL",
    "taxId": {
      "cnpj": "00000000000000"
    },
    "tradeName": "COCA-COLA BRASIL"
  },
  "links": {
    "self": "/api/cdp/v1/suppliers/SUP-902341"
  },
  "message": "Mocked response (example.com); no real Datasul call executed.",
  "status": "CREATED",
  "supplierCode": "FORN-000902341",
  "supplierId": "SUP-902341"
}
```



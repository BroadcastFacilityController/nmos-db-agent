# BFC NMOS DB Agent

This package provides a go service which will maintain a connection with an external NMOS registry (such as [Easy-NMOS](https://github.com/rhastie/easy-nmos)) and will continuously report changes from the Registry into a live-updating NMOS database to be used by Broadcast Facility Controller (BFC).

## Environment Variables
| Variable              | Description                      | Options (*Default)               | Example                           |
|-----------------------|----------------------------------|----------------------------------|-----------------------------------|
| DEBUG_LEVEL           | Logging level                    | debug, info*, warn, error, fatal | DEBUG_LEVEL=info                  |
| DB_URL                | Postgres Database URL            | string                           | DB_URL=localhost                  |
| DB_PORT               | Postgres Database Port           | int                              | DB_PORT=5432                      |
| DB_USER               | Postgres Database Username       | string                           | DB_USER=admin                     |
| DB_PASSWORD           | Postgres Database Password       | string                           | DB_PASSWORD=admin                 |
| NMOS_REGISTRY_ADDRESS | NMOS Registry (QueryAPI) Address | string                           | NMOS_REGISTRY_ADDRESS=10.10.60.10 |
| NMOS_REGISTRY_PORT    | NMOS Registry (QueryAPI) Port    | int                              | NMOS_REGISTRY_PORT=8080           |

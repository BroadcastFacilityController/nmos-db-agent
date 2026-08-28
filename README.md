# BFC NMOS DB Agent 2

This package provides a go service which will maintain a connection with an external NMOS registry (such as [Easy-NMOS](https://github.com/rhastie/easy-nmos)) and will continuously report changes from the Registry into a live-updating NMOS database to be used by Broadcast Facility Controller (BFC).

## Environment Variables
| Variable               | Description                                     | Options (*Default)               | Example                                                                |
| ---------------------- | ----------------------------------------------- | -------------------------------- | ---------------------------------------------------------------------- |
| LOG_LEVEL              | Logging level                                   | debug, info*, warn, error, fatal | DEBUG_LEVEL=info                                                       |
| NMOS_DB_URL            | Postgres Database URL                           | string                           | NMOS_DB_URL=postgres://admin:admin@localhost:5432/nmos?sslmode=disable |
| NMOS_DB_RUN_MIGRATIONS | Run migrations on the database (take ownership) | true, false*                     | NMOS_DB_RUN_MIGRATIONS=true                                            |
| NMOS_QUERY_URL         | NMOS Registry (QueryAPI) Address                | string                           | NMOS_QUERY_URL=http://172.25.50.101/x-nmos                             |


## Docker Compose
```yml
services:
  bfc-nmos-db-agent:
    image: ghcr.io/broadcastfacilitycontroller/nmos-db-agent:latest
    container_name: nmos-db-agent
    hostname: bfc-nmos-db-agent
    restart: unless-stopped
    environment:
      - DEBUG_LEVEL=info
      - NMOS_DB_URL=postgres://admin:admin@localhost:5432/nmos?sslmode=disable
      - NMOS_DB_RUN_MIGRATIONS=true
      - NMOS_QUERY_URL=http://172.25.50.101/x-nmos
```
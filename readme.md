# crud-docker

Simple CRUD application to test if the multi container application is working on another machine.

## Project structure

```
curd-docker
    |_cmd
    |   |_main.go
    |_config
    |   |_local.yaml
    |_internal
    |   |_config
    |   |   |_config.go
    |   |_handlers
    |   |   |_userHandler.go
    |   |_middlewares
    |   |   |_middleware.go
    |   |_storage
    |   |   |_postgresql
    |   |   |   |_postgresql.go
    |   |   |_storage.go
    |   |_types
    |   |   |_types.go
    |   |_utils
    |       |_error
    |       |   |_error.go
    |       |_json
    |           |_json.go
    |_.dockeringore
    |_.env
    |_.gitignore
    |_compose.yml
    |_Dockerfile
    |_go.mod
    |_go.sum
    |_readme.md
```

## Steps for configuration

1. Create a `.env` file and declare the following variables:

   ```
   POSTGRES_USER=username           (postgres username)
   POSDTGRES_PASSWORD=password      (postgres password)
   POSTGRES_DB=dbname               (database name)
   DB_HOST=db                       (database host)
   DB_PORT=5432                     (database port)
   CONFIG_PATH=config/local.yaml    (path to config file)
   ```

2. Create a directory named `config` and declare `local.yaml` file within it

   ```
   curd-docker
       |_config
            |_local.yaml
   ```

3. Declare the following configurations within `local.yaml`
   ```
    env: "dev"
    http_server:
      address: ":8080"
    database:
      postgresql:
         user: "postgres"
         password: "secret"
         dbname: "database"
         host: "db"
         port: "5432"
   ```

## Starting up the containers

Build and run the containers by executing the following command:
<br>
`docker compose up --build -d`

## Shutting down the containers

Shut down the containers by executing the following command:
<br>
`docker compose down`

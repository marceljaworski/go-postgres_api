# Golang postgreSQL products API

## Features

- PostgresSQL DB
- Docker compose
- CRUD API

## Build App image

`docker build -t marceljaworski/go-postgres_api:1.0 .`

## Build Postgres container

`p` project name
`docker compose -p postgres up -d`
or..
up
`docker-compose -p postgres -f docker-compose.yaml up -d`
down
`docker-compose -p postgres -f docker-compose.yaml down`
stop
`docker-compose -p postgres -f docker-compose.yaml stop`
start
`docker-compose -p postgres -f docker-compose.yaml start`

## Postgres usefull commands

exec container
`docker exec -it containerName bash`

Run the psql command to connect to  the PostgreSQL database

`psql -U username -d database_name` or..

Change to postgres user
`su - postgres`

Sign in into psql
`psql`

list databases
`\l`

Sign into the database
`psql productsdb`

Ensure you are conected to a database
`\c productsdb`

List of the available tables
`\d`

Show table products'
`\d products`

Use CASCADE with DROP TABLE (and DROP SCHEMA)

`DROP TABLE table_name CASCADE;`

In this project, the user, db and tables have already been created. But there are these commands to do it:

`CREATE USER postgres_user WITH PASSWORD 'password';`
`CREATE DATABASE my_postgres_db OWNER postgres_user;`

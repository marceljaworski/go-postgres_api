# Golang postgreSQL products API

## Features

- PostgresSQL DB
- Docker compose
- CRUD API

## Build Postgres container

`p` project name
`docker compose -p postgres up -d`

## Postgres usefull commands

exec container
`docker exec -it containerName bash`

Change to postgres user
`su - postgres`

Sign in into psql
`psql`

list databases
`\l`

Sign into the database
`psql productsdb`

Ensure you are conected ro database
`\c products`

List of the available tables
`\d`

Show table products
`\d products`

Use CASCADE with DROP TABLE (and DROP SCHEMA)

`DROP TABLE table_name CASCADE;`

In this project, the user, db and tables have already been created. But there are these commands to do it:

`CREATE USER postgres_user WITH PASSWORD 'password';`
`CREATE DATABASE my_postgres_db OWNER postgres_user;`

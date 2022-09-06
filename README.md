
<!--toc-->
- [Challenge](#challenge)
- [Setup](#setup)
    * [Go](#go)
    * [Docker-compose](#docker-compose)
    * [Configuration](#configuration)
    * [Migrations](#migrations)
    * [Run](#run)
- [Documentation](#documentation)
    * [Framework](#framework)
    * [ORM](#orm)
    * [Model Generation](#model-generation)
    * [Tests](#tests)
    * [API](#api)
    * [Possible additions](#possible-additions)

<!-- tocstop -->

# Challenge

See [here](CHALLENGE.md).

# Setup

## Go

Version 1.19+ needed.

## Docker-compose

Docker & docker-compose(2.7.0+)

`docker-compose up -d`

## Configuration

For local development just copy `config/.example.challenge.yaml` to `config/.challenge.yaml`

**NB! Change default user/password for `docsAuth` inside config before deploying anywhere.** 

## Migrations

Install [goose](https://github.com/pressly/goose):

`go install github.com/pressly/goose/v3/cmd/goose@latest`

Then use goose directly or via makefile commands.

To apply all missing migrations:

`make migrations_up`

To go down by one:

`make migrations_down`

## Linter

See [golangci-lint installation](https://golangci-lint.run/usage/install/#local-installation)

## Run

To import data_dump.csv run `make run_import`

To start api run `make run_api`, [api docs](#api)

# Documentation

## Framework

[Echo](https://echo.labstack.com/) is high performance and minimalist framework, which is useful for this 
particular task. Whole initialization and routing is short and transparent. It is well maintained and has a lot of stars
 on github.

## ORM

[SQLBoiler](https://github.com/volatiletech/sqlboiler) is a "database-first" ORM. Has a model generator which creates 
type safe models. Also uses very convenient null package for working with nullable values. Although it wasn't really 
necessary to use sqlboiler here, it's so easy to install and use that I decided to still do it.

## Model Generation

Install first
```
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
```
Then for dev db
```
make generate_models
```
Models will be placed under `internal/repository/model`

## Tests

Run api `make run_api` then `make test` (includes integration tests)

To show cover profile `make coverage`.

For generated model tests in `internal/repository/model` we need postgresql client version >= 14. Otherwise, those tests will fail.

## API

API described as OpenAPI scheme.
Available [here](http://localhost:3011/docs/) after `make run_api`. Protected by basic auth, credentials are 
described in configuration.  

OpenAPI scheme is [here](api/docs/openapi.yaml)

## Possible additions

 - dockerfile
 - [load testing](https://github.com/tsenart/vegeta)
 - code-based openapi scheme generation
 - profiling
 - different environments support
 - [leak test](https://github.com/uber-go/goleak)
 - gRPC api interface
 - cross-platform
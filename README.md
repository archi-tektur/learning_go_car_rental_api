# Food API

[API Doc](api/swagger.json)

## How to build

```bash
make build
```

## How to run tests

```bash
make test
```

## How to run the server

From source code:

```bash
go run cmd/*.go --CORS_ALLOW_ORIGIN=localhost server
```

Or after build:

```bash
bin/food-app --CORS_ALLOW_ORIGIN=localhost server
```

Using env vars:

```bash
export CORS_ALLOW_ORIGIN=localhost
bin/food-app server
```

## How to build a docker container

```bash
make container
```

## How to update API Doc

```bash
make swagger
```

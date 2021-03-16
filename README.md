# Lookum

Lookum is an online shopping market for various products.
Docker container image supported.
Docker compose file is added to the repo for easy install

## Prerequisites

| Program | Usage |
|----|----|
| docker | Container image building |
| docker-compose | Local deployment |
| make | Repo building process |
| go | go project building(version: 1.16) |

## Build

```bash
make all
```

## Container Image Build

```bash
make docker
```

## Run

### Run individually

```bash
cd docker && ./run.sh -d api
```

### Run using the compose file

```bash
cd docker && ./compose.sh up -d
```

or

```bash
docker-compose -f docker/docker-compose.yml up -d 
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

# Micromon

## About

This is an experiment to study Go and Microservices.  This isn't meant to be practical, it's more about learning than anything.  And it's also for the fun of it.

## Components

### Breed Service

- Stores the "Breed" information for each species of Pokemon.
- Uses BoltDB for Database

### Move Service

- Stores the "Move" or Attack information
- Uses Vanilla Golang (`map[string]pokemon.Move`)

### Pokemon Service

- Stores instantiated Pokemon (the ones that fight and grow)
- Uses Postgres (`go-pg`, `protoc-go-inject-tag`, and `protobuf`)
- Experiment based on [this Medium Article](https://medium.com/@vptech/complexity-is-the-bane-of-every-software-engineer-e2878d0ad45a)

### Gateway API

- Proxy for gRPC Services
- Uses gRPC-Gateway  

## Setup

This project runs in Docker Compose, and so a simple 
```
git pull
```

and 

```
docker-compose up
``` 
should do the trick


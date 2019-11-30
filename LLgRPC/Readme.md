# LLgRPC

Go implementation of a procedural middleware similar to RPC.

## Table of Contents
- [LLgRPC](#llgrpc)
  - [Table of Contents](#table-of-contents)
  - [Requirements](#requirements)
  - [Installing](#installing)
  - [Testing](#testing)

## Requirements

- go 1.13
- gomod
- git

## Installing

1. Install golang version 1.13.
2. Clone the project:
   
```sh
https://github.com/lucas625/Middleware.git
```

## Testing

- Run the naming server:

```sh
go run LLgRPC/client-server/naming-server/naming-server.go
```

- Run the server:

```sh
go run LLgRPC/client-server/server/server.go
```

- Run the client:

```sh
go run LLgRPC/client-server/client/client.go
```
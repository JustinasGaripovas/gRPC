# Workshop for gRPC
## Intro
This is not production ready code, this is test for gRPC possibilities. Project consists of 3 sub-projects.

**nodeServer** - gRPC `node` server
**protos** - Protoc compiler with go language
**pythonClient** - gRPC `python` stub/client

**nodeServer** and **pythonClient** are services that communicate with each other using gRPC protocol.

**protos** is used to check out **protoc** compiler for `go` language that generates code from `.protos` files.

## Usage
In `root` directory run
```
docker-compose up -d
```

### Client
In order to access clients container
```
docker-compose exec python sh
```

afterwards if order to run the `python` client
```
python client.py
```

Which allows you to select one of 4 modes of gRPC:
```
PingPong = 1 (Unary)
PingsPong = 2 (Client->Server stream)
PingPongs = 3 (Server->Client stream)
PingsPongs = 4 (Bi-directional stream)
```

Note: `Bi-directional` stream is not implemented yet

## More info
Read more about gRPC [here](./gRPC.pdf)
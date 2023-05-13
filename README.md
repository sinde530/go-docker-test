# go-docker-test

```go
├go-docker-test
├── cmd
└── main.go
├── pkg
│ └── db
│   └── db.go
├── models
└── go.mod
└── go.sum
└── DockerFile
└── README.md
```

## docker cli

```docker
docker build . -t go-dock

docker image ls

```

# raspberry pi 4

**all containers remove**

```docker
docker rm $(docker ps -aq)
```

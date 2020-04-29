# go-docker-realize

```bash
export GOFLAGS=-mod=vendor
export GO111MODULE=on
go mod init github.com/yangpeng-chn/go-docker-realize  (go.mod will be generated)
go mod download
go mod vendor
go mod verify

docker image build --tag go-realize:latest .
docker container run -v $(pwd):/work -w /work -p 8080:8080 go-realize:latest
```

Access http://localhost:8080

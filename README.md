# go-docker-realize

Use `realize` package to test Go application with hot-reload when changing source files.

### Use realize without docker

```bash
$ GO111MODULE=off go get github.com/oxequa/realize
$ vi ~/.zprofile
$ source ~/.zprofile
---
export PATH=$PATH:$GOPATH/bin
---
$ which realize

$ realize start --run # also works without --run
[14:36:13][API] : Watching 1 file/s 1 folder/s
[14:36:13][API] : Build started
[14:36:14][API] : Build completed in 0.892 s
[14:36:14][API] : Running..
[14:36:14][API] : Starting web server on port 8080
```

### Use realize with docker

Note: removing the following lines also works

```
ENV APP_HOME /app

RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME

EXPOSE 8080
```

```bash
$ export GOFLAGS=-mod=vendor
$ export GO111MODULE=on (may need to set to off if error occured)
$ go mod init github.com/yangpeng-chn/go-docker-realize  (go.mod will be generated)
$ go mod download
$ go mod vendor
$ go mod verify

$ docker image build --tag go-realize:latest .
$ docker container run -v $(pwd):/app -w /app -p 8080:8080 go-realize:latest
[06:33:34][API] : Watching 1 file/s 1 folder/s
[06:33:34][API] : Build started
[06:33:38][API] : Build completed in 3.512 s
[06:33:38][API] : Running..
[06:33:38][API] : Starting web server on port 8080
```

### Use realize with docker and docker-compose

```bash
$ docker-compose up --build
Building app
Step 1/7 : FROM golang:1.14
 ---> 2421885b04da
Step 2/7 : RUN go get github.com/oxequa/realize
 ---> Using cache
 ---> 2131ca7f8662
Step 3/7 : ENV APP_HOME /app
 ---> Using cache
 ---> a6a9a670c9cb
Step 4/7 : RUN mkdir -p $APP_HOME
 ---> Using cache
 ---> c2467a23fca1
Step 5/7 : WORKDIR $APP_HOME
 ---> Using cache
 ---> 0356ac1555ac
Step 6/7 : EXPOSE 8080
 ---> Using cache
 ---> eef7ee715adb
Step 7/7 : CMD [ "realize", "start", "--run" ]
 ---> Using cache
 ---> abbabedc1a0a

Successfully built abbabedc1a0a
Successfully tagged go-docker-realize_app:latest
Recreating api ... done
Attaching to api
api    | [06:57:03][API] : Watching 1 file/s 1 folder/s
api    | [06:57:03][API] : Build started
api    | [06:57:03][API] : Build completed in 0.356 s
api    | [06:57:03][API] : Running..
api    | [06:57:03][API] : Starting web server on port 8080

 (CTL+C, stop)
$ docker-compose down --remove-orphans --volumes
Removing api ... done
Removing network go-docker-realize_default
```

Access [http://localhost:8080](http://localhost:8080/)
FROM golang:1.14

RUN go get github.com/oxequa/realize

ENV APP_HOME /app

RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME

EXPOSE 8080

CMD [ "realize", "start", "--run" ]
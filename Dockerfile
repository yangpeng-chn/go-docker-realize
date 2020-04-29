FROM golang:latest

RUN go get github.com/oxequa/realize
ENTRYPOINT ["realize"]
CMD ["start"]

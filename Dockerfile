FROM golang:latest

RUN go get github.com/oxequa/realize
ENTRYPOINT ["realize"]
CMD ["start"]

# docker image build --tag go-realize:latest .
# docker container run -v $(pwd):/work -w /work -p 8080:8080 go-realize:latest

FROM golang as builder

ENV MONGOURI="temp"

WORKDIR /work

COPY . .

RUN set -x && \
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o jokesapi

FROM debian:stretch-slim

RUN apt-get update -y && apt-get install ca-certificates -y

EXPOSE 8000

WORKDIR /work
COPY --from=builder /work .
CMD ["./jokesapi"]
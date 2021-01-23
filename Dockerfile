FROM golang as builder

WORKDIR /work

COPY . .

RUN set -x && \
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o jokesapi

FROM public.ecr.aws/lambda/go

# RUN apt-get update -y && apt-get install ca-certificates -y 

COPY --from=builder /work/jokesapi ${LAMBDA_TASK_ROOT}
CMD ["jokesapi"]
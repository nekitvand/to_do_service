FROM golang:alpine AS builder
RUN apk add --update make git protoc protobuf protobuf-dev curl gcc musl-dev

ARG GITHUB_PATH=github.com/nekitvand/to_do_service

COPY Makefile /home/${GITHUB_PATH}/Makefile
COPY go.mod /home/${GITHUB_PATH}/go.mod 
COPY go.sum /home/${GITHUB_PATH}/go.sum
COPY pkg /home/${GITHUB_PATH}/pkg

WORKDIR /home/${GITHUB_PATH}
# RUN go mod tidy
RUN make deps

COPY . /home/${GITHUB_PATH}
RUN make build

# gRPC

FROM alpine:latest as server
RUN apk --no-cache add ca-certificates
WORKDIR /root/

ARG GITHUB_PATH=github.com/nekitvand/to_do_service

COPY --from=builder /home/${GITHUB_PATH}/bin/grpc-server .
COPY --from=builder /home/${GITHUB_PATH}/config.yaml .
COPY --from=builder /home/${GITHUB_PATH}/swagger ./swagger
COPY --from=builder /home/${GITHUB_PATH}/migrations/ ./migrations

RUN chown root:root grpc-server

EXPOSE 50051
EXPOSE 8080
EXPOSE 9100

CMD ["./grpc-server"]
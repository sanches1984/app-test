FROM golang:alpine as builder
RUN apk update && \
    mkdir /app
ADD . /app/
WORKDIR /app
COPY ./ ./

RUN go install -v ./cmd/harmony && \
    go build -o ./bin/harmony ./cmd/harmony/main.go

FROM alpine
RUN apk update && \
    mkdir /app

COPY --from=builder /app/bin/harmony /app

WORKDIR /app

EXPOSE 8080

CMD /app/harmony
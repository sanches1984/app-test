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
    adduser -D -H -h /app harmony && \
    mkdir -p /app/config  && \
    chown -R harmony:harmony /app

USER harmony

COPY --chown=harmony --from=builder /app/bin/harmony /app

WORKDIR /app

CMD /harmony

FROM golang:alpine as builder
RUN apk update && mkdir /app
ADD . /app/
WORKDIR /app
COPY ./ ./

RUN go install -v ./cmd/referral && \
    go build -o ./bin/referral ./cmd/referral/main.go

FROM alpine
RUN apk update && mkdir /app

COPY --from=builder /app/bin/referral /app

WORKDIR /app

CMD /app/referral
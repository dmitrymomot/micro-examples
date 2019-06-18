FROM golang:alpine AS container
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh curl ca-certificates
RUN mkdir -p /app_source
WORKDIR /app_source
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download


FROM container AS builder
RUN rm -Rvf /go/pkg/mod/github.com/hashicorp/consul@v1.4.4/api 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -i -o /health ./main.go


FROM alpine:latest
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*
COPY --from=builder /health /
WORKDIR /
ENTRYPOINT [ "/health" ]

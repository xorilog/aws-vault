FROM golang:1.22 as aws-vault-proxy
WORKDIR /usr/src/aws-vault-proxy
COPY ./http-proxy/* /usr/src/aws-vault-proxy/
RUN go build -v -o /usr/local/bin/aws-vault-proxy ./...
CMD ["/usr/local/bin/aws-vault-proxy"]

FROM golang:1.22-alpine as socat-proxy
RUN apk add --no-cache socat
WORKDIR /usr/src/aws-vault-socat-proxy/
COPY ./socat/* /usr/src/aws-vault-socat-proxy/
RUN go build -v -o /usr/local/bin/aws-vault-socat-proxy ./...
ENTRYPOINT ["/usr/local/bin/aws-vault-socat-proxy"]

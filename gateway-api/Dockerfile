FROM golang:latest

WORKDIR /go/src/github.com/Mrcampbell/pgo2/gateway-api

RUN go get -u github.com/oxequa/realize
RUN go get -u github.com/kardianos/govendor

COPY . .
ADD .realize.yaml .realize.yaml
ADD vendor/vendor.json vendor/vendor.json
RUN govendor add +e

EXPOSE 8081

CMD realize start
FROM golang:latest

WORKDIR /go/src/github.com/Mrcampbell/pgo2/pokemon-service

RUN go get -u github.com/oxequa/realize
RUN go get -u github.com/kardianos/govendor

COPY . .
ADD .realize.yaml .realize.yaml
ADD vendor/vendor.json vendor/vendor.json
RUN govendor add +e

EXPOSE 9090

CMD realize start
FROM golang:alpine AS build

WORKDIR /go/src/github.com/urcomputeringpal/nuc-node-led-controller
COPY ./vendor ./vendor
RUN go install -v ./vendor/...
COPY . .
RUN go install -v github.com/urcomputeringpal/nuc-node-led-controller


FROM alpine
COPY --from=build /go/bin/nuc-node-led-controller /go/bin/nuc-node-led-controller
ENTRYPOINT ["/go/bin/nuc-node-led-controller"]

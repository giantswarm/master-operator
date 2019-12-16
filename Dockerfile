FROM alpine:3.10

RUN apk add --no-cache ca-certificates

ADD ./master-operator /master-operator

ENTRYPOINT ["/master-operator"]

# BUILD
FROM golang:latest as build

WORKDIR /service
ADD . /service

RUN cd /service && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /http-service .

CMD /http-service

# TEST
FROM build as test

RUN chmod -R 700 /go
RUN curl -LO https://github.com/pact-foundation/pact-ruby-standalone/releases/download/v1.81.0/pact-1.81.0-linux-x86_64.tar.gz
RUN tar xzf pact-1.81.0-linux-x86_64.tar.gz
ENV PATH /service/pact/bin/:$PATH


# PRODUCTION
FROM alpine:latest as production

RUN apk --no-cache add ca-certificates
COPY --from=build /http-service ./
RUN chmod +x ./http-service

ENTRYPOINT ["./http-service"]

EXPOSE 8080
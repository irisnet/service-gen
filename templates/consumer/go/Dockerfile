FROM golang-alpine as builder

WORKDIR /app
ADD . .

RUN make install

FROM golang-alpine

COPY --from=builder /go/bin/{{service_name}}-sp /usr/local/bin

ENTRYPOINT ["{{service_name}}-sp", "start"]

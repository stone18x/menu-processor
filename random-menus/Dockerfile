FROM golang:1.16.2-alpine3.13 AS builder

WORKDIR /dockersrc/
COPY main.go /dockersrc/
COPY go.* /dockersrc/
COPY ./function /dockersrc/function

RUN ls

RUN CGO_ENABLED=0 go build -o /bin/image

FROM scratch
COPY --from=builder /bin/image /bin/app
ENTRYPOINT ["/bin/app"]
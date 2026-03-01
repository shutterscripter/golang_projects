FROM golang:alpine as builder

RUN mkdir /build

ADD . /build

WORKDIR /build

RUN CGO_ENABLED=0 go build -o main .

# stage 2

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder --chmod=0755 /build/main /app/

WORKDIR /app

EXPOSE 3000 

CMD ["./main"]
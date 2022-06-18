FROM golang:1.17-alpine3.15 as building-stage

WORKDIR /building-stage

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN  go build -o dockerized-macaiki ./cmd/macaiki

FROM alpine:3.15
WORKDIR /app
COPY --from=building-stage /building-stage/dockerized-macaiki .
COPY --from=building-stage /building-stage/.env .

EXPOSE 8080

CMD [ "/app/dockerized-macaiki" ]
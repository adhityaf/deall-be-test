FROM golang:latest as builder

ENV APP deall-be-test

WORKDIR /app

COPY . ./

COPY go.mod go.sum ./
RUN go mod download 

RUN touch .env
COPY .env-example .env

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/${APP} .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /bin/${APP} /bin/
COPY --from=builder /app/.env /bin/

COPY . ./
COPY .env-example /app/.env

ENTRYPOINT ["/bin/deall-be-test"]
FROM golang:1.23.2-bullseye AS build

WORKDIR /usr/src/app

COPY go.mod ./

RUN go mod download

COPY . .
# CGO disables the c libraries we dont need for this app. GOOS sets OS to linux.
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:3.18 AS app

RUN adduser -D appuser

WORKDIR /app

COPY --from=build /usr/src/app/app .

USER appuser

EXPOSE 8080

CMD ["./app"]




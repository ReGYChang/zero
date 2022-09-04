FROM r19750927/nexdata-go:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN make build

FROM alpine:3.6

WORKDIR /

COPY --from=builder /app/release/auth ./auth

RUN /auth -v

EXPOSE 80 8080

ENTRYPOINT ["/auth"]

FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /myprogram

FROM gcr.io/distroless/base-debian11 AS release

WORKDIR /
COPY --from=builder /myprogram /myprogram

USER nonroot:nonroot

ENTRYPOINT ["/myprogram"]

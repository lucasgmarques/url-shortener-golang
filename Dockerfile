FROM golang:1.18.0-bullseye AS builder

WORKDIR /app

COPY go.* ./

# build the Go application
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo -o app .

# Stage 2: Create a minimal runtime image
FROM alpine:3.19.1

RUN adduser -D -g '' user

RUN apk --no-cache add ca-certificates && \
    apk update

WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/app ./

RUN chown -R user:user /app

USER user

EXPOSE 8080

CMD [ "./app" ]


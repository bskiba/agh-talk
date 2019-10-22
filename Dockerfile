FROM golang:1.11.3
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o agh-demo .

FROM alpine:latest
COPY --from=0 /app/agh-demo .
ENTRYPOINT ["/agh-demo"]

# Build the Go API
FROM golang:1.18-alpine AS builder
ADD . /app
WORKDIR /app/backend
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .
# Build the React application
FROM node:16-alpine3.15 AS node_builder
COPY --from=builder /app/frontend ./
RUN npm install
RUN npm run build
# Final stage build, this will be the container
# that we will deploy to production
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
COPY --from=node_builder /build ./web
RUN chmod +x ./main
EXPOSE 8080
CMD ./main
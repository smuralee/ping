FROM golang:alpine as builder
RUN apk update && apk add --no-cache git

RUN addgroup -S alpha && adduser -S alpha -G alpha
RUN mkdir -p /home/alpha/app && chown -R alpha:alpha /home/alpha/app

WORKDIR /home/alpha/app
USER alpha

# For building Go Module required
ENV GOPROXY=direct
ENV GO111MODULE=on
ENV GOARCH=amd64
ENV GOOS=linux
ENV CGO_ENABLED=0

# Copy the Go Modules manifests
COPY --chown=alpha:alpha go.mod go.sum ./

# Cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY --chown=alpha:alpha . .

# Build
RUN go build -a -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN addgroup -S alpha && adduser -S alpha -G alpha
RUN mkdir -p /home/alpha/app && chown -R alpha:alpha /home/alpha/app

WORKDIR /home/alpha/app
USER alpha

COPY --chown=alpha:alpha --from=builder /home/alpha/app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]

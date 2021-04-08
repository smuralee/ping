FROM golang:alpine
RUN apk update && apk add --no-cache git

RUN addgroup -S alpha && adduser -S alpha -G alpha
RUN mkdir -p /home/alpha/app && chown -R alpha:alpha /home/alpha/app

WORKDIR /home/alpha/app
USER alpha

# Copy the go source
COPY --chown=alpha:alpha . .

# Build
RUN go build main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]


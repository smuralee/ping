# Ping Container

Container image for testing purpose. Returns the client IP address and status message

## Build and run the container

```shell
docker built -t ping .
docker run -p 8080:8080 ping
```

## Response structure
```json
{"IP":"172.31.0.1:55667","Status":"Success"}
```
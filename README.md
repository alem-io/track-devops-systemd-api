# track-devops-systemd-api

Simple API for systemd-api exercise

## Build

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o api .
```

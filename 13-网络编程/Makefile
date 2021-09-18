build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /tmp/server 1-tcp-server.go

scp:
	scp -r /tmp/server stx:/home/qiaocc/code/tcp

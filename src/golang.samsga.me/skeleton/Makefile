.PHONY:	client server
all:	client server
client:
	cd client && go install client.go `ls -1 *.go|grep -v client.go`
server:
	cd server && go install server.go `ls -1 *.go|grep -v server.go`

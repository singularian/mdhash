# This is how we want to name the binary output
BINARY=mdhash

# Project name
PROJECT="mdhash"

# These are the values we want to pass for Version and BuildTime
VERSION=1.0.6
BUILD_TIME=`date +%FT%T%z`

# command path
CMD=cmd


all:
	go build -o testCubehash testCubehash.go
	go build -o testdjb2 testDJB2.go 
	go build -o testFarmHash testFarmHash.go
	go build -o testJenkinsHash testJenkinsHash.go
	go build -o testHW32 testHW32.go
	go build -o testpoly1305 testpoly1305.go
	go build -o testsdbm testSDBM.go 
	go build -o testSHA1-128 testSHA1-128.go
	go build -o testspooky testSpookyHash.go
	go build -o testXXhash128 testXXhash128.go
clean:
	rm -r testCubehash testdjb2 testHW32 testFarmHash testJenkinsHash testpoly1305 testsdbm testSHA1-128 testspooky testXXhash128

# Display general help about this command
help:
	@echo "$(PROJECT) Makefile."
	@echo "GOPATH=$(GOPATH)"
	@echo "The following commands are available:"
	@echo "    make all         : Build the mdhash files"
	@echo "    make clean       : Remove any build artifact"
list:
	echo "all clean list"

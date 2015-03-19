.PHONY: deps run slow-lo0 slugish-lo0 normal-lo0 proto

deps:
	go get -u github.com/libgit2/git2go && \
		cd $GOPATH/src/github.com/libgit2/git2go && \
		git checkout next && git submodule update --init && \
		make install
	go get -u github.com/ddollar/forego
	go get -u ./...

run:
	mkdir -p ~/git-explorer-data
	GIT_EXPLORER_STORAGE_PATH=~/git-explorer-data forego start -p 9090

slow-lo0:
	comcast --device=lo0 --latency=250 --target-bw=1000

slugish-lo0:
	comcast --device=lo0 --latency=700 --target-bw=1000

normal-lo0:
	comcast --mode stop

proto:
	protoc -I pb --go_out=plugins=grpc:pb pb/*.proto

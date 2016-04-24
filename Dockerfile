FROM golang
RUN go get github.com/tools/godep
RUN mkdir -p /go/src/github.com/nkobber/example-go-webapp
WORKDIR /go/src/github.com/nkobber/example-go-webapp
ADD . ./
RUN godep go build

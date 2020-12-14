FROM golang:latest

ENV GO111MODULE=on

WORKDIR /go/src/app

#RUN mkdir module
#RUN cd module
#RUN go mod init github.com/maihiro/apimodule
#RUN cd ..


#COPY module .
COPY server server
#COPY sampleApi.go .
COPY apiServer.go .
COPY go.mod .


#RUN go get -d -v ./...
##RUN go install -v ./...

##RUN go mod init sampleApi.go
RUN  go build  apiServer.go

CMD /go/src/app/apiServer

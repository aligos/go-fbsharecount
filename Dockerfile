FROM golang:alpine
ADD . /go/src/github.com/aligos/echo-quote
RUN go install github.com/aligos/echo-quote
CMD ["/go/bin/echo-quote"]
EXPOSE 3000
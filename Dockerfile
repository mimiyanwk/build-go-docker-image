FROM golang:latest
WORKDIR $GOPATH/src/github.com/mimiyanwk
COPY . .

# RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["mimiyanwk"]

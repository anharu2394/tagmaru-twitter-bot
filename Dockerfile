FROM golang:latest

WORKDIR /go
ADD . /go

RUN go get -u github.com/ChimeraCoder/anaconda \
      && go get -u github.com/joho/godotenv
      CMD ["go", "run", "tweet.go"]

FROM golang:latest

WORKDIR /go
ADD . /go

RUN go get -u github.com/ChimeraCoder/anaconda && go get -u github.com/joho/godotenv

RUN apt-get update && apt-get install -y \
      busybox-static \
      && apt-get clean
      ENV TZ=Asia/Tokyo

      RUN go build tweet.go
      COPY crontab /var/spool/cron/crontabs/root

      CMD ["busybox", "crond", "-f", "-L", "/dev/stderr"]

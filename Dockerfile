FROM golang:latest
LABEL authors="rinat"


ENTRYPOINT ["top", "-b"]
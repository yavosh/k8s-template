FROM golang:alpine as build

ENV GOPATH /go
ENV USER root

CMD echo "Running on $(uname -m)"
RUN mkdir /build
COPY . /build
WORKDIR /build

RUN go build -o server ./cmd/server

FROM alpine

COPY --from=build /build/server /app/
CMD [ "/app/server" ]

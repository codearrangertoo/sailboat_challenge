# syntax=docker/dockerfile:1
FROM golang:1.15 AS build
WORKDIR /go/src/github.com/jgarland/sailboat_challenge/
COPY ./src .
RUN go build .

FROM scratch AS final
WORKDIR /root/
COPY --from=build /go/src/github.com/jgarland/sailboat_challenge/sailboat_challenge .
ENTRYPOINT ["./sailboat_challenge"]  
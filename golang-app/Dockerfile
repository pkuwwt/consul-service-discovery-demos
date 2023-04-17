FROM golang
WORKDIR /dist/

ENV GOPROXY=https://goproxy.cn

ADD main.go /dist/main.go
ADD consul /dist/consul
ADD go.sum /dist/
ADD go.mod /dist/
RUN go build


FROM debian:11
WORKDIR /dist/
COPY --from=0 /dist/example /dist/

ENTRYPOINT ["/dist/example"]

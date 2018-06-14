FROM golang:1.9-alpine

COPY root /

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN apk add --no-cache ca-certificates \
        dpkg \
        gcc \
        git \
        musl-dev \
        bash \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && chmod +x /entrypoint.sh

RUN go get github.com/derekparker/delve/cmd/dlv

WORKDIR $GOPATH
ENTRYPOINT ["/entrypoint.sh"]

CMD ["go", "run", "main.go"]

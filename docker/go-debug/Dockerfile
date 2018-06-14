FROM golang:1.10-alpine

COPY root /

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN apk add --no-cache ca-certificates \
        dpkg \
        gcc \
        git \
        musl-dev \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && chmod +x /entrypoint.sh \
    && go get github.com/derekparker/delve/cmd/dlv

WORKDIR $GOPATH
ENTRYPOINT ["/entrypoint.sh"]

CMD ["dlv", "debug", "--headless", "--listen=:2345", "--api-version=2"]

#
# Stage 1
#
FROM library/golang as builder

RUN go get github.com/tools/godep                     \
    && go get github.com/astaxie/beego                \ 
    && go get github.com/ericchiang/k8s               \
    && go get github.com/ghodss/yaml                  \
    && go get github.com/bradfitz/gomemcache/memcache \
    && go get github.com/gorilla/websocket

RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/flux-web
RUN mkdir -p $APP_DIR

ADD . $APP_DIR

WORKDIR $APP_DIR

RUN CGO_ENABLED=0 godep get -v

RUN CGO_ENABLED=0 go build -ldflags '-w -s' -o /flux-web && \
    cp -r views/ /views && \
    cp -r conf/ /conf
#
# Stage 2
#
FROM alpine:3.11
RUN adduser -D -u 1000 flux-web
COPY --from=builder /flux-web /flux-web
COPY --from=builder /views /views
COPY --from=builder /conf /conf
USER 1000
ENTRYPOINT ["/flux-web"]

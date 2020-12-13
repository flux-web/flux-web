#
# Stage 1
#
FROM library/golang as builder

ENV APP_DIR $GOPATH/src/flux-web/flux-web
RUN mkdir -p $APP_DIR

WORKDIR $GOPATH/src/flux-web/flux-web

ADD go.* $APP_DIR/

RUN go mod download

ADD . $APP_DIR

RUN CGO_ENABLED=0 go build -ldflags '-w -s' -o /flux-web && cp -r conf/ /conf && cp -r views/ /views
    
#
# Stage 2
#
FROM alpine:3.12
RUN adduser -D -u 1000 flux-web
COPY --from=builder /flux-web /flux-web
COPY --from=builder /conf /conf
COPY --from=builder /views /views
USER 1000
ENTRYPOINT ["/flux-web"]

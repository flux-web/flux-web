
FROM library/golang

# Godep for vendoring
RUN go get github.com/tools/godep

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/flux-web
RUN mkdir -p $APP_DIR

ADD . $APP_DIR

WORKDIR $GOPATH/src/flux-web

# Compile the binary and statically link
RUN CGO_ENABLED=0 godep get

CMD [ "go run main.go" ]

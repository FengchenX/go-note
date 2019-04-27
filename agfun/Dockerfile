FROM golang:alpine as agfun-build
WORKDIR /agfun
COPY ./go-note /agfun
COPY ./go-third /agfun
COPY ./go-service /agfun
RUN cd /agfun/go-note/src/agfun/auth \
        && export GOPATH=$GOPATH:/agfun/go-third \
        && export GOPATH=$GOPATH:/agfun/go-service \
        && export GOPATH=$GOPATH:/agfun/go-note \
        && GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build

FROM alpine
WORKDIR /usr/local/bin
COPY --from=agfun-build /agfun/go-note/src/agfun/auth/auth /usr/local/bin/
#COPY --from=agfun-build /titangrm/src/titan-grm/translation /usr/local/bin/translation/
RUN apk add curl

EXPOSE 8080

#ENV GRM_SERVER=data-manager \
#    GRM_SERVER_ADDRESS=:8080 \
#    GRM_SERVER_NAMESPACE=titangrm \
#    GRM_REGISTRY_ADDRESS=consul.titangrm:8500 \
#    GRM_AUTH_DB=auth_db:5432 \
#    GRM_AUTH_DB_USER=postgres \
#    GRM_AUTH_DB_PWD=otitan123 \
#    GRM_ETCD_ENDPOINTS=etcd:2379

ENTRYPOINT auth
#ENTRYPOINT titan-grm "$GRM_SERVER" -registry_address "$GRM_REGISTRY_ADDRESS" -server_address "$GRM_SERVER_ADDRESS" -server_namespace "$GRM_SERVER_NAMESPACE"

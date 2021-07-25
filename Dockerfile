FROM golang:1.14.0-alpine3.11 as builder

#RUN mkdir -p /opt/app

WORKDIR /opt/app

RUN apk --update add curl git make gcc libc-dev

COPY ./go.mod ./go.sum ./

#RUN go mod download

COPY ./ ./

#RUN GOOS=linux go build -i -v -ldflags "-X main.version=${VERSION}" ./

ENTRYPOINT go run -ldflags "-X main.version=${VERSION}" ./ --listen ${APP_LISTEN_MARKETING} --uri_api ${URI_API} --uri_static ${URI_STATIC} --host_contracts ${HOST_CONTRACTS} --pguser ${DB_USER} --pghost ${DB_HOST} --pgport ${DB_PORT} --pgdbname ${DB_NAME} --pgpassword ${DB_PASSWORD}


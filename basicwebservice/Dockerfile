From golang:1.13-alpine AS build

WORKDIR /src/
COPY webserver.go go.* /src/

RUN CGO_ENABLED=0 go build -o /bin/webserver

FROM scratch
COPY --from=build /bin/webserver /bin/webserver
ENTRYPOINT ["bin/webserver"]

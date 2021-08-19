From golang:1.16.5-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/

RUN CGO_ENABLED=0 go build -o /bin/webservice

FROM scratch
COPY --from=build /bin/webservice /bin/webservice
ENTRYPOINT ["bin/webservice"]

FROM golang:1.23 AS build

WORKDIR /app
COPY . .
RUN go mod download
RUN cd /cmd && CGO_ENABLED=0 go build -o ../bin/app

FROM debian:bullseye-slim AS tz
RUN apt-get update && apt-get install -y tzdata
RUN ln -fs /usr/share/zoneinfo/Europe/Moscow /etc/localtime && dpkg-reconfigure -f noninteractive tzdata

FROM gcr.io/distroless/static-debian11
COPY --from=build /cmd/bin /
COPY --from=build /cmd/config/env /
COPY --from=tz /etc/localtime /etc/localtime

ARG DEPLOY=true
ENV DEPLOY="${DEPLOY}"

CMD ["./cmd/main"]
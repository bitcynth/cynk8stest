FROM golang:1.13.5-alpine
ARG build_version
ARG build_time
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main -ldflags "-X main.buildVersion=${build_version} -X main.buildTime=${build_time}" .
EXPOSE 8080
CMD ["/app/main"]
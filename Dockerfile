FROM golang:alpine AS build-env
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o /responder

FROM scratch
COPY --from=build-env /responder /responder
EXPOSE 8080
CMD ["/responder"]

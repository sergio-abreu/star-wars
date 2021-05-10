FROM golang:1.16-alpine AS build
RUN apk add --update alpine-sdk
WORKDIR /src
COPY . .
RUN go mod vendor
RUN go build -o .bin/star-wars ./

FROM alpine
COPY --from=build /src/.bin/ /
ENTRYPOINT ["/star-wars"]

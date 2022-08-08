FROM golang:1.18-bullseye AS build

ENV APP_HOME /go/src/fasthttp
WORKDIR "$APP_HOME"
COPY src/ .

# ENV GO111MODULE=on
# ENV GOFLAGS=-mod=vendor

#COPY go.mod ./

RUN go mod download
RUN go mod verify
RUN go build -o fasthttp

#COPY *.go ./

#RUN go build -o /godocker

# FROM scratch

# ENV APP_HOME /go/src/fasthttp
# #RUN mkdir -p "$APP_HOME"
# WORKDIR "$APP_HOME"

# COPY --from=build "$APP_HOME"/fasthttp "$APP_HOME"

# EXPOSE 8080

CMD ["./fasthttp"]
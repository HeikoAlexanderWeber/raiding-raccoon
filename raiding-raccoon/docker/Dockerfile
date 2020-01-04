### === Builder stage. === ###
FROM golang:1.13-alpine as build-env

# Some updates.
RUN apk update && apk upgrade
RUN apk add make git ca-certificates

# First copying the dependency information and download dependencies in order to
# enable docker to cache the filesystem in this state.
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

# Add all the sourcecode and resources and execute a build.
ADD . .
RUN make build


### === This is the runtime stage. === ###
# This part of the image must be modified very carefully to not clash with security guidelines for instance.
FROM alpine:3.10 as runtime-env

# Set misc information
LABEL maintainer="Heiko Alexander Weber <heiko.a.weber@gmail.com>"
LABEL license="MIT"
USER root

# Copy the necessary artifacts.
COPY --from=build-env \
    /app/bin \
    /app
WORKDIR /app

# Executing user should never be root
RUN addgroup -S app && adduser -S app -G app
RUN chown -R app:app .
# Important!
USER app

# Printing meta information for debugging.
RUN ls -alghR
RUN env

# Starting the app here.
ENTRYPOINT [ "/app/program" ]

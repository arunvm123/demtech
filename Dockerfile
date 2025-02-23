#Build Stage
FROM golang:1.23.4 AS build-env

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# All these steps will be cached
RUN mkdir /demtech
WORKDIR /demtech

# <- COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN go build -o server .
# <- Second step to build minimal image
FROM scratch

COPY --from=build-env /demtech/server /
# ADD ./ssl ssl
ENTRYPOINT [ "./server","-config-env","true" ]
FROM golang:1.19-alpine

RUN mkdir /app

ADD . /app

# Creates an app directory to hold your app’s source code
WORKDIR /app
# Installs Go dependencies
RUN go mod download

# Builds your app with optional configuration
RUN go build -o main .

# Tells Docker which network port your container listens on
EXPOSE 8080

# Specifies the executable command that runs when the container starts
CMD ["/app/main"]

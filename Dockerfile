#FROM golang:1.17-alpine
#
#WORKDIR  ./go/src/git.paygear.ir/cloud/${PROJECT}
#
#RUN go get -v
#
#RUN go build -v -o /go/bin/chemex
#
#COPY . .
#COPY chemex /
#
#CMD ["chemex"]



# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Haj Hamid"

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY chemex /

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /build

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "/build" ]
FROM golang:1.17-bullseye

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

# Create and change to the app directory.
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o server .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/server .

COPY sample.db .
COPY frontend/build/ frontend/build/

# # Run the web service on container startup.
ENTRYPOINT ["./server"] 
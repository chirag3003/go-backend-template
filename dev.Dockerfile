FROM golang:latest

# Install libvips and libvips-dev
RUN apt update && apt install -y libvips-dev libvips && rm -rf /var/lib/apt/lists/*

# Install go air package for live reloading
RUN go install github.com/air-verse/air@latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Expose port 5000 to the outside world
EXPOSE 5000

# Command to run the executable
CMD ["air"]


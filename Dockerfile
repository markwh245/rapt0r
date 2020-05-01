FROM golang:1.14

# Maintainer
LABEL maintainer="Gabriel Dutra <dtr0x80@protonmail.com>"

# Current working directory
WORKDIR /app

# Copy all files
COPY . .

# Download dependencies
RUN go mod download

# Build app
RUN make

# Command to run the executable
ENTRYPOINT ["./jowfuzz"]

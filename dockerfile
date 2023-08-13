# Start from the golang image
FROM golang:1.20-alpine as compiler

# Set the working directory for golang image
WORKDIR /app

# Copy all project files inside the container to /app
COPY . .

# Download Go modules
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o project


#specify the base image for the Docker container as Alpine Linux
FROM alpine

# Set the working directory for our application
WORKDIR /app

# Copy all project files from golang working directory to Alpine Linux Docker container 
COPY --from=compiler ./app/project ./binary

# Specify the initial command that should be executed
ENTRYPOINT [ "./binary" ]

# To build the Image of our golang project you need to input the command:
# docker build --go-web 
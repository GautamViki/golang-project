# FROM golang

# WORKDIR /go/src/docker-app

# RUN go mod init docker-app
# RUN go mod tidy
# RUN go get github.com/go-chi/chi/v5

# COPY . .

# # Build the Go application
# RUN go build -o docker-app .

# ENTRYPOINT [ "./docker-app" ]

# Use a lightweight Go base image
FROM golang

# Set the working directory inside the container
WORKDIR /go/src/docker-app

RUN go mod init docker-app
RUN go mod tidy
RUN go get github.com/go-chi/chi/v5

# Copy all the files from your project directory into the container
COPY . .

# Build the Go application inside the container
RUN go build -o docker-app .

# Run the Go application
CMD ["./docker-app"]
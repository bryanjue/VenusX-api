# 1. Use the official Go 1.24 image
FROM golang:1.25-alpine

# 2. Install necessary tools
RUN apk add --no-cache git gcc musl-dev

# 3. Set the working directory
WORKDIR /app

# 4. Copy dependency management files
COPY go.mod go.sum ./
RUN go mod download

# 5. Instalar Air para el modo "Live Reload"
RUN go install github.com/air-verse/air@latest

# 6. Copy the rest of the application
COPY . .

# 7. Expose port 8080
EXPOSE 8080

# 8. run with air to have live server
CMD ["air"]
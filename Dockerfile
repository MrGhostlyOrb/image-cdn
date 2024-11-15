FROM golang:1.23.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env.example .env

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /richmond-paper-supply-cdn

# Run the Go binary
CMD ["/richmond-paper-supply-cdn"]

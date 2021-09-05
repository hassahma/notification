FROM golang:1.17-alpine as builder

# Set destination for COPY
WORKDIR /root

# Download Go modules
COPY go.mod /root
COPY go.sum /root
RUN go mod download

# Copy the go code
COPY . /root

# Build
RUN go build -o /marvel
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /marvel


FROM golang:1.17-alpine
RUN apk --no-cache add ca-certificates

WORKDIR /

# Copy the pre-built binary file from the previous stage
COPY --from=builder /root /marvel

EXPOSE 9091

# Run
CMD [ "go", "run", "/marvel/*.go" ]

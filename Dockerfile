# Load an image builder
FROM golang:1.17-alpine as builder

RUN apk update \
    && apk upgrade \
    && apk add --no-cache git make

# Copy everything from this directory to "/project workdir of the container"
WORKDIR /project
COPY . .

# Build our application executable under /bin folder
RUN go build -o ./bin/main main.go

# Now our application executable is compiled let's copy it to a ligher linux distro like alpine
FROM alpine:3.7

# Change our workdir from / and copy everything from our builders project directory
WORKDIR /
COPY --from=builder /project/bin/ .

#Run the application
CMD ["./main"]
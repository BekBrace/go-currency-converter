# Set the Go version as an argument
ARG GO_VERSION=1.19

# Use the specified Chainguard Go image
FROM cgr.dev/chainguard/go:latest as build

# Set the working directory in the container
WORKDIR /GO-CURRENCY-CONVERTER

# Copy the go.mod and go.sum file if you have one (you may need to create this in your project root)
COPY go.mod .
COPY go.sum .

# Copy the entire source code from the current directory to the working directory in the container
COPY . .

# Build the application, output named 'app' based on previous Dockerfile but corrected in COPY
RUN go build -o app .

# # Use the Chainguard static image for the runtime
# FROM cgr.dev/chainguard/static:latest

# Copy the built executable from the build stage
COPY --from=build /GO-CURRENCY-CONVERTER/main /main.exe

# Set the command to run the executable, using the correct executable name
CMD ["/conv"]

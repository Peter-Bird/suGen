package docker

type Docker struct{}

func (d *Docker) GetDirs(path, name string) []string {
	return []string{}
}

func (d *Docker) GetFiles(dirName, name string) map[string]string {
	return map[string]string{
		dirName + "/DOCKERFILE": d.GenDockerFile(),
	}
}

func (d *Docker) GenDockerFile() string {

	return `# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
`
}

FROM golang:1.21
WORKDIR /app
ADD go.mod go.sum ./
RUN go install ./
ENTRYPOINT tilt-example-go
FROM golang:latest
WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
# Build the app
RUN go build -v -o /usr/local/bin/app ./
EXPOSE 80

CMD ["/usr/local/bin/app"]

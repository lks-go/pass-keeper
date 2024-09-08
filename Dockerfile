FROM golang:1.22.2-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN go build -v -o /bin/server ./cmd/server
EXPOSE 9000/tcp

CMD ["/bin/server"]


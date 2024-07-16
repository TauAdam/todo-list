FROM golang:1.22.5
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-list

EXPOSE 8080
CMD ["/todo-list"]
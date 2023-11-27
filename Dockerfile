FROM golang:1.20

RUN apt-get update && apt-get install -y

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify
RUN go get -v -u github.com/golang-migrate/migrate/v4

COPY . .

RUN go build -o main ./cmd/api/

EXPOSE 4000
EXPOSE 5432

CMD ["migrate -path ./migrations -database postgres://playground:playground@db/playground?sslmode=disable up"]
CMD ["./main"]

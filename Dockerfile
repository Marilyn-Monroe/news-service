FROM golang:1.22.1 AS api

WORKDIR /compiler

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./main.go

FROM scratch AS prod

WORKDIR /app

COPY --from=api /compiler/server .

COPY /postgresql/schemas ./postgresql/schemas
COPY .env ./

EXPOSE 8080
CMD ["./server"]

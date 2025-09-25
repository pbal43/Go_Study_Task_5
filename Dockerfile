FROM golang:1.25-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o todolist cmd/to_do_list/main.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/todolist .
EXPOSE 8080
CMD ["./todolist"]
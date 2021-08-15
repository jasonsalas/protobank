###################
# BUILD CONTAINER
###################
FROM golang:alpine AS builder

RUN mkdir app
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o protobank-server ./cmd/server/main.go

###################
# RUN CONTAINER
###################
FROM alpine
WORKDIR /app
COPY --from=builder /app/protobank-server .
EXPOSE 50051

CMD [ "./app/protobank-server" ]

FROM golang:1.19-alpine3.15 AS builder
WORKDIR /home/admin1/Documents/recruitment
COPY . .
RUN go build -o /recruit

# Run stage
FROM alpine:3.15
#RUN addgroup -S appgroup && adduser -S admin1 -G appgroup


WORKDIR /app
COPY --from=builder /recruit /recruit
#COPY  /files ./files
#COPY  /INTERFACE ./INTERFACE
#COPY  /keys  ./keys

#VOLUME [ "app/files","app/INTERFACE" ]

EXPOSE 8080

ENTRYPOINT [ "/recruit" ]

#USER "1000"
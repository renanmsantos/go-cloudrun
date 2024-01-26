FROM golang:1.21 as build
WORKDIR /app
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o go-cloudrun ./cmd/main.go

FROM scratch
COPY --from=build /app/go-cloudrun .
ENTRYPOINT [ "./go-cloudrun" ]
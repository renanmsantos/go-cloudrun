FROM golang:1.21 as build
WORKDIR /app
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o cloudrun

FROM scratch
COPY --from=build /app/cloudrun .
ENTRYPOINT [ "./cloudrun" ]
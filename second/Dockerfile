FROM ubuntu as base
RUN apt-get update && apt-get install -y golang-go
ENV GO111MODULE=off
COPY . .
RUN CGO_ENABLED=0 go build -o /app .
FROM scratch
COPY --from=base /app /app
ENTRYPOINT ["/app" ]
EXPOSE 80

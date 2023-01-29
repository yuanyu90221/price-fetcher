FROM golang:1.19.4-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . ./
RUN go build -o /pricefetcher
EXPOSE 8080
CMD ["/pricefetcher"]
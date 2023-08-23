FROM golang:1.20.7-alpine


WORKDIR /usr/src/app 
COPY . .
RUN go mod tidy
RUN go build -o /usr/local/bin/app cmd/*.go
EXPOSE 3000
CMD ["sh", "-c", "/usr/local/bin/app"]


FROM golang:1.22
WORKDIR /app
ENV GOPROXY=direct
COPY go.mod go.sum ./
COPY .env ./
RUN go mod download
COPY src ./src
RUN go mod tidy
RUN go build -C src -o ./docker-app
EXPOSE 8080
CMD [ "./src/docker-app" ]
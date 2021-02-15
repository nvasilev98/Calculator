FROM golang:1.14 as build
RUN mkdir /calculator
WORKDIR /calculator
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o calculator cmd/webapp/main.go

FROM scratch
COPY --from=build /calculator/calculator /bin/calculator
CMD ["/bin/calculator"]
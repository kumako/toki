FROM golang AS builder

# Fetch Golang CLI Dependencies
RUN go get github.com/markbates/pkger/cmd/pkger

# Copy Source Code
WORKDIR /code
COPY . /code

# Static Build
RUN pkger
RUN go build -o /app -tags netgo -ldflags '-extldflags "-static"'


FROM debian
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app /app
EXPOSE 4242
ENTRYPOINT ["/app"]

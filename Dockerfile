# Initial stage: download modules
FROM golang:1.14 as builder
LABEL stage=builder

ADD go.mod go.sum /m/
RUN cd /m && go mod download

# add a non-privileged user
RUN useradd -u 10001 user

RUN mkdir -p /untypical
ADD . /untypical
WORKDIR /untypical

# Build the binary with go build
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
go build -ldflags="-w -s" -o ./bin/untypical ./

# Final stage: Run the binary
FROM alpine
LABEL stage=final

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /untypical /untypical

WORKDIR /untypical

USER user

CMD ["./bin/untypical"]
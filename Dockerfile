FROM golang:latest AS builder

RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir /build
RUN mkdir /build/myip
WORKDIR /build/myip

#RUN go get github.com/emolina82/myip
#RUN cd /build && git clone https://github.com/emolina82/myip.git
#EXPOSE 8080 - No necesario

COPY . .

RUN cd /build/myip/src/ && go mod download && go build




FROM scratch
#FROM golang:latest


COPY --from=builder /build/myip/src/myip bin/myip
COPY certs certs

ENTRYPOINT [ "./bin/myip" ]
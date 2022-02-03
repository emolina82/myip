FROM golang:latest

RUN mkdir /build
RUN mkdir /build/myip
WORKDIR /build/myip

RUN export GO111MODULE=on 
#RUN go get github.com/emolina82/myip
#RUN cd /build && git clone https://github.com/emolina82/myip.git

COPY . .

RUN cd /build/myip/src/ && go build

EXPOSE 8080

ENTRYPOINT [ "/build/myip/src/myip" ]
FROM golang:latest
ENV HOME=/root
RUN go get github.com/dvdmuckle/curl-a-joke
COPY jokes.json /root/
WORKDIR "/root"
ENTRYPOINT ["/go/bin/curl-a-joke"]
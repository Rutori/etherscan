FROM golang:1.16

WORKDIR /workdir
COPY . /workdir

RUN make

ENTRYPOINT ["bin/app"]
CMD ["-k=${APIKEY}", "-r=${RATE}","-a=${BLOCKAMOUNT}"]

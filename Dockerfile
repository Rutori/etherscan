FROM golang:1.16

WORKDIR /workdir
COPY . /workdir

RUN make

ENTRYPOINT ["bin/app"]
CMD ["-k=${APIKEY}", "-r=5","-a=100"]

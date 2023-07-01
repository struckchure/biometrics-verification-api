FROM --platform=linux/amd64 ubuntu:23.10

RUN apt-get -y update && apt-get -y upgrade
RUN apt-get -y install build-essential curl

RUN curl -OL https://go.dev/dl/go1.20.5.linux-amd64.tar.gz
RUN sha256sum go1.20.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xvf go1.20.5.linux-amd64.tar.gz

ENV PATH="${PATH}:/usr/local/go/bin"

COPY . .

RUN make install

RUN go mod download
RUN go build -o ./bv_api

CMD [ "./bv_api" ]

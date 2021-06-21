FROM golang:rc-alpine3.13

# Setup nginx
RUN apk update
RUN apk add openrc
RUN apk add nginx
RUN mkdir /etc/nginx/sites-available
RUN mkdir /etc/nginx/sites-enabled
COPY ./assets/default.conf /etc/nginx/sites-available
COPY ./assets/nginx.conf /etc/nginx/nginx.conf
RUN ln -s /etc/nginx/sites-available/default.conf /etc/nginx/sites-enabled/default.conf
RUN rc-update add nginx default

WORKDIR /go/src/app
COPY . .

RUN go mod tidy
RUN go install

CMD ["/bin/sh"]
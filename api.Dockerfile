FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app/cmd
RUN go build -o main .
RUN git clone https://github.com/vishnubob/wait-for-it.git
FROM golang:latest
RUN mkdir /app
ADD /migrations /app/
WORKDIR /app
RUN chmod +w,+x,+r /app/goosemigration.sh
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN git clone https://github.com/vishnubob/wait-for-it.git
FROM golang:1.18

WORKDIR /app

COPY dist/api .
COPY .env .

RUN chmod +x api

CMD ["./api"]
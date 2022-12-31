FROM node:16-alpine

WORKDIR /app

COPY dist/ .

RUN npm install

ENV API_URL http://localhost:3000

CMD ["npm", "start", "--", "-p", "4200"]
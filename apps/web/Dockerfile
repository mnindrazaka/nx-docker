FROM node:16-alpine

WORKDIR /app

COPY dist/ .

RUN npm install

CMD ["npm", "start", "--", "-p", "4200"]
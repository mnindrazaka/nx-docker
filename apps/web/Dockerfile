FROM node:16-alpine

WORKDIR /app

COPY /dist/apps/web/ .

RUN npm install

CMD ["npm", "start", "--", "-p", "4300"]
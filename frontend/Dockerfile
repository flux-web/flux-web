FROM node:alpine

WORKDIR /app

RUN apk add python

COPY package.json ./
COPY yarn.lock ./

RUN yarn

COPY . .

RUN npm run build

CMD npm start
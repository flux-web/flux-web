FROM node:alpine as builder

WORKDIR /app

COPY package*.json ./

RUN npm i

COPY . .

RUN npm run build

FROM nginx:alpine

COPY .docker/default.conf /etc/nginx/conf.d/default.conf

COPY --from=builder /app/dist /app
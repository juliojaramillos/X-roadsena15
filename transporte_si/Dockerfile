#Image
FROM node:latest

#API Directory
WORKDIR /usr/src/app

COPY ./transporte_si/package.json ./

RUN npm install

COPY ./transporte_si ./

EXPOSE 1001

CMD ["npm", "run", "dev"]
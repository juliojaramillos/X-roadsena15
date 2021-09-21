#Image
FROM node:latest

#API Directory
WORKDIR /usr/src/app

COPY ./movilidad_si/package.json ./

RUN npm install

COPY ./movilidad_si ./

EXPOSE 1005

CMD ["npm", "run", "dev"]
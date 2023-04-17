FROM node:16-alpine
WORKDIR /dist/

ADD package.json /dist/
RUN npm i
ADD app.js /dist/
ADD consul.js /dist/

EXPOSE 3000

ENTRYPOINT ["node", "app.js"]

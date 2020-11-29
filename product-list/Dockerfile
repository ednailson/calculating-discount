FROM node:12.14.1
ENV NODE_ENV=production

WORKDIR /usr/src/app
COPY package*.json ./
RUN yarn
COPY . .
RUN yarn tsc

EXPOSE 3333
CMD [ "node", "src/server.js" ]
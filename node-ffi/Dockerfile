FROM node:18-alpine
WORKDIR /app

COPY package.json .
COPY yarn.lock .
RUN yarn install

COPY ecosystem.config.js .
COPY index.mjs .
RUN yarn global add pm2

CMD ["pm2-runtime", "ecosystem.config.js"]
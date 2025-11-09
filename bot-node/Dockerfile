FROM node:18-alpine
WORKDIR /app

COPY package.json package-lock.json ./

# instalar uuid
RUN npm install

COPY bot.js .

CMD ["node", "bot.js"]

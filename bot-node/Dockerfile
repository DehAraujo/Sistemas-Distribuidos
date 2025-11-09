FROM node:18-alpine
WORKDIR /app

# 1. Copia arquivos de configuração (package.json e package-lock.json)
COPY package*.json ./

# 2. Instala as dependências (onde 'uuid' está)
RUN npm install

# 3. Copia o código e define o comando
COPY . .
CMD ["node", "bot.js"]
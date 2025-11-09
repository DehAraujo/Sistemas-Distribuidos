FROM node:18-alpine
WORKDIR /app

# Copia os arquivos de dependência
COPY package*.json ./

# Instala as dependências de Node.js (isto deve ser reexecutado)
RUN npm install

# Copia o restante do código
COPY . .
CMD ["node", "bot.js"]
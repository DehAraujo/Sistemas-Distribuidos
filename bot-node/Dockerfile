FROM node:18-alpine
WORKDIR /app
# 1. Copia os arquivos de configuração de dependência
COPY package*.json ./

# 2. Instala as dependências (incluindo 'uuid')
RUN npm install

# 3. Copia o código real e define o comando
COPY . .
CMD ["node", "bot.js"]
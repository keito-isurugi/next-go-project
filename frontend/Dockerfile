FROM node:22.3-alpine

WORKDIR /app/src

EXPOSE 3000

COPY . .

RUN npm install -g && pnpm install

RUN pnpm run build
CMD [ "pnpm", "nuxt", "start" ]
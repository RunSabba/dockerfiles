FROM node:23-alpine AS build 

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm ci --only=production

COPY . .


FROM node:23-alpine AS app

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/node_modules ./node_modules
COPY --from=build /usr/src/app/myapp.js ./myapp.js

EXPOSE 3000

ENV PORT=3000 \
  GREETING=Howdy \
  NAME=Dockerizer

RUN addgroup -g 1001 -S nodejs && \
    adduser -S nodejs -u 1001

USER nodejs

CMD ["node","myapp.js"]
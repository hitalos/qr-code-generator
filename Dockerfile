FROM docker.io/library/node:18-alpine AS builder

WORKDIR /var/www
COPY package.json package-lock.json .
RUN npm ci
COPY . .
RUN npm run build:all

FROM docker.io/library/node:18-alpine

WORKDIR /var/www
COPY package.json package-lock.json .
RUN npm ci --omit=dev
COPY --from=builder /var/www/public ./public
COPY . .

CMD npm start

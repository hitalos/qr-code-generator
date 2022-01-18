FROM node:16

ADD . /var/www
WORKDIR /var/www
RUN npm install --production

CMD npm start

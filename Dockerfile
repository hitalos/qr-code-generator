FROM node:6.9.1

ADD . /var/www
WORKDIR /var/www
RUN npm install --save

CMD npm start

ARG PHP_VERSION
FROM php:${PHP_VERSION}-fpm-alpine

ARG PHP_XDEBUG_VERSION
ARG PHP_SWOOLE_VERSION
ARG ALPINE_REPOSITORIES

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories \

&& docker-php-ext-install -j$(nproc) exif \
&& docker-php-ext-install -j$(nproc) mysqli \
&& docker-php-ext-install -j$(nproc) bcmath \
&& apk add --no-cache freetype-dev libjpeg-turbo-dev libpng-dev \
&& docker-php-ext-configure gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/ \
&& docker-php-ext-install -j$(nproc) gd \
&& docker-php-ext-install -j$(nproc) pdo_mysql \
&& docker-php-ext-install opcache \
&& docker-php-ext-enable opcache \
&& apk add --no-cache libxml2-dev \
&& docker-php-ext-install -j$(nproc) soap \
&& apk add --no-cache libzip-dev \
&& docker-php-ext-configure zip --with-libzip=/usr/include \
&& docker-php-ext-install zip \
&& docker-php-ext-install calendar \

&& apk add --no-cache $PHPIZE_DEPS \
&& pecl install xdebug-${PHP_XDEBUG_VERSION} \
&& docker-php-ext-enable xdebug \
&& pecl install swoole-${PHP_SWOOLE_VERSION} \
&& docker-php-ext-enable swoole \
&& echo "build success"
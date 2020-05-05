ARG PHP_VERSION
FROM php:${PHP_VERSION}-fpm-alpine

ARG PHP_XDEBUG_VERSION
ARG PHP_SWOOLE_VERSION
ARG PHP_COMPOSER_VERSION
ARG ALPINE_REPOSITORIES

COPY ./resources/php/xdebug-${PHP_XDEBUG_VERSION}.tgz /tmp/xdebug-${PHP_XDEBUG_VERSION}.tgz
COPY ./resources/php/swoole-${PHP_SWOOLE_VERSION}.tgz /tmp/swoole-${PHP_SWOOLE_VERSION}.tgz
COPY ./resources/php/composer-${PHP_COMPOSER_VERSION}.phar /usr/local/bin/composer

RUN sed -i "s/dl-cdn.alpinelinux.org/${ALPINE_REPOSITORIES}/g" /etc/apk/repositories \
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
# For gd extension
&& apk add --no-cache freetype-dev libjpeg-turbo-dev libpng-dev \
&& docker-php-ext-configure gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/ \
# For mcrypt extension
&& apk add --no-cache libmcrypt-dev libltdl \
&& docker-php-ext-configure mcrypt \
# For soap extension
&& apk add --no-cache libxml2-dev \
# For zip extension
&& apk add --no-cache zip zlib-dev \
&& docker-php-ext-configure zip --with-zlib-dir=/usr \
# Install php extensions
&& docker-php-ext-install -j$(nproc) bcmath calendar exif gd mcrypt mysqli opcache pdo_mysql soap zip \
&& apk add --no-cache $PHPIZE_DEPS \
&& pecl install /tmp/xdebug-${PHP_XDEBUG_VERSION}.tgz \
&& docker-php-ext-enable xdebug \
&& pecl install /tmp/swoole-${PHP_SWOOLE_VERSION}.tgz \
&& docker-php-ext-enable swoole \
&& chmod +x /usr/local/bin/composer \
&& composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/
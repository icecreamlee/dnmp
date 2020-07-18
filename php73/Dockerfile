ARG PHP_VERSION
FROM php:${PHP_VERSION}-fpm-alpine

ARG PHP_REDIS_VERSION
ARG PHP_XDEBUG_VERSION
ARG PHP_SWOOLE_VERSION
ARG PHP_COMPOSER_VERSION
ARG ALPINE_REPOSITORIES

COPY ./resources/php/xdebug-${PHP_XDEBUG_VERSION}.tgz /tmp/xdebug-${PHP_XDEBUG_VERSION}.tgz
COPY ./resources/php/swoole-${PHP_SWOOLE_VERSION}.tgz /tmp/swoole-${PHP_SWOOLE_VERSION}.tgz
COPY ./resources/php/redis-${PHP_REDIS_VERSION}.tgz /tmp/redis-${PHP_REDIS_VERSION}.tgz
COPY ./resources/php/composer-${PHP_COMPOSER_VERSION}.phar /usr/local/bin/composer

RUN sed -i "s/dl-cdn.alpinelinux.org/${ALPINE_REPOSITORIES}/g" /etc/apk/repositories \
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
# for gd extension
&& apk add --no-cache freetype-dev libjpeg-turbo-dev libpng-dev \
&& docker-php-ext-configure gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/ \
# for soap extension
&& apk add --no-cache libxml2-dev \
# for zip extension
&& apk add --no-cache libzip-dev \
&& docker-php-ext-configure zip --with-libzip=/usr/include \
# install php extensions
&& docker-php-ext-install -j$(nproc) bcmath calendar exif gd mysqli opcache pdo_mysql soap zip pcntl \
&& docker-php-ext-enable bcmath calendar exif gd mysqli opcache pdo_mysql soap zip pcntl \
&& apk add --no-cache $PHPIZE_DEPS \
&& pecl install /tmp/xdebug-${PHP_XDEBUG_VERSION}.tgz \
&& docker-php-ext-enable xdebug \
&& pecl install /tmp/redis-${PHP_REDIS_VERSION}.tgz \
&& docker-php-ext-enable redis \
# for install swoole
&& apk add --no-cache openssl-dev \
&& ( \
    cd /tmp && tar xzf swoole-${PHP_SWOOLE_VERSION}.tgz \
    && cd swoole-${PHP_SWOOLE_VERSION} \
    && phpize && ./configure --enable-openssl --with-openssl-dir=/opt/openssl/ --enable-http2 \
    && make -j$(nproc) && make install \
    ) \
&& docker-php-ext-enable swoole \
# install composer
&& chmod +x /usr/local/bin/composer \
&& composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/
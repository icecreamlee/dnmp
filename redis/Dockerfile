ARG REDIS_VERSION
FROM redis:${REDIS_VERSION}-alpine
COPY ./redis/redis.conf /usr/local/etc/redis/redis.conf
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
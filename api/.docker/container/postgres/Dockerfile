FROM postgres:16.3-alpine3.20

ARG DEBIAN_FRONTEND=noninteractive

ENV TZ=UTC

COPY .docker/config/postgres/postgres.conf /etc/postgres.conf

CMD [ "postgres", "-c", "config_file=/etc/postgres.conf" ]

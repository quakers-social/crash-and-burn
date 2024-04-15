#!/usr/bin/env bash

set -x
set -e
set -u

cat /home/mastodon/.rbenv/live/.env.production
. /home/mastodon/.rbenv/live/.env.production

echo "DB_HOST ${DB_HOST}"
echo "DB_PORT ${DB_PORT}"
echo "DB_NAME ${DB_NAME}"
echo "DB_USER ${DB_USER}"
echo "DB_PASS ${DB_PASS}"


export PGPASSWORD=${DB_PASS}

psql -U ${DB_USER} -h ${DB_HOST} -d ${DB_NAME} -p ${DB_PORT} -S -c "\d"

/home/mastodon/.rbenv/live/bin/bundle exec puma -C config/puma.rb
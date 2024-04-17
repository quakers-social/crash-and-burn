#!/usr/bin/env bash

set -x
set -e
set -u

pwd

cat /home/mastodon/.rbenv/live/.env.production
. /home/mastodon/.rbenv/live/.env.production

echo "DB_HOST ${DB_HOST}"
echo "DB_PORT ${DB_PORT}"
echo "DB_NAME ${DB_NAME}"
echo "DB_USER ${DB_USER}"
echo "DB_PASS ${DB_PASS}"

export PGPASSWORD=${DB_PASS}

psql -U ${DB_USER} -h ${DB_HOST} -d ${DB_NAME} -p ${DB_PORT} -S -c "\d"

ls -lah /home/mastodon/.rbenv/live/public/assets/

# rails assets:precompile || true # Dockerfile
# bundle exec rails db:setup # one time
/home/mastodon/.rbenv/live/bin/bundle exec puma -C config/puma.rb
MASTODON DOCKER CONTAINER
=========================


BUILD
-----

```bash
podman build \
	--build-arg M_VERSION=v4.2.8 \
	--ulimit nofile=10240:10240 \
	-t quakers-social/mastodon-web:latest .
```

or without cache

```bash
podman build \
	--build-arg M_VERSION=v4.2.8 \
	--ulimit nofile=10240:10240 \
	--no-cache=true \
    -t quakers-social/mastodon-web:latest .
```

IMAGE DEBUGGING
---------------

Start

```bash
podman run --rm -it -p 3000:3000 quakers-social/mastodon-web:latest
```


For investigation in the image enter:

```bash
podman run --rm -it --entrypoint=sh quakers-social/mastodon-web:latest
```

IMAGE PUSH BY HAND
------------------

```bash
podman login docker.io
LATES_VERSION=4.2.8.8
DOCKER_ACCOUNT=olafradicke
podman tag  quakers-social/mastodon-web:latest  ${DOCKER_ACCOUNT}/mastodon-web:${LATES_VERSION}
podman push ${DOCKER_ACCOUNT}/mastodon-web:${LATES_VERSION}
```



NOTES
-----

https://www.digitalocean.com/community/tutorials/how-to-install-mastodon-on-ubuntu-20-04

https://docs.joinmastodon.org/admin/config/

web-container port:3000
stream container port:4000



```
ExecStart=/usr/bin/node ./streaming
```


This configuration will be written to .env.production
Save configuration? Yes
Below is your configuration, save it to an .env.production file outside Docker:

# Generated with mastodon:setup on 2024-04-12 20:25:30 UTC

# Some variables in this file will be interpreted differently whether you are
# using docker-compose or not.

LOCAL_DOMAIN=mastodon.the-independent-friend.de
SINGLE_USER_MODE=false
SECRET_KEY_BASE=1b4ee9aba8c2d297187f8f1860bc670bf790793a795bbec3c027fe4f3ad694ee7be8e2d0763549320d5bf32e73365a5cdf0b5d4e06d6b02b27c2304a69ae1184
OTP_SECRET=14b1ebbdd322b44d337b396e0ba9e6d2a60c58619c58d143784370b69ad1bbdf232bf757f9a5fd83577b59c17f5265fba412b676ec7865c3f2c4fa831f28c0ee
VAPID_PRIVATE_KEY=z_qGhzPeqjaDz1GldcNtz0wQ5tSID3jY7-doDXTyyMQ=
VAPID_PUBLIC_KEY=BD-NlWLlExCWkFPwseSkdU9nHVEF_nzsOtoh3bhqp392AzLsi73YD-BQl-mJfWa02gA4XfhXwcqiQVsDRy0IUmE=
DB_HOST=postgres
DB_PORT=5432
DB_NAME=mastodon
DB_USER=mastodon
DB_PASS=mastodon
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=redis
S3_ENABLED=true
S3_PROTOCOL=https
S3_BUCKET=files.mastodon.the-independent-friend.de
S3_REGION=us-east-1
S3_HOSTNAME=s3.us-east-1.amazonaws.com
AWS_ACCESS_KEY_ID=XXX
AWS_SECRET_ACCESS_KEY=XXX
S3_ALIAS_HOST=files.mastodon.the-independent-friend.de
SMTP_SERVER=smtp.mailgun.org
SMTP_PORT=587
SMTP_LOGIN=mastodon
SMTP_PASSWORD=XXX
SMTP_AUTH_METHOD=plain
SMTP_OPENSSL_VERIFY_MODE=none
SMTP_ENABLE_STARTTLS=auto
SMTP_FROM_ADDRESS=Mastodon <notifications@mastodon.the-independent-friend.de>
UPDATE_CHECK_URL=
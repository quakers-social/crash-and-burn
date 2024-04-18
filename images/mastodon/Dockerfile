FROM ruby:3.3-bookworm


ARG M_VERSION=v4.2.8

ENV DEBIAN_FRONTEND="noninteractive"
ENV HOME="/home/mastodon"
ENV PATH="$HOME/.rbenv/live/bin:$HOME/node_modules/.bin:$PATH"
ENV RUBY_CONFIGURE_OPTS="--with-jemalloc rbenv install 3.2.3"
ENV RAILS_ENV=production
# ENV RAILS_SERVE_STATIC_FILES=true
# ENV RUN_DB_MIGRATIONS=true
ENV PORT=3000
# ENV LD_PRELOAD=libjemalloc.so
ENV LD_PRELOAD=/usr/lib/x86_64-linux-gnu/libjemalloc.so.2

RUN apt-get update \
	&& apt install -y --no-install-recommends \
    curl \
	wget \
	gnupg \
	apt-transport-https \
	lsb-release \
	ca-certificates \
	libldap2-dev \
	libidn11-dev \
	nodejs \
	npm \
	postgresql \
	libjemalloc2 \
	libicu-dev \
	libidn-dev \
	libpq-dev \
	libjemalloc-dev \
	zlib1g-dev \
	libgdbm-dev \
	libgmp-dev \
	libssl-dev \
	libyaml-0-2 \
	ca-certificates \
	libreadline8 \
	python3 \
	shared-mime-info \
&& rm -rf /var/lib/apt/lists/*
	# node-webpack

RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg \
    | apt-key add \
	&& echo "deb https://dl.yarnpkg.com/debian/ stable main" \
    > /etc/apt/sources.list.d/yarn.list \
	&& apt update \
    && apt install yarn

# RUN find / -name "libjemalloc*"

RUN adduser --disabled-login mastodon
WORKDIR /home/mastodon
RUN chown mastodon:mastodon -Rv /home/mastodon
USER mastodon
RUN npm install --save-dev webpack \
	&& git clone https://github.com/rbenv/rbenv.git ${HOME}/.rbenv --depth=1
WORKDIR /home/mastodon/.rbenv
RUN ${HOME}/.rbenv/src/configure \
    && make -C ${HOME}/.rbenv/src \
    && eval "$(rbenv init -)" \
    && git clone https://github.com/rbenv/ruby-build.git \
	   --depth=1 \
	   /home/mastodon/.rbenv/plugins/ruby-build \
    && gem install bundler --no-document \
    && git clone --depth 1 --branch  ${M_VERSION} \
	   https://github.com/mastodon/mastodon.git \
	   /home/mastodon/.rbenv/live
WORKDIR /home/mastodon/.rbenv/live
# RUN git checkout ${M_VERSION} \
RUN bundle config deployment 'true' \
    && bundle config without 'development test' \
    && bundle install -j$(nproc) \
	&& yarn install --pure-lockfile --production --network-timeout 600000 \
	&& yarn cache clean
COPY .env.production.example /home/mastodon/.rbenv/live/.env.production
COPY entryscript.sh /entryscript.sh
# RUN rails assets:precompile
RUN OTP_SECRET=precompile_placeholder \
    SECRET_KEY_BASE=precompile_placeholder \
	rails assets:precompile
ENTRYPOINT ["/entryscript.sh"]


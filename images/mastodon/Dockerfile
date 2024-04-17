FROM ruby:3.2-bookworm


ARG M_VERSION=v4.2.8

ENV HOME="/home/mastodon"
ENV PATH="$HOME/.rbenv/live/bin:$PATH"
ENV RUBY_CONFIGURE_OPTS="--with-jemalloc rbenv install 3.2.3"
ENV RAILS_ENV=production
# ENV RAILS_SERVE_STATIC_FILES=true
# ENV RUN_DB_MIGRATIONS=true
ENV PORT=3000
# ENV LD_PRELOAD=libjemalloc.so

RUN apt-get update
RUN apt install -y \
    curl \
	wget \
	gnupg \
	apt-transport-https \
	lsb-release \
	ca-certificates \
	libldap2-dev \
	libidn11-dev \
	postgresql

# RUN find / -name "libjemalloc*"

RUN adduser --disabled-login mastodon
WORKDIR /home/mastodon
USER mastodon
RUN git clone https://github.com/rbenv/rbenv.git ${HOME}/.rbenv --depth=1
WORKDIR /home/mastodon/.rbenv
RUN ${HOME}/.rbenv/src/configure \
    && make -C ${HOME}/.rbenv/src \
    && eval "$(rbenv init -)" \
    && git clone https://github.com/rbenv/ruby-build.git \
	   --depth=1 \
	   /home/mastodon/.rbenv/plugins/ruby-build \
    && gem install bundler --no-document \
    && git clone https://github.com/mastodon/mastodon.git \
	   /home/mastodon/.rbenv/live
WORKDIR /home/mastodon/.rbenv/live
RUN git checkout ${M_VERSION} \
    && bundle config deployment 'true' \
    && bundle config without 'development test' \
    && bundle install -j$(getconf _NPROCESSORS_ONLN)
# RUN yarn install --pure-lockfile
COPY .env.production.example /home/mastodon/.rbenv/live/.env.production

COPY entryscript.sh /entryscript.sh
# RUN yarn install --pure-lockfile
# RUN bundle exec rake mastodon:setup
ENTRYPOINT ["/entryscript.sh"]


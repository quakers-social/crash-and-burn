FROM ruby:3-bookworm

ENV HOME="/home/mastodon"
ENV PATH="$HOME/.rbenv/bin:$PATH"
ENV RUBY_CONFIGURE_OPTS="--with-jemalloc rbenv install 3.2.3"
ENV RAILS_ENV=production

RUN apt-get update
RUN apt install -y \
    curl \
	wget \
	gnupg \
	apt-transport-https \
	lsb-release \
	ca-certificates \
	libldap2-dev \
	libidn11-dev

RUN adduser --disabled-login mastodon
WORKDIR /home/mastodon
USER mastodon
RUN git clone https://github.com/rbenv/rbenv.git ${HOME}/.rbenv --depth=1
WORKDIR /home/mastodon/.rbenv
RUN cat ${HOME}/.rbenv/src/configure
RUN chmod 777 ${HOME}/.rbenv/src/configure
RUN ls -lah ${HOME}/.rbenv/src/configure
RUN ${HOME}/.rbenv/src/configure
RUN make -C ${HOME}/.rbenv/src
RUN eval "$(rbenv init -)"
RUN git clone https://github.com/rbenv/ruby-build.git ~/.rbenv/plugins/ruby-build
RUN gem install bundler --no-document
RUN git clone https://github.com/mastodon/mastodon.git live
WORKDIR /home/mastodon/.rbenv/live
RUN git checkout $(git tag -l | grep '^v[0-9.]*$' | sort -V | tail -n 1)
RUN bundle config deployment 'true'
RUN bundle config without 'development test'
RUN bundle install -j$(getconf _NPROCESSORS_ONLN)
# RUN yarn install --pure-lockfile
# RUN bundle exec rake mastodon:setup
ENTRYPOINT ["/home/mastodon/.rbenv/shims/bundle", "exec", "puma", "-C", "config/puma.rb"]


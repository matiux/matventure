FROM golang:1.12.4-stretch

ARG HOST_UID
ARG HOST_GID

ENV HOST_UID ${HOST_UID:-1000}
ENV HOST_GID ${HOST_GID:-1000}

RUN apt-get update && apt-get install -y \
        freeglut3-dev \
        git \
        git-all \
        libasound2-dev \
        libgl1-mesa-dev \
        libglu1-mesa-dev \
        mesa-common-dev \
        xorg-dev \
        zsh \
        && apt-get remove --purge -y software-properties-common \
        && apt-get autoremove -y \
        && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN mkdir /home/utente \
    && groupadd -g $HOST_GID utente \
    && useradd -u $HOST_UID --gid $HOST_GID -d /home/utente utente \
    #&& usermod -a -G www-data utente \
    && chown utente /home/utente

RUN go get golang.org/x/tools/cmd/goimports

WORKDIR /home/utente/matventure

USER utente

COPY conf/custom_zshrc /home/utente
RUN wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh || true \
    && echo 'source ~/custom_zshrc' >> /home/utente/.zshrc

# USER root

# RUN apk update \
# #    && apk add --no-cache \
# #    bash \
# #    mysql-client \
# #    rsync \
#     && apk add --virtual .ext-build-deps \
#     $PHPIZE_DEPS \
#     && pecl install -o \
#     inotify \
# #    mysql \
#     redis \
#     && docker-php-ext-enable \
#     inotify \
#     redis \
#     && apk del .ext-build-deps \
#     && pecl clear-cache \
#     && docker-php-source delete


# USER utente

# RUN echo 'alias test="./vendor/bin/simple-phpunit --exclude-group slow"' >> /home/utente/.zshrc \
#     && echo 'alias test-all="./vendor/bin/simple-phpunit"' >> /home/utente/.zshrc \
#     && echo 'alias sfcc="rm -Rf var/cache/*"' >> /home/utente/.zshrc

# COPY --chown=utente:utente ssh/* /home/utente/.ssh/
# COPY --chown=utente:utente git/.gitconfig /home/utente/.gitconfig
# RUN  chmod 600 /home/utente/.ssh/id_rsa


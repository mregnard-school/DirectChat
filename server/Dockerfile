FROM rust:1.26.0

# Creating a directory to work from
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY . .
EXPOSE 80
#running config
RUN rustup default nightly-2018-05-14 \
    && rustup update \
    && cargo update \
    && cargo install diesel_cli

ENTRYPOINT ["./entrypoint.sh"]

CMD ["cargo", "run"]

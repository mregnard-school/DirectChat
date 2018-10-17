FROM rust:latest

# Creating a directory to work from
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY . .
EXPOSE 80
#running config
RUN rustup default nightly \
    && rustup update \
    && cargo update \
    && cargo install diesel_cli \
    && diesel migration run \
    && cargo build
CMD ["cargo", "run"]

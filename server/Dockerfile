# select image
FROM rust:1.23

# create a new empty shell project
RUN USER=root cargo new --bin server
WORKDIR /server

# Set rustup 1.14.0 (1e51b07cc 2018-10-04) && cargo 1.31.0-nightly (09ce4b519 2018-10-16)
# Verify with 'rustup --version && cargo --version'
RUN rustup default nightly-2018-10-20

# Install orm cli
RUN cargo install diesel_cli



# copy over your manifests
COPY ./Cargo.lock ./Cargo.lock
COPY ./Cargo.toml ./Cargo.toml
COPY ./.env ./.env
RUN cargo build


# this build step will cache your dependencies
RUN rm src/*.rs

# copy your source tree
COPY ./src ./src

# copy migrations
COPY ./migrations ./migrations

COPY ./entryfile.sh ./entryfile.sh
ENTRYPOINT ["./entryfile.sh"]

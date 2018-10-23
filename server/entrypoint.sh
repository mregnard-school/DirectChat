#!/bin/sh
echo "Entrypoint script start"
diesel migration run
echo "migration has been run"
cargo build
echo "application build"
cargo run
echo "application run"

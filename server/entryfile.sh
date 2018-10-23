#!/bin/sh
echo "runing migrations"
diesel setup

echo "print migrations"
diesel print-schema > src/schema.rs

echo "launching application"
cargo run

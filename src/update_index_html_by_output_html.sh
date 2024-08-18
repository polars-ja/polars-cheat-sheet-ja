#!/bin/bash -eux

cp html_builder/target/output.html ./
prettier --write output.html
rm -rf index.html
cp output.html index.html

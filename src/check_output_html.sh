#!/bin/bash -eux

cp html_builder/target/output.html ./
prettier --write output.html
open output.html

#!/usr/bin/env bash

git clone --depth=1 -b master git@github.com:swagger-api/swagger-ui.git $TMP_WORK_DIR

statik -src=$TMP_WORK_DIR/dist -dest ./internal/ -package=webui

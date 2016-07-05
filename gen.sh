#!/bin/bash

dirs=$(find .files/src -type d -print0 | xargs -0)

go-bindata -prefix .files/src -pkg embedded -o embedded/embedded.go $dirs

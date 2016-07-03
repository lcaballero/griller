#!/bin/bash

dirs=$(find .files -type d -print0 | xargs -0)

go-bindata -prefix .files -pkg embedded -o embedded/embedded.go $dirs

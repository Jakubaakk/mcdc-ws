#!/bin/sh

curl -k --fail http://localhost:8080/api/health || exit 1

#!/bin/bash

./scripts/watch.sh &

present -http localhost:3999 -notes

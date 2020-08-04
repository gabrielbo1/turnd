#!/bin/bash
eval "docker build -t gabrielbo1/turnd-$TRAVIS_CPU_ARCH:${TAG} ."
#!/bin/#!/usr/bin/env bash

WORKDIR=~/go/src/github.com/insolar
echo checking insolar-api...
cd ${WORKDIR}/insolar-api || exit
git checkout 1.x.x
git pull

echo checking insolar-internal-api...
cd ${WORKDIR}/insolar-internal-api || exit
git checkout 1.x.x
git pull

echo checking insolar-observer-api...
cd ${WORKDIR}/insolar-observer-api || exit
git checkout 1.x.x
git pull

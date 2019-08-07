#!/bin/#!/usr/bin/env bash

WORKDIR=~/go/src/github.com/insolar
cd ${WORKDIR}/insolar-api || exit
git checkout 1.0.0
git pull

cd ${WORKDIR}/insolar-internal-api || exit
git checkout master
git pull

cd ${WORKDIR}/insolar-observer-api || exit
git checkout master
git pull

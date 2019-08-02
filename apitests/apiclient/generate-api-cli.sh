#!/usr/bin/env bash
# shellcheck disable=SC2086
set -x

WORKDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
SPEC_BASE_DIR=${WORKDIR}
#REPOS=( "insolar-observer-api" "insolar-internal-api" "insolar-api" )
#REPOS=( "insolar-api" )
#
#for repo_name in "${REPOS[@]}"
#do
#  if [[ ! -f ${SPEC_BASE_DIR}/${repo_name} ]]; then
#	  git clone git@github.com:insolar/${repo_name}.git ${SPEC_BASE_DIR}/${repo_name}
#	fi
#	cd ${SPEC_BASE_DIR}/${repo_name} || exit 0
#  if [[ ${repo_name} == "insolar-api" ]]; then
#    git checkout 1.0.0
#    git pull
#  else
#    git pull origin master
#  fi
#
#	npm install
#  npm run export -- --collapse
#  mkdir openapi
#  openapi-generator generate \
#      --input-spec api-exported.yaml \
#      --generator-name go \
#      --output ${SPEC_BASE_DIR}/${repo_name} \
#      --package-name insolar-api-cli \
#      --skip-validate-spec
#done

# TODO kostil
openapi-generator generate \
      --input-spec api-exported-template.yaml \
      --generator-name go \
      --output ${SPEC_BASE_DIR}/apiclient \
      --package-name apiclient \
      --model-package models \
      --skip-validate-spec

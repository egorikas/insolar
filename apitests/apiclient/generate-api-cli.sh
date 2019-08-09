#!/usr/bin/env bash
# shellcheck disable=SC2086
set -x

WORKDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
SRC_DIR=~/go/src/github.com/insolar
SPEC_BASE_DIR=${WORKDIR}
REPOS=( "insolar-api" "insolar-observer-api" "insolar-internal-api" )
#REPOS=( "insolar-api" "insolar-internal-api")

for repo_name in "${REPOS[@]}"
do
  repo_dir=${SRC_DIR}/${repo_name}
  if [[ ! -d ${repo_dir} ]]; then
    echo repository ${repo_dir} not found;
    exit 0;
  fi
	cd ${repo_dir} || exit

	npm install
  npm run export -- --collapse

  package=$(echo ${repo_name} | tr '-' '_')
  output_dir=${SPEC_BASE_DIR}/${package}
  rm -rf ${output_dir}
  openapi-generator generate \
      --input-spec api-exported.yaml \
      --generator-name go \
      --output ${output_dir} \
      --package-name ${package} \
      --skip-validate-spec

  cp ${repo_dir}/api-exported.yaml ${WORKDIR}/
  cd ${WORKDIR} || exit
done

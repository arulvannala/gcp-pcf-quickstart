#!/usr/bin/env bash

set -e


my_dir="$( cd $(dirname $0) && pwd )"
release_dir="$( cd ${my_dir} && cd ../.. && pwd )"
workspace_dir="$( cd ${release_dir} && cd .. && pwd )"
omg_tf_dir="${release_dir}/src/omg-tf"
env_dir="${workspace_dir}/env"
env_file="${workspace_dir}/omg-env-in/${env_file_name}"
env_output_dir="${workspace_dir}/omg-env-out"

pushd ${release_dir} > /dev/null
	source ci/tasks/utils.sh
popd > /dev/null
check_param ${env_file_name}

# This task may run as part of a failed job. In that case
# the env_dir will already exist and contain the state.
if [ ! -d ${env_dir} ]; then
	mkdir -p ${env_dir}
	pushd ${env_dir}
		tar zxvf ${env_file}
	popd
fi

export GOPATH=${release_dir}
export PATH=${GOPATH}/bin:${PATH}

terraform_output="${env_dir}/env.json"
terraform_config="${env_dir}/terraform.tfvars"
terraform_state="${env_dir}/terraform.tfstate"

pushd "${omg_tf_dir}"
	terraform init
	terraform get
	yes "yes" | terraform destroy --parallelism=100 -state=${terraform_state} -var-file=${terraform_config}
popd

env_file="${env_output_dir}/${env_file_name}"
pushd "${env_dir}"
	tar czvf ${env_file} .
popd
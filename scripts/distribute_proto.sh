#!/usr/bin/env bash
services=('user' 'auth' 'gateway')
gateway=('user' 'auth')
user=('auth')
auth=()
follow=()
post=()
timeline=()

for service in "${services[@]}"; do
	srv_proto_dir="${service}/proto"
	echo "Service: $service"
	eval "dependencies=(\"\${${service}[@]}\")"
	for dependency in "${dependencies[@]}"; do
		echo "Dependency: $dependency"
		cp "${dependency}"/proto/*.proto "${srv_proto_dir}"
	done
done

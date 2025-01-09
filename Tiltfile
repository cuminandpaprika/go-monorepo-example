# -*- mode: Python -*-

# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')

local_resource(
  'order-local-bin',
  'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/order ./order',
  deps=['./order/main.go', './order/start.go'],
)

# Build docker and live sync files  
docker_build_with_restart(
  'order-image',
  '.',
  entrypoint=['/app/build/order'],
  dockerfile='order/deploy/Dockerfile',
  only=[
    './build',
    './order/web',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./order/web', '/app/web'),
  ],
)

k8s_yaml('order/deploy/k8s.yaml')
k8s_resource('order', port_forwards=8000, resource_deps=['order-local-bin'])

local_resource(
  'kitchen-local-bin',
  'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/kitchen ./kitchen',
  deps=['./kitchen/main.go', './kitchen/start.go'],
)

# Build docker and live sync files  
docker_build_with_restart(
  'kitchen-image',
  '.',
  entrypoint=['/app/build/kitchen'],
  dockerfile='kitchen/deploy/Dockerfile',
  only=[
    './build',
    './kitchen/web',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./kitchen/web', '/app/web'),
  ],
)

k8s_yaml('kitchen/deploy/k8s.yaml')
k8s_resource('kitchen', port_forwards='8001:8000', resource_deps=['kitchen-local-bin'])
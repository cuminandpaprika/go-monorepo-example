# -*- mode: Python -*-

# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/order ./order'

local_resource(
  'order-local-bin',
  compile_cmd,
  deps=['./order/main.go', './order/start.go'],
  )
  
docker_build_with_restart(
  'example-go-image',
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
k8s_resource('example-go', port_forwards=8000, resource_deps=['order-local-bin'])
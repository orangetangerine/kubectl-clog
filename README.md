# clog kubectl

A `kubectl` plugin to colorize you logs

## Quick Start

```
kubectl krew install clog
# everything just like kubectl logs, just replace logs with clog
kubectl clog deploy/helloworld --tail=1 -f
```

## Feature
Some content are detected to colorizing.
* json format log with level field. e.g. `{"level":"debug"}`
* envoy format log via istio-proxy. e.g. `2023-12-26T07:01:24.212130Z     debug   envoy upstream`
* istio access log. e.g. `[2023-12-26T05:45:58.421Z] "POST /package.service/method HTTP/2" 200 ...` 


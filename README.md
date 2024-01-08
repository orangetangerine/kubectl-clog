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

<img width="753" alt="image" src="https://github.com/orangetangerine/kubectl-clog/assets/4987543/93c50646-d51e-4dfc-ba67-b0841cdd0c9b">
<img width="1490" alt="image" src="https://github.com/orangetangerine/kubectl-clog/assets/4987543/226c1532-4283-4ba2-b0ce-80d2c41e0edf">
<img width="1310" alt="image" src="https://github.com/orangetangerine/kubectl-clog/assets/4987543/97b8f391-37c3-4639-8dff-35ab0a244482">

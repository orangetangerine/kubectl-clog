
## Usage
The following assumes you have the plugin installed via

```shell
kubectl krew install clog
```

### Colorize Your kubectl logs

```shell
# usage: just like kubectl logs, use clog instead of logs
kubectl clog ...
```

## How it works
Just A wrapper to kubectl logs, with some filter writers.
# featurelabd
Daemon client for Feature Lab.

## Running the daemon
```bash
# install featurelabd to the bin folder in GOPATH
go install github.com/torresjeff/featurelabd 

# Run featurelabd
featurelabd 
```

Supported arguments for `featurelabd`:
* `featureLabUrl`:  the URL of the Feature Lab Server. Default: http://localhost:3000
* `port`: the port for the daemon to listen on. Default: 43743
* `ttlMinutes`: the ttl for the cache in minutes. The Feature Lab daemon will automatically fetch features at intervals specified by this argument. Default: 10
* `cleanUpInterval`: the interval to clean up stale entries in the cache in minutes. Default: 30

Example:

`$ featurelabd --featureLabUrl http://featurelab.com --port 5555 --ttlMinutes 20 --cleanUpInterval 15`

## Build
In `featurelabpb` package:
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative featurelabd.proto
```

## Testing
To test with `grpcurl` (Mac OS):
```bash
brew install grpcurl
```

### Get treatment
```bash
grpcurl -plaintext -d '{"app": "FeatureLab", "feature": "ShowRecommendations", "criteria": "123456"}' localhost:43743 pb.FeatureLab/GetTreatment
```

### Fetch a specific feature
```bash
grpcurl -plaintext -d '{"app": "FeatureLab", "feature": "ShowRecommendations"}' localhost:43743 pb.FeatureLab/FetchFeature
```

### Fetch features
```bash
grpcurl -plaintext -d '{"app": "FeatureLab"}' localhost:43743 pb.FeatureLab/FetchFeatures
```

# featurelabd
Daemon client for Feature Lab.


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

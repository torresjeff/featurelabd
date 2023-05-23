# featurelabd
Daemon client for Feature Lab.

The purpose of featurelabd is to reduce latency when evaluating the treatment for a feature. Instead of making a round-trip to the server
each time you want to get the treatment for a user, the request is sent to the daemon instead, which continually caches the treatment allocations.
By default, the daemon fetches the treatments for features at 10 minute intervals, but this value is configurable. The daemon uses GRPC protocol.

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

## Make calls to the daemon from your code
```go
package main

import (
	"fmt"
	"github.com/torresjeff/go-feature-lab/featurelab"
	"google.golang.org/grpc"
	"log"
)

func main() {
	apps := []string{"MyApp", "AnotherApp"}
	// Creates a client that connects to featurelabd which is listening in port 43743.
	// The daemon will pre-cache all features that belong to the specified apps.
	featureLab, conn, err := featurelab.NewFeatureLabDaemonClient(43743, apps...)
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	// Get the treatment assignment for the feature named "MyFeature" which belongs to the app "MyApp", using "userId" as the criteria
	treatment, flError := featureLab.GetTreatment("MyApp", "MyFeature", "userId")
	if flError != nil {
		log.Printf(flError.Error())
	} else {
		log.Println(fmt.Sprintf("Treatment for feature %s using criteria %s is: %+v", "MyFeature", "userId", treatment))
	}
}
```

## Development

### Build
In `featurelabpb` package:
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative featurelabd.proto
```

### Testing
To test with `grpcurl` (Mac OS):
```bash
brew install grpcurl
```

#### Get treatment
```bash
grpcurl -plaintext -d '{"app": "FeatureLab", "feature": "ShowRecommendations", "criteria": "123456"}' localhost:43743 pb.FeatureLab/GetTreatment
```

#### Fetch a specific feature
```bash
grpcurl -plaintext -d '{"app": "FeatureLab", "feature": "ShowRecommendations"}' localhost:43743 pb.FeatureLab/FetchFeature
```

#### Fetch features
```bash
grpcurl -plaintext -d '{"app": "FeatureLab"}' localhost:43743 pb.FeatureLab/FetchFeatures
```

package featurelabd

import (
	"fmt"
	"github.com/torresjeff/go-feature-lab/featurelab"
	"log"
	"time"
)

type CacheableFeatureLabClient struct {
	featureLabClient  featurelab.FeatureLab
	treatmentAssigner featurelab.TreatmentAssigner
	featureCache      FeatureCache

	// Keeps track of all the apps that have been fetched from the backend service
	apps map[string]bool
}

// GetTreatment calculates the treatment that is assigned for a criteria in a particular Feature stored in the cache.
// If the Feature is not in the cache, then featurelab.FeatureOff is returned.
// This method won't make calls to the backend service, it will only operate on features stored in cache.
// To force a cache update, you must call FetchFeatures (to cache all features) or FetchFeature (for a specific Feature).
func (f *CacheableFeatureLabClient) GetTreatment(app string, featureName string, criteria string) (featurelab.TreatmentAssignment, *featurelab.Error) {
	feature, found := f.featureCache.GetFeature(app, featureName)
	if !found {
		return featurelab.FeatureOff, nil
	}

	assignment, err := f.treatmentAssigner.GetTreatmentAssignment(feature, criteria)
	if err == featurelab.ErrInvalidTreatmentAllocation {
		return featurelab.TreatmentAssignment{}, featurelab.NewError(featurelab.ErrInvalidTreatment,
			fmt.Sprintf("Feature %s has invalid treatment allocation", getCacheKey(app, featureName)))
	}

	return assignment, nil
}

// FetchFeatures fetches all features from the Feature Lab backend service and stores them in cache, overwriting any values that already
// exist in the cache with the same Feature names.
func (f *CacheableFeatureLabClient) FetchFeatures(app string) ([]featurelab.Feature, *featurelab.Error) {
	features, err := f.featureLabClient.FetchFeatures(app)
	if err != nil {
		return nil, err
	}

	f.featureCache.PutFeatures(features)

	f.apps[app] = true

	log.Printf("Finished fetching features and cached features for app: %s\n", app)

	return features, nil
}

// FetchFeature fetches the Feature information of a Feature from the Feature Lab backend service and stores it in cache, overwriting any value
// that already exists in the cache with the same Feature Treatment.
func (f *CacheableFeatureLabClient) FetchFeature(app string, featureName string) (featurelab.Feature, *featurelab.Error) {
	feature, err := f.featureLabClient.FetchFeature(app, featureName)
	if err != nil {
		return featurelab.Feature{}, err
	}

	f.featureCache.PutFeature(app, featureName, feature)

	f.apps[app] = true

	return feature, nil
}

func (f *CacheableFeatureLabClient) UpdateCache() *featurelab.Error {
	for app, _ := range f.apps {
		_, err := f.FetchFeatures(app)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewCacheableFeatureLabClient(featureLabHost string, ttl, cleanUpInterval time.Duration) *CacheableFeatureLabClient {
	return &CacheableFeatureLabClient{
		featureLabClient:  featurelab.NewFeatureLabClient(featureLabHost),
		treatmentAssigner: featurelab.NewTreatmentAssigner(),
		featureCache:      NewDefaultFeatureCache(ttl, cleanUpInterval),
		apps:              make(map[string]bool),
	}
}

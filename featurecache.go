package featurelabd

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/torresjeff/go-feature-lab/featurelab"
	"log"
	"time"
)

type FeatureCache interface {
	GetFeature(app, name string) (featurelab.Feature, bool)
	PutFeature(app, name string, feature featurelab.Feature)
	PutFeatures(features []featurelab.Feature)
}

type defaultFeatureCache struct {
	cache *cache.Cache
}

func (d *defaultFeatureCache) GetFeature(app, name string) (featurelab.Feature, bool) {
	feature, found := d.cache.Get(getCacheKey(app, name))
	if !found {
		return featurelab.Feature{}, false
	}

	f, ok := feature.(featurelab.Feature)
	if !ok {
		panic(fmt.Sprintf("expected to find a Feature in cache, but instead found %+v", f))
	}

	log.Printf("Found Name %s in cache\n", getCacheKey(app, name))
	return f, true
}

func (d *defaultFeatureCache) PutFeature(app, name string, feature featurelab.Feature) {
	d.cache.Set(getCacheKey(app, name), feature, cache.DefaultExpiration)

	log.Printf("Cached feature: %s\n", getCacheKey(app, name))
}

func (d *defaultFeatureCache) PutFeatures(features []featurelab.Feature) {
	for _, f := range features {
		d.PutFeature(f.App, f.Name, f)
	}
}

func NewDefaultFeatureCache(ttl time.Duration, cleanUpInterval time.Duration) FeatureCache {
	return &defaultFeatureCache{
		cache: cache.New(ttl, cleanUpInterval),
	}
}

func getCacheKey(app, featureName string) string {
	return fmt.Sprintf("%s:%s", app, featureName)
}

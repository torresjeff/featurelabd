package featurelabd

import (
	"github.com/torresjeff/go-feature-lab/featurelab"
	"reflect"
	"testing"
	"time"
)

func TestNewDefaultFeatureCache(t *testing.T) {
	duration := 10 * time.Minute

	got := NewDefaultFeatureCache(duration, duration)

	if got == nil {
		t.Errorf("got nil, wantTreatment not nil")
	}

	app := "FeatureLab"
	showRecommendations := "ShowRecommendations"
	allocations := []featurelab.FeatureAllocation{
		featurelab.NewFeatureAllocation("C", 10),
		featurelab.NewFeatureAllocation("T1", 10),
	}

	feature := featurelab.NewFeature(app, showRecommendations, allocations)
	got.PutFeature(app, showRecommendations, feature)

	if f, found := got.GetFeature(app, showRecommendations); !found {
		t.Errorf("got not found, want found")
	} else if !reflect.DeepEqual(f, feature) {
		t.Errorf("got %+v, want %+v", f, feature)
	}

	nonExistingFeatureName := "NonExistingFeature"
	if f, found := got.GetFeature(app, nonExistingFeatureName); found {
		t.Errorf("found feature %+v, want not found", f)
	}

}

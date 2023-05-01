package featurelabd

import (
	"context"
	"github.com/torresjeff/featurelabd/pb"
	"github.com/torresjeff/go-feature-lab/featurelab"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type FeatureLabService struct {
	pb.FeatureLabServer
	featureLabClient featurelab.FeatureLab
}

func (fls *FeatureLabService) GetTreatment(ctx context.Context, getTreatmentRequest *pb.GetTreatmentRequest) (*pb.GetTreatmentResponse, error) {
	app, feature, criteria := getTreatmentRequest.App, getTreatmentRequest.Feature, getTreatmentRequest.Criteria
	treatment, err := fls.featureLabClient.GetTreatment(app, feature, criteria)
	if err != nil {
		return nil, toGrpcError(err)
	}

	return &pb.GetTreatmentResponse{
		Treatment: treatment.Treatment,
	}, nil
}

func (fls *FeatureLabService) FetchFeature(ctx context.Context, fetchFeatureRequest *pb.FetchFeatureRequest) (*pb.FetchFeatureResponse, error) {
	app, featureName := fetchFeatureRequest.App, fetchFeatureRequest.Feature

	feature, err := fls.featureLabClient.FetchFeature(app, featureName)
	log.Printf("Fetched feature, got: %+v, err: %+v\n", feature, err)
	if err != nil {
		return nil, toGrpcError(err)
	}

	return &pb.FetchFeatureResponse{
		Feature: &pb.Feature{
			App:         app,
			Name:        feature.Name,
			Allocations: toGrpcTreatmentAllocations(feature.Allocations),
		},
	}, nil
}

func (fls *FeatureLabService) FetchFeatures(ctx context.Context, fetchFeaturesRequest *pb.FetchFeaturesRequest) (*pb.FetchFeaturesResponse, error) {
	app := fetchFeaturesRequest.App

	features, err := fls.featureLabClient.FetchFeatures(app)
	if err != nil {
		return nil, toGrpcError(err)
	}

	return &pb.FetchFeaturesResponse{
		Features: toGrpcFeatures(features),
	}, nil
}

func toGrpcError(error *featurelab.Error) error {
	return status.Errorf(toGrpcStatusCode(error.Code), error.Error())
}

func toGrpcStatusCode(code uint32) codes.Code {
	switch code {
	case featurelab.ErrBadRequest:
		return codes.InvalidArgument
	case featurelab.ErrNotFound:
		return codes.NotFound
	case featurelab.ErrInvalidTreatment:
		fallthrough
	default:
		return codes.Internal
	}
}

func toGrpcFeatures(features []featurelab.Feature) []*pb.Feature {
	grpcFeatures := make([]*pb.Feature, len(features))
	for i, feature := range features {
		grpcFeatures[i] = &pb.Feature{
			App:         feature.App,
			Name:        feature.Name,
			Allocations: toGrpcTreatmentAllocations(feature.Allocations),
		}
	}
	return grpcFeatures
}

func toGrpcTreatmentAllocations(featureAllocations []featurelab.FeatureAllocation) []*pb.TreatmentAllocation {
	treatmentAllocations := make([]*pb.TreatmentAllocation, len(featureAllocations))
	for i, featureAllocation := range featureAllocations {
		treatmentAllocations[i] = &pb.TreatmentAllocation{
			Treatment: featureAllocation.Treatment,
			Weight:    featureAllocation.Weight,
		}
	}
	return treatmentAllocations
}

func NewFeatureLabService(featureLabClient featurelab.FeatureLab, apps ...string) pb.FeatureLabServer {
	fls := &FeatureLabService{
		featureLabClient: featureLabClient,
	}

	// Initial fetch of features to cache them
	for _, app := range apps {
		_, err := fls.featureLabClient.FetchFeatures(app)
		if err != nil {
			log.Printf("error fetching features for app %s: %s", app, err)
		}
	}

	return fls
}

package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/packages"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

// FeedService handles communication with feed-related methods of the Octopus
// API.
type FeedService struct {
	builtInFeedStats string

	services.CanDeleteService
}

// NewFeedService returns an feed service with a preconfigured client.
func NewFeedService(sling *sling.Sling, uriTemplate string, builtInFeedStats string) *FeedService {
	return &FeedService{
		builtInFeedStats: builtInFeedStats,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceFeedService, sling, uriTemplate),
		},
	}
}

// Add creates a new feed.
func (s *FeedService) Add(feed IFeed) (IFeed, error) {
	if IsNil(feed) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterFeed)
	}

	feedResource, err := ToFeedResource(feed)
	if err != nil {
		return nil, err
	}

	response, err := services.ApiAdd(s.GetClient(), feedResource, new(FeedResource), s.BasePath)
	if err != nil {
		return nil, err
	}

	return ToFeed(response.(*FeedResource))
}

// Get returns a collection of feeds based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s *FeedService) Get(feedsQuery FeedsQuery) (*Feeds, error) {
	// TODO this method is wired for /api/Spaces-1/feeds?ids=feeds-builtin
	// but the server also supports a simpler single-value at /api/Spaces-1/feeds/feeds-builtin
	// we should support that too.

	v, _ := query.Values(feedsQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*FeedResource]), path)
	if err != nil {
		return &Feeds{}, err
	}

	return ToFeeds(response.(*resources.Resources[*FeedResource])), nil
}

// GetAll returns all feeds. If none can be found or an error occurs, it
// returns an empty collection.
func (s *FeedService) GetAll() ([]IFeed, error) {
	items := []*FeedResource{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return ToFeedArray(items), err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return ToFeedArray(items), err
}

// GetByID returns the feed that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s *FeedService) GetByID(id string) (IFeed, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, "id")
	}

	path := s.BasePath + "/" + id
	resp, err := api.ApiGet(s.GetClient(), new(FeedResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(IFeed), nil
}

// GetBuiltInFeedStatistics returns statistics for the built-in feeds.
func (s *FeedService) GetBuiltInFeedStatistics() (*BuiltInFeedStatistics, error) {
	path := s.builtInFeedStats
	resp, err := api.ApiGet(s.GetClient(), new(BuiltInFeedStatistics), path)
	if err != nil {
		return nil, err
	}

	return resp.(*BuiltInFeedStatistics), nil
}

// TODO remove or rename this method in API Client v3; the first parameter wants to be an IFeed, not a PackageDescription
func (s *FeedService) SearchPackageVersions(packageDescription *packages.PackageDescription, searchPackageVersionsQuery SearchPackageVersionsQuery) (*resources.Resources[*packages.PackageVersion], error) {
	if packageDescription == nil {
		return nil, internal.CreateInvalidParameterError("SearchPackageVersions", "packageDescription")
	}

	uriTemplate, err := uritemplates.Parse(packageDescription.GetLinks()[constants.LinkSearchPackageVersionsTemplate])
	if err != nil {
		return &resources.Resources[*packages.PackageVersion]{}, err
	}

	path, err := uriTemplate.Expand(searchPackageVersionsQuery)
	if err != nil {
		return &resources.Resources[*packages.PackageVersion]{}, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*packages.PackageVersion]), path)
	if err != nil {
		return &resources.Resources[*packages.PackageVersion]{}, err
	}

	return resp.(*resources.Resources[*packages.PackageVersion]), nil
}

// TODO this method should be called SearchFeedPackageVersions for consistency, but that would be a breaking change in v2 of the client; defer to v3
func (s *FeedService) SearchFeedPackageVersions(feed IFeed, searchPackageVersionsQuery SearchPackageVersionsQuery) (*resources.Resources[*packages.PackageVersion], error) {
	if feed == nil {
		return nil, internal.CreateInvalidParameterError("SearchFeedPackageVersions", "feed")
	}

	uriTemplate, err := uritemplates.Parse(feed.GetLinks()[constants.LinkSearchPackageVersionsTemplate])
	if err != nil {
		return nil, err
	}

	path, err := uriTemplate.Expand(searchPackageVersionsQuery)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*packages.PackageVersion]), path)
	if err != nil {
		return nil, err
	}

	return resp.(*resources.Resources[*packages.PackageVersion]), nil
}

func (s *FeedService) SearchPackages(feed IFeed, searchPackagesQuery SearchPackagesQuery) (*resources.Resources[*packages.PackageDescription], error) {
	if feed == nil {
		return nil, internal.CreateInvalidParameterError("SearchPackages", "feed")
	}

	uriTemplate, err := uritemplates.Parse(feed.GetLinks()[constants.LinkSearchPackagesTemplate])
	if err != nil {
		return &resources.Resources[*packages.PackageDescription]{}, err
	}

	path, err := uriTemplate.Expand(searchPackagesQuery)
	if err != nil {
		return &resources.Resources[*packages.PackageDescription]{}, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*packages.PackageDescription]), path)
	if err != nil {
		return &resources.Resources[*packages.PackageDescription]{}, err
	}

	return resp.(*resources.Resources[*packages.PackageDescription]), nil
}

// Update modifies a feed based on the one provided as input.
func (s *FeedService) Update(feed IFeed) (IFeed, error) {
	if feed == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, "feed")
	}

	path, err := services.GetUpdatePath(s, feed)
	if err != nil {
		return nil, err
	}

	feedResource, err := ToFeedResource(feed)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), feedResource, new(FeedResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(IFeed), nil
}

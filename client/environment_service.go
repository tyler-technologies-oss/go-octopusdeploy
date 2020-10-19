package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type environmentService struct {
	sortOrderPath string
	summaryPath   string

	service
}

func newEnvironmentService(sling *sling.Sling, uriTemplate string, sortOrderPath string, summaryPath string) *environmentService {
	environmentService := &environmentService{
		sortOrderPath: sortOrderPath,
		summaryPath:   summaryPath,
	}
	environmentService.service = newService(serviceEnvironmentService, sling, uriTemplate, new(model.Environment))

	return environmentService
}

func (s environmentService) getPagedResponse(path string) ([]*model.Environment, error) {
	resources := []*model.Environment{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Environments), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Environments)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new environment.
func (s environmentService) Add(environment *model.Environment) (*model.Environment, error) {
	if environment == nil {
		return nil, createInvalidParameterError("Add", parameterEnvironment)
	}

	path, err := getAddPath(s, environment)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), environment, new(model.Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

// GetAll returns all environments. If none can be found or an error occurs, it
// returns an empty collection.
func (s environmentService) GetAll() ([]*model.Environment, error) {
	items := []*model.Environment{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the environment that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s environmentService) GetByID(id string) (*model.Environment, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Environment), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.Environment), nil
}

// GetByIDs returns the environments that match the input IDs.
func (s environmentService) GetByIDs(ids []string) ([]*model.Environment, error) {
	if len(ids) == 0 {
		return []*model.Environment{}, nil
	}

	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []*model.Environment{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName returns the environments with a matching partial name.
func (s environmentService) GetByName(name string) ([]*model.Environment, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []*model.Environment{}, err
	}

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns enironments with a matching
// partial name.
func (s environmentService) GetByPartialName(name string) ([]*model.Environment, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*model.Environment{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies an environment based on the one provided as input.
func (s environmentService) Update(environment *model.Environment) (*model.Environment, error) {
	if environment == nil {
		return nil, createInvalidParameterError(operationUpdate, parameterEnvironment)
	}

	path, err := getUpdatePath(s, environment)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), environment, new(model.Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

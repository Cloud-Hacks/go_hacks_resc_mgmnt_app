package service

import (
	"resource-service/src/repository"
	"resource-service/src/service/dto"
)

//ResourceService - Resource Service
type ResourceService struct {
}

var resourceRepository repository.ResourceRepository = repository.ResourceRepository{}

//GetResource - Service to Get Resource
func (resource *ResourceService) GetResource(data dto.RequestGetResourceDTO) (dto.ResponseGetResourceDTO, error) {

	//Calling GetListOfResources Repository
	resources, err := resourceRepository.GetResource(data)
	if err != nil {
		return dto.ResponseGetResourceDTO{}, err
	}

	return resources, nil

}

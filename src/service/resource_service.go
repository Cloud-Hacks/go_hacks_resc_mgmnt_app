package service

import (
	"resource-service/src/repository"
	"resource-service/src/service/dto"
)

//ResourceService - Resource Service
type ResourceService struct{}

//AddResource - Service to Add Resource
func (resource *ResourceService) AddResource(data dto.RequestAddResourceDTO) (string, error) {
	//Calling AddResource Repository
	msg, err := resourceRepository.AddResource(data)
	if err != nil {
		return "", err
	}

	return msg, nil
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

//GetFile and img Link - Service to Get Resource
func (resource *ResourceService) Getfile_imgLink(data dto.RequestGetLinkDTO) (dto.ResponseGetLinksDTO, error) {

	//Calling GetListOfResources Repository
	resources, err := resourceRepository.Getfile_imgLink(data)
	if err != nil {
		return dto.ResponseGetLinksDTO{}, err
	}

	return resources, nil

}

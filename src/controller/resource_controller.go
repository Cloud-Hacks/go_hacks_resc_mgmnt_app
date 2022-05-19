package controller

import (
	"resource-service/src/service"
	"resource-service/src/service/dto"
	"resource-service/utils/constants"
	logger "resource-service/utils/logging"
	"resource-service/utils/wrapper/request"
	"resource-service/utils/wrapper/response"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

//ResourceController - Resource Controller
type ResourceController struct {
}

var log *zerolog.Logger = logger.GetInstance()

var resourceService = service.ResourceService{}

//GetResource - Controller to Get Resource
func (resource *ResourceController) GetResource(c *gin.Context) {
	//Bind json according to given structure
	var data dto.RequestGetResourceDTO
	if request.CheckJSON(c, &data) {
		return // exit
	}
	//Validator to Check Request Body is correct or not
	if request.CheckRequestBodyValidator(c, data) {
		return // exit
	}

	//Calling GetListOfResources Service
	resources, err := resourceService.GetResource(data)
	if err != nil {
		//Returns Error Code 500
		response.Send500(c, constants.ERROR_IN_GET_RESOURCE, err)
		return // exit
	}
	response.SuccessWithData(c, resources)
}

//Get file and image links to return links
func (resource *ResourceController) Getfile_imgLink(c *gin.Context) {

	var data dto.RequestGetLinkDTO

	// Bind form data
	c.Bind(&data)

	//Validator to Check Request Body is correct or not
	if request.CheckRequestBodyValidator(c, data) {
		return // exit
	}

	//Calling GetListOfResources Service
	resources, err := resourceService.Getfile_imgLink(data)
	if err != nil {
		//Returns Error Code 500
		response.Send500(c, constants.ERROR_IN_GET_RESOURCE, err)
		return // exit
	}
	response.SuccessWithData(c, resources)

}

//AddResource - Controller to Add Resource
func (resource *ResourceController) AddResource(c *gin.Context) {
	//Bind json according to given structure
	var data dto.RequestAddResourceDTO
	if request.CheckJSON(c, &data) {
		return // exit
	}
	//Validator to Check Request Body is correct or not
	if request.CheckRequestBodyValidator(c, data) {
		return // exit
	}
	//Calling AddResource Service
	msg, err := resourceService.AddResource(data)
	if err != nil {

		//Returns Error Code 500
		response.Send500(c, constants.ERROR_IN_ADD_RESOURCE, err)
		return // exit
	}
	response.SuccessWithMessage(c, msg)
}

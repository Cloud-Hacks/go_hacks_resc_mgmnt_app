package repository

import (
	"fmt"
	"time"

	"resource-service/dbmodel"
	"resource-service/src/model"
	"resource-service/src/service/dto"
)

var res = []model.Resource{
	{ID: 1, Title: "SPIFEE", Category: "Tech", Status: "Published", Types: "PDF", Content: "FGY", FileLink: "/hyt.pdf", CreatedBy: 32, UpdatedBy: 54},
	{ID: 2, Title: "Devtron", Category: "CI/CD", Status: "Draft", Types: "TXT", Content: "DRE", FileLink: "/grew.txt", CreatedBy: 21, UpdatedBy: 12},
}

var img_resource = []model.Image{
	{ID: 1, ResourceId: 1, Link: "/gty.jpg", Status: "Draft", CreatedBy: 43, UpdatedBy: 23},
	{ID: 2, ResourceId: 2, Link: "/gry.png", Status: "Published", CreatedBy: 23, UpdatedBy: 12},
}

//ResourceRepository- Resource Repository
type ResourceRepository struct{}

//AddResource - Store Resource into DB
func (resourceRepository *ResourceRepository) AddResource(data dto.RequestAddResourceDTO) (string, error) {
	//To store the resource
	var resource model.Resource
	resource.Title = data.Title
	resource.Category = data.Category
	resource.Status = data.Status
	resource.Types = data.Type
	resource.Content = data.Content
	resource.FileLink = data.FileLink
	resource.CreatedBy = data.CreatedBy
	resource.CreatedAt = time.Now().UTC().Unix()
	resource.UpdatedBy = data.CreatedBy
	resource.UpdatedAt = time.Now().UTC().Unix()

	//Create the record in Resource Table
	if err := dbmodel.SetupDB().Create(&resource).Error; err != nil {
		return "", err
	}
	return "Success", nil
}

//GetResource - Get Resource based on ID from DB
func (resourceRepository *ResourceRepository) GetResource(data dto.RequestGetResourceDTO) (dto.ResponseGetResourceDTO, error) {

	for _, q := range res {
		for _, r := range img_resource {
			if q.ID == data.ID && q.ID == r.ResourceId {

				//Return response
				return dto.ResponseGetResourceDTO{
					ID:         data.ID,
					Title:      q.Title,
					Category:   q.Category,
					Status:     q.Status,
					Types:      q.Types,
					FileLink:   q.FileLink,
					Content:    q.Content,
					ImageLinks: r.Link,
					UpdatedBy:  q.UpdatedBy,
					UpdatedAt:  q.UpdatedAt,
				}, nil
			}
		}
	}
	return dto.ResponseGetResourceDTO{}, fmt.Errorf("No Data")
}

func (rescRepo *ResourceRepository) Getfile_imgLink(data dto.RequestGetLinkDTO) (dto.ResponseGetLinksDTO, error) {

	for _, q := range res {
		for _, r := range img_resource {
			if q.ID == data.ID && q.ID == r.ResourceId {

				//Return response
				return dto.ResponseGetLinksDTO{

					FileLink: q.FileLink,

					ImageLinks: r.Link,
				}, nil
			}
		}
	}
	return dto.ResponseGetLinksDTO{}, fmt.Errorf("No Data")
}

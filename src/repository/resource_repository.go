package repository

import (
	"fmt"

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

package article

import "forum/entity"

type tagListResponse struct {
	Tags []string `json:"tags"`
}

func NewTagListResponse(tags []entity.Tag) *tagListResponse {
	r := new(tagListResponse)
	for _, t := range tags {
		r.Tags = append(r.Tags, t.Tag.String)
	}
	return r
}

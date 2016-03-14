package listing

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Request struct {
	Offset int `json:"offset" form:"offset"`
	Size   int `json:"size" form:"size"`
}

func NewRequest(c *gin.Context) (*Request, map[string]string) {
	errors := make(map[string]string)

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		errors["offset"] = "Invalid value"
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		errors["size"] = "Invalid value"
	}
	if len(errors) > 0 {
		return nil, errors
	}
	return &Request{offset, size}, nil
}

func (r *Request) Validate() map[string]string {
	errors := make(map[string]string)
	if r.Offset < 0 {
		errors["offset"] = "Offset should be greater than or equal to 0."
	}
	if r.Size == 0 {
		r.Size = 10
	} else if r.Size < 0 {
		errors["size"] = "Size should be greater than 0."
	} else if r.Size > 100 {
		errors["size"] = "Size should be smaller than or equal to 100."
	}
	if len(errors) > 0 {
		return errors
	}
	return nil
}

type Result struct {
	Offset  int           `json:"offset"`
	Size    int           `json:"size"`
	Total   int64         `json:"total"`
	Results []interface{} `json:"results"`
}

func (r *Result) MarshalJSON() ([]byte, error) {
	if r.Results == nil {
		r.Results = make([]interface{}, 0)
	}

	type alias Result
	return json.Marshal(&struct {
		*alias
	}{(*alias)(r)})
}

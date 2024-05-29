package handler

import "fmt"

// CreateOpportunity
func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type %s) is required", name, typ)
}

type CreateOpportunityRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpportunityRequest) Validate() error {
	if r.Role == "" {
		errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		errParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		errParamIsRequired("salary", "Int64")
	}
	return nil
}

type UpdateOpportunityRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpportunityRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}

package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *CreateOpeningRequest) Validate() error {
	if r == nil {
		return fmt.Errorf("request body is empty")
	}

	if r.Role == "" {
		return errParamIsRequired("Role", "string")
	}

	if r.Company == "" {
		return errParamIsRequired("Company", "string")
	}

	if r.Location == "" {
		return errParamIsRequired("Location", "string")
	}

	if r.Link == "" {
		return errParamIsRequired("Link", "string")
	}

	if r.Remote == nil {
		return errParamIsRequired("Remote", "bool")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("Salary", "int64")
	}

	return nil
}

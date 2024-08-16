package vld

type Issue struct {
	Code    string
	Message string
	Value   any
}

func (issue Issue) Error() string {
	return issue.Message
}

// The `Issue` struct implements the error interface, which makes it tricky
// to serialize. This `IssueDTO` struct is required because we do need
// serialization of the issues.
type IssueDTO struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Value   any    `json:"value"`
}

type Rule func(any) (any, error)

type Validation struct {
	Tag   string
	Data  any
	Rules []Rule
}

type ValidationErrors struct {
	Errors map[string]IssueDTO `json:"errors"`
}

func (v ValidationErrors) Error() string {
	return "Validation of provided data failed"
}

func Validate(validations []Validation) error {
	errors := ValidationErrors{
		Errors: make(map[string]IssueDTO),
	}

	for _, validation := range validations {
		if _, exists := errors.Errors[validation.Tag]; exists {
			continue
		}

		data := validation.Data
		var err error

		for _, rule := range validation.Rules {
			data, err = rule(data)
			if err != nil {
				validationIssue, ok := err.(Issue)
				if ok {
					errors.Errors[validation.Tag] = IssueDTO(validationIssue)
				} else {
					errors.Errors[validation.Tag] = IssueDTO{
						Code:    CODE_UNKNOWN,
						Message: err.Error(),
					}
				}
			}
		}
	}

	if len(errors.Errors) != 0 {
		return errors
	}

	return nil
}

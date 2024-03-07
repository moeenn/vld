package vld

type Rule func(any) error

type Validation struct {
	Tag   string
	Data  any
	Rules []Rule
}

type ValidationErrors struct {
	Errors map[string]string `json:"errors"`
}

func (v ValidationErrors) Error() string {
	return "Validation of provided data failed"
}

func Validate(validations []Validation) error {
	errors := ValidationErrors{
		Errors: make(map[string]string),
	}

	for _, validation := range validations {
		if _, exists := errors.Errors[validation.Tag]; exists {
			continue
		}

		for _, rule := range validation.Rules {
			err := rule(validation.Data)
			if err != nil {
				errors.Errors[validation.Tag] = err.Error()
			}
		}
	}

	if len(errors.Errors) != 0 {
		return errors
	}

	return nil
}

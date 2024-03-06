package main

import ()

/*
TODO: Validators required
- [ ] Email
- [ ] UUID
- [ ] URL
- [ ] Date
- [ ] Time
- [ ] Required
- [ ] Min (number)
- [ ] Max (number)
- [ ] Min length
- [ ] Max length
- [ ] Greater than
- [ ] Greater than or equal
- [ ] Less than
- [ ] Less than or equal
- [ ] Before date
- [ ] Before or equal to date
- [ ] After date
- [ ] After or equal to date
- [ ] Starts with
- [ ] Doesn't start with
- [ ] Ends with
- [ ] Doesn't end with
- [ ] Enum
- [ ] Password
- [ ] Base64
- [ ] JSON
- [ ] Regex
- [ ] Same
- [ ] Array
- [ ] Min items
- [ ] Max items
*/

/*
validator := v.Struct {
	"Id": []v.Validators{ v.Required, v.Min(0) },
	"Email": []v.Validators{ v.Required, v.Email },
	"Role": []v.Validators{ v.Enum("ADMIN", "CUSTOMER") },
	"posts": []v.Validators {
		v.Array(
			v.Struct {
				"Id": []v.Validators { v.Required },
				"Title": []v.Validators { v.Required },
			}
		),
		v.MinItems(1),
	}
}

var errors Map[string][[]v.ValidationError] = v.Validate(data)
*/

/*

type LoginForm struct {
	Email string
	Password string
}

func (f LoginForm) Validate() []RuleErrors {
	validations := []v.Validation {
		{
			Tag: "email",
			Data: f.Email,
			Rules: []v.Rules {v.NonEmptyString, v.Email},
		},
  		{
			Tag: "password",
			Data: f.Password,
			Rules: []v.Rules {v.NonEmptyString, v.MinLength(8)},
		},
	}

	return v.Validate(validations)
}

*/

type Rule func(any) error

type Validation struct {
	Tag   string
	Data  any
	Rules []Rule
}

// only capture singular errors ?!!
type ValidationErrors map[string]string

func main() {
}

package main

import (
	"fmt"
)

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

func main() {
	user := User{
		Id:    10,
		Email: "admin@site.com",
	}

	fmt.Printf("%+v\n", user)
}

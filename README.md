### VLD
A data validation library for GoLang.


#### Installation

Download the library 

```bash
go get github.com/moeenn/vld
```


#### Why not `go-playground/validator`?

There is no denying that `validator` is an amazing package and is the ecosystem standard. However there are some issues with it

- Struct tags are used to define field validators. Since tags are nothing more than strings, this method can be very error-prone.
- The validation errors are structured non user-friedly way. This makes it difficult to format the validation errors in such a way that they can be displayed on the front-end clients properly.
- The mechanism for defining custom validators is not intuitive. 

`vld` attempts to solve these problems.


#### Basic usage

```go
package main

import (
	"fmt"
	v "github.com/moeenn/vld"
)

type LoginForm struct {
	Email    string
	Password string
}

func main() {
	form := LoginForm{
		Email:    "admin-site.com",
		Password: "q1w2e3r4",
	}

	validations := []v.Validation{
		{
			Tag:   "email",
			Data:  form.Email,
			Rules: []v.Rule{v.NonEmptyString, v.Email},
		},
		{
			Tag:   "password",
			Data:  form.Password,
			Rules: []v.Rule{v.NonEmptyString, v.Min(8)},
		},
	}

	err := v.Validate(validations)
	if err != nil {
		validationErrors := err.(v.ValidationErrors)
		fmt.Printf("validation errors: %v\n", validationErrors.Errors)
		return
	}

	fmt.Println("validation successful")
}
```


#### Included validators
- [x] NonEmptyString: The input must be valid non-empty string value.

```go
v.NonEmptyString
```

- [x] Length: The input string length must be exactly equal to the provided length.

```go
// input string must be exactly 10 characters in length
v.Length(10)
```

- [x] Min: The input string length must be equal or more than the provided length.

```go
// input string must be 5 characters or more in length
// input int / float must be greater than 5
v.Min(5)
```

- [x] MaxLength: The input string length must be equal or less than the provided length.

```go
// input string must be 5 characters or less in length
v.MaxLength(20)
```

- [x] MinFloat: The input floating-point number must be equal or more than the provided limit.

```go
// input number must be at least 5.0 or more
v.MinFloat(5.0)
```

- [x] MaxFloat: The input floating-point number must be equal or less than the provided limit.

```go
// input number must be 500.5 or less
v.MaxFloat(500.5)
```

- [x] MinInt: The input integer number must be equal or more than the provided limit.

```go
// input number must be at least 100 or more
v.MinInt(100)
```

- [x] MaxInt: The input integer number must be equal or less than the provided limit.

```go
// input number must be 200 or less
v.MaxInt(200)
```

- [x] GreaterThan: The input floating-point number must be greater than the provided limit.

```go
// input number must be greater than 10.5 (non-inclusive)
v.GreaterThan(10.5)
```

- [x] LessThanInt: The input integer number must be less than the provided limit.

```go
// input number must be less than 24 (non-inclusive)
v.LessThanInt(24)
```

- [x] GreaterThanInt: The input integer number must be greater than the provided limit.

```go
// input number must be greater than 100 (non-inclusive)
v.GreaterThanInt(100)
```

- [x] Email: The input must be a valid email address. 

```go
v.Email
```

- [x] UUID: The input must be a valid UUID string

```go
v.UUID
```

- [x] URL: The input must be a valid URL string

```go
v.URL
```

- [x] Password: The input must be a strong password. The following rules are applied. 
    - Minimum eight characters
    - At least one uppercase letter
    - At least one lowercase letter
    - At least one number
    - At least one special character

```go
v.Password
```

- [x] JSON: Input must be a well-formed JSON string

```go
v.JSON
```

- [x] ISODate: The input must be valid ISO timestamp according to RFC3339: [Link](https://pkg.go.dev/time#pkg-constants). 

```go
// input string must conform to the format e.g 2024-03-22T12:35:05.115Z
v.ISODate
```

- [x] Date: The input must be a valid date-only string. [Link](https://pkg.go.dev/time#pkg-constants).

```go
// input must be in format e.g. 2023-10-05
v.Date
```

- [x] Time: The input must be a valid time-only string. [Link](https://pkg.go.dev/time#pkg-constants).

```go
// input must be in 24-hours format: e.g. 10:20:00
v.Time
```

- [ ] DateEqual

- [ ] Before date

- [ ] Before or equal to date: TODO

- [ ] After date

- [ ] After or equal to date: TODO

- [ ] Before time: TODO

- [ ] Before or equal to time: TODO

- [x] StartsWith: The input string must begin with the provided prefix.

```
v.StartsWith("data:")
```


- [x] DoesntStartWith: The input string must not start with the provided prefix.

```go
v.DoesntStartWith("mysql")
```

- [x] EndsWith: The input string must end with the provided suffix.

```go
v.EndsWith("example")
```


- [x] DoesntEndWith: The input string must not end with the provided prefix.

```go
v.DoesntEndWith("sample")
```


- [x] Enum: The input string must be equal to one of the provided valid values.

```go
// values are provided in variadic fashion
v.Enum("Value One", "Value Two", "Value Three")
```


- [x] Regexp: The input value must satisfy the provided regular expression.

```go
v.Regexp("^hello$")
```


- [x] Same: The input value must be the same as the required input. This validator can be used to confirm passwords.

```go
import (
	"fmt"
	v "github.com/moeenn/vld"
)

type RegisterForm struct {
	Password        string
	ConfirmPassword string
}

func main() {
	form := RegisterForm{
		Password:        "A832KCN284506b@",
		ConfirmPassword: "A832KCN284506b@",
	}

	validations := []v.Validation{
		{
			Tag:   "password",
			Data:  form.Password,
			Rules: []v.Rule{v.NonEmptyString, v.Password},
		},
		{
			Tag:   "confirm_password",
			Data:  form.ConfirmPassword,
			Rules: []v.Rule{v.Same("Password", form.Password)},
		},
	}

	err := v.Validate(validations)
	if err != nil {
		validationErrors := err.(v.ValidationErrors)
		fmt.Printf("validation errors: %v\n", validationErrors.Errors)
		return
	}

	fmt.Println("validation successful")
}
```

- [ ] Positive: TODO

- [ ] Negative: TODO

- [ ] RangeBetween: TODO

- [ ] Latitude: TODO

- [ ] Longitude: TODO

- [ ] Array: TODO

- [ ] MinItems: TODO
 
- [ ] MaxItems: TODO


#### Custom validators
In `vld` validators are plain functions. They can be defined as follows.

```go
type ExampleForm struct {
	Slug string
}

func Slug(input any) (any, error) {
	err := errors.New("The input must be a valid slug")
	asString, ok := input.(string)
	if !ok {
		return nil, err
	}

	if strings.Contains(asString, "_") {
		return nil, err
	}
	return asString, nil
}

func main() {
	form := ExampleForm{
		Slug: "some-slug-here",
	}

	validations := []v.Validation{
		{
			Tag:   "slug",
			Data:  form.Slug,
			Rules: []v.Rule{v.NonEmptyString, Slug}, // notice the user-defined rule
		},
	}

	err := v.Validate(validations)
	if err != nil {
		validationErrors := err.(v.ValidationErrors)
		fmt.Printf("validation errors: %v\n", validationErrors.Errors)
		return
	}
	
	// the input data is valid
    ...	
}
```

If the custom validator function required additional arguments, they can be defined as follows.

```go
func StartsWith(prefix string) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The input must start with '%s'", prefix)
		asString, ok := input.(string)
		if !ok || !strings.HasPrefix(asString, prefix) {
			return nil, err
		}
		return asString, nil
	}
}
```

### VLD
A data validation library for GoLang.


#### Why not `go-playground/validator`?

There is no denying that `validator` is an amazing package and is the ecosystem standard. However there are some issues with it

- Struct tags are used to define field validators. Since tags are nothing more than strings, this method can be very error-prone.
- The validation errors are structured non user-friedly way. This makes it difficult to format the validation errors in such a way that they can be displayed on the front-end clients properly.
- The mechanism for defining custom validators is not intuitive. 

`vld` attempts to solve these problems.


#### Installation

Download the library 

```bash
go get github.com/moeenn/vld
```


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

	if err := v.Validate(validations); err != nil {
		validationErrors := err.(v.ValidationErrors)
		fmt.Printf("validation errors: %v\n", validationErrors.Errors)
		return
	}

	fmt.Println("validation successful")
}
```


#### Included validators

|                             Validator | Description                                                                                                                                                                                                                           |
| ------------------------------------: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
|                      `NonEmptyString` | Check if provided input is a non-empty string                                                                                                                                                                                         |
|                         `Length(int)` | Check if the provided input is a string and its length is equal to the provided length.                                                                                                                                               |
|         `Min(int \| float \| string)` | If the provided number is an `int` / `float(64)`, check input is greater than or equal to the target. If the provided input is a `string`, check its length is more than or equal to the target.                                      |
|         `Max(int \| float \| string)` | If the provided number is an `int` / `float(64)`, check input is less than or equal to the target. If the provided input is a `string`, check its length is less than or equal to the target.                                         |
| `GreaterThan(int \| float \| string)` | If the provided number is an `int` / `float(64)`, check input is more than (but not equal) to the target. If the provided input is a `string`, check its length is more than (but not equal) to the target.                           |
|    `LessThan(int \| float \| string)` | If the provided number is an `int` / `float(64)`, check input is less than (but not equal) to the target. If the provided input is a `string`, check its length is less than (but not equal) to the target.                           |
|                               `Email` | Check if the provide input is a valid email address                                                                                                                                                                                   |
|                  `HasPrefix(string)` | Check if the provided input is a valid string and starts with the provided substring.                                                                                                                                                 |
|             `NotHasPrefix(string)` | Check if the provided input is a valid string and doesn't starts with the provided substring.                                                                                                                                         |
|                    `HasSuffix(string)` | Check if the provided input is a valid string and ends with the provided substring.                                                                                                                                                   |
|               `NotHasSuffix(string)` | Check if the provided input is a valid string and ends with the provided substring.                                                                                                                                                   |
|                        `Equals(string)` | Check if the provided input is the same as the target input.                                                                                                                                                                          |
|                     `Enum(...string)` | Check if the provided input matches any of the listed enumerations values.                                                                                                                                                            |
|                                 `URL` | Check if the provided input is a valid string and a valid URL.                                                                                                                                                                        |
|                      `Regexp(string)` | Check if the provided input is a valid string and matches the required regular expression.                                                                                                                                            |
|                                `UUID` | Check if the provided input is a valid string and a valid UUID.                                                                                                                                                                       |
|                            `Password` | Check if the provided input is a valid string and a reasonably strong password. Password rules <br>- Minimum eight characters<br>- At least one uppercase letter<br>- One lowercase letter<br>- One number<br>- One special character |
|                                `JSON` | Check if the provided code is a valid string and a valid json.                                                                                                                                                                        |
|                            `DateTime` | Check if the provided input is a valid string and a valid ISO timestamp according to RFC3339: [Link](https://pkg.go.dev/time#pkg-constants).                                                                                          |
|                                `Date` | Check if the provided input is a valid date-only string. Date string must be in format e.g. 2023-10-05. [Link](https://pkg.go.dev/time#pkg-constants).                                                                                |
|                                `Time` | Check if the provided input is a valid string and a valid time-only string. Time string must be in 24-hours format: e.g. 10:20:00. [Link](https://pkg.go.dev/time#pkg-constants).                                                     |
|                `DateEqual(time.Time)` | Check if the provided date is a date equal to the target date.                                                                                                                                                                        |
|               `DateBefore(time.Time)` | Check if the provided input is a date before (but not equal) to the target date.                                                                                                                                                      |
|                `DateAfter(time.Time)` | Check if the provided input is a date after the target date. If inclusive is set to true, target date will be included.                                                                                                               |
|                            `Latitude` | Check if the provided input a valid map latitude value.                                                                                                                                                                               |
|                           `Longitude` | Check if the provided input a valid map longitude value.                                                                                                                                                                              |


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

	if err := v.Validate(validations); err != nil {
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


#### TODO

- [ ] Extend LessThan: Add option to check BeforeTime
- [ ] Extend Equals: Add option to check ExactTime
- Array
	- [ ] Extend Min to allow checking MinItems 
	- [ ] Extend Max to allow checking MaxItems

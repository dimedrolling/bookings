package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

//Form a custom struct with url.Values
type Form struct {
	url.Values
	Errors errors
}

//New initializes a Form
func New(data url.Values) *Form {
	return &Form{
		data,
		map[string][]string{},
	}
}

//Valid returns true if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//Has checks if the field exist in request
func (f *Form) Has(field string, r *http.Request) bool {
	x := f.Get(field)
	if x == "" {
		return false
	}
	return true
}

//Required checks required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blanc")
		}
	}
}

//MinLength checks for min length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

//IsEmail checks is email valid
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}

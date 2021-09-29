package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/smth", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Form supposed to be valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/smth", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form is valid while has to be not")
	}

	postedData := url.Values{}
	postedData.Add("a", "1")
	postedData.Add("b", "2")
	postedData.Add("c", "3")

	r = httptest.NewRequest("POST", "/smth", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Form should be invalid")
	}

}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/smth", nil)
	form := New(r.Form)
	ok := form.Has("smth", r)
	if ok {
		t.Error("This form has no attribute")
	}

	postedData := url.Values{}
	postedData.Add("smth", "dff")

	form = New(postedData)
	ok = form.Has("smth", r)
	if !ok {
		t.Error("This form has attribute, but shows like not, got")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/smth", nil)
	form := New(r.Form)
	form.IsEmail("smth")
	if form.Valid() {
		t.Error("This form has no email at all")
	}

	postedData := url.Values{}
	postedData.Add("smth1", "dff@yandex.ru")
	postedData.Add("smth2", "90_12")

	form = New(postedData)
	form.IsEmail("smth1")
	if !form.Valid() {
		t.Error("This form has good form email, but shows like not")
	}

	form.IsEmail("smth2")
	if form.Valid() {
		t.Error("This form has no good form email, but shows like it is")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/smth", nil)
	form := New(r.Form)
	if form.MinLength("smth", 1, r) {
		t.Error("This form has no field at all")
	}

	isError := form.Errors.Get("smth")
	if isError == "" {
		t.Error("Should have an error but didnt get")
	}

	postedData := url.Values{}
	postedData.Add("smth1", "123")
	postedData.Add("smth2", "90")

	form = New(postedData)
	if !form.MinLength("smth1", 3, r) {
		t.Error("This form has good length, but shows like not")
	}
	isError = form.Errors.Get("smth1")
	if isError != "" {
		t.Error("Should not have an error but didnt get")
	}

	if form.MinLength("smth2", 3, r) {
		t.Error("This form has not enough length, but shows otherwise")
	}
}

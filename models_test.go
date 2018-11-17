package gofcourse

import (
	"testing"
)

func TestCourse(t *testing.T) {
	expected_title := "Laber"
	c := Course{Title: expected_title + " "}
	c.Normalize()
	if title := c.Title; title != "Laber" {
		t.Errorf("Title %s != %s", title, expected_title)
	}
}

func TestPerson(t *testing.T) {
	p := Person{FirstName: "Aline", LastName: "Alma"}
	p.Normalize()
	if full_name := p.FullName(); full_name != "Aline Alma" {
		t.Errorf("FullName %s != Aline Alma", full_name)
	}
}

func TestContact(t *testing.T) {
	raw_phone := "0791234567"
	phone := "+41 79 123 45 67"
	c := Contact{Mobile: []string{raw_phone}}
	c.Normalize()
	if mobile := c.Mobile[0]; mobile != phone {
		t.Errorf("Mobile %s != %s", mobile, phone)
	}
}

package gofcourse

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"strings"
	"time"
)

type Contact struct {
	Mobile    []string
	Phone     []string
	Email     []string
	Emergency []string
}

type Person struct {
	FirstName string
	LastName  string
	Address   string
	City      string
	Plz       string
	Country   string
	Birthdate time.Time
	Gender    rune
	Contact   Contact
	Notes     []string
}

type Course struct {
	Title        string
	FirstDate    time.Time
	LastDate     time.Time
	Number       uint
	Code         string
	Notes        []string
	Participants []Person
	Waitlist     []Person
}

func (c *Course) Normalize() error {
	c.Title = strings.Trim(c.Title, " ")
	return nil
}

func (p *Person) Normalize() error {
	switch p.Gender {
	case rune('m'):
	case rune('f'):
	case rune('o'):
	case rune(0):
	default:
		return fmt.Errorf("invalid gender '%v'", p.Gender)
	}
	return nil
}

func (p *Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) Id() string {
	return strings.Replace(p.FullName(), " ", "_", -1)
}

func (p *Person) Compare(p2 *Person) bool {
	return p.FullName() < p2.FullName()
}

func (c *Contact) Normalize() error {
	for _, list := range [][]string{c.Mobile, c.Phone, c.Emergency} {
		for i, v := range list {
			parsed, err := phonenumbers.Parse(v, "CH")
			if err != nil {
				return err
			}
			if ok := phonenumbers.IsValidNumber(parsed); !ok {
				return fmt.Errorf("invalid phone number %s", v)
			}
			c.Mobile[i] = phonenumbers.Format(parsed, phonenumbers.INTERNATIONAL)
		}
	}
	return nil
}

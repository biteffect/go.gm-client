package gmapi

import (
	"errors"
	"github.com/biteffect/go.gm-fin"
	"strings"
)

type GmBalance struct {
	Balance   gmfin.Amount `xml:"balance,attr"`
	Overdraft gmfin.Amount `xml:"overdraft,attr"`
}

type GmAttribute struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type GmVerify struct {
	Code      int           `xml:"code,attr"`
	Attribute []GmAttribute `xml:"attribute"`
}

type GmPayment struct {
	ID            string        `xml:"id,attr"`
	Sum           int64         `xml:"sum,attr"`
	Check         int           `xml:"check,attr"`
	Service       int           `xml:"service,attr"`
	Source        string        `xml:"source,attr"`
	Account       string        `xml:"account,attr"`
	Date          string        `xml:"date,attr"`
	Delayed       int           `xml:"delayed,attr"`
	TerminalVPSID string        `xml:"terminal-vps-id,attr"`
	Attribute     []GmAttribute `xml:"attribute"`
}

func (p *GmPayment) Validate() error {
	out := make([]string, 0)
	if len(out) > 0 {
		return errors.New(strings.Join(out, "; "))
	}
	return nil
}

type GmStatus struct {
	ID         string        `xml:"id,attr"`
	State      int           `xml:"state,attr"`
	Substate   int           `xml:"substate,attr"`
	Code       int           `xml:"code,attr"`
	Final      int           `xml:"final,attr"`
	Trans      string        `xml:"trans,attr"`
	Attributes []GmAttribute `xml:"attribute"`
}

func (s *GmStatus) Attribute(key string) *GmAttribute {
	if s.Attribute == nil {
		return nil
	}
	for _, v := range s.Attributes {
		if v.Name == key {
			return &v
		}
	}
	return nil
}

type GmAdvancedInput struct {
	Key        string `xml:"key,attr"`
	Title      string `xml:"title,attr"`
	Value      string `xml:"value,attr"`
	ValueTitle string `xml:"valueTitle,attr"`
}

type GmAdvancedData struct {
	Input []GmAdvancedInput `xml:"input"`
}

type GmAdvanced struct {
	Code    int            `xml:"code,attr"`
	Service int            `xml:"service,attr"`
	Data    GmAdvancedData `xml:"attribute"`
}

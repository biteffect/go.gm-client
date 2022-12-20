package gmapi

import (
	"github.com/biteffect/go.gm-fin"
)

type Balance struct {
	Amount    gmfin.Amount `xml:"balance,attr"`
	Overdraft gmfin.Amount `xml:"overdraft,attr"`
}

type Attribute struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type withAttribute struct {
	Attributes []Attribute `xml:"attribute"`
}

func (box *withAttribute) Attribute(key string) *Attribute {
	if box.Attribute == nil {
		return nil
	}
	for _, v := range box.Attributes {
		if v.Name == key {
			return &v
		}
	}
	return nil
}

type Status struct {
	ID       string `xml:"id,attr"`
	State    int    `xml:"state,attr"`
	Substate int    `xml:"substate,attr"`
	Sode     int    `xml:"code,attr"`
	Final    int    `xml:"final,attr"`
	Trans    string `xml:"trans,attr"`
	withAttribute
}

func (s Status) IsFinal() bool {
	return s.State == StatusSuccess ||
		s.State == StatusError ||
		s.State == -2
}

func (s Status) IsSuccess() bool {
	return s.State == StatusSuccess
}

func (s Status) IsError() bool {
	return s.State == StatusError
}

type VerifyStatus struct {
	Code int `xml:"code,attr"`
	withAttribute
}

func (s VerifyStatus) IsOk() bool {
	return s.Code == 0
}

type PaymentOptions struct {
	Check         int         `xml:"check,attr"`
	Source        string      `xml:"source,attr"`
	Delayed       int         `xml:"delayed,attr"`
	TerminalVPSID string      `xml:"terminal-vps-id,attr"`
	Attribute     []Attribute `xml:"attribute"`
}

func (o *PaymentOptions) apply(p *gmPayment) {
	if o == nil || p == nil {
		return
	}
	if len(o.Source) > 0 {
		p.Source = o.Source
	} else {
		p.Source = GmApiSourceGate
	}
	if o.Attribute != nil {
		if p.Attribute == nil {
			p.Attribute = make([]Attribute, 0)
		}
		for _, attr := range o.Attribute {
			p.Attribute = append(p.Attribute, attr)
		}
	}
}

package gmapi

import (
	"errors"
	"github.com/nooize/go-assist"
	"github.com/nooize/go-assist/env"
	"log"
	"strings"
)

const gmTimeFormat = "2006-01-02T15:04:05Z0700"

type gmPayment struct {
	ID            string      `xml:"id,attr"`
	Sum           int64       `xml:"sum,attr"`
	Check         int         `xml:"check,attr"`
	Service       int         `xml:"service,attr"`
	Source        string      `xml:"source,attr"`
	Account       string      `xml:"account,attr"`
	Date          string      `xml:"date,attr"`
	Delayed       int         `xml:"delayed,attr"`
	TerminalVPSID string      `xml:"terminal-vps-id,attr"`
	Attribute     []Attribute `xml:"attribute"`
}

func (p *gmPayment) Validate() error {
	out := make([]string, 0)
	if len(out) > 0 {
		return errors.New(strings.Join(out, "; "))
	}
	return nil
}

// ClientInstance ClientInstance
var apiInst *Client

func init() {
	point := env.GetInt(EnvKeyGmApiPoint, 0)
	if point == 0 {
		return
	}

	url := env.GetUrl(EnvKeyGmApiUrl, DefaultGmApiUrl) // DefaultGmApiUrl
	cert, err := assist.LoadPemCertificate(env.GetStr(EnvKeyGmCertPath, ""))
	if err != nil {
		log.Panicf("GM Api certificate load eror : %s", err.Error())
	}
	apiInst, _ = NewGmClient(*url, point, *cert)
}

package gmapi

import (
	"bytes"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	gmfin "github.com/biteffect/go.gm-fin"
	"github.com/nooize/go-logz"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client GM API client
type Client struct {
	url    url.URL
	point  int
	client *http.Client
	certs  *x509.CertPool
	logger logz.Logger
}

// SetLogger set logger for dump all requests & responses
func (g *Client) SetLogger(l logz.Logger) {
	if g != nil && l != nil {
		g.logger = l
	}
}

func (g *Client) GetBalance() (*Balance, error) {
	req := struct {
		XMLName xml.Name `xml:"request"`
		Point   int      `xml:"point,attr"`
		Balance string   `xml:"balance"`
	}{
		Point: g.point,
	}

	resp := struct {
		XMLName xml.Name `xml:"response"`
		Balance struct {
			Balance   int `xml:"balance,attr"`
			Overdraft int `xml:"overdraft,attr"`
		} `xml:"balance"`
	}{}

	if err := g.callApi(&req, &resp); err != nil {
		return nil, err
	}
	return &Balance{
		Amount:    gmfin.AmountFromCents(resp.Balance.Balance),
		Overdraft: gmfin.AmountFromCents(resp.Balance.Overdraft),
	}, nil
}

func (g *Client) Verify(service int, account string, attrs []Attribute) (*VerifyStatus, error) {
	req := struct {
		XMLName xml.Name `xml:"request"`
		Point   int      `xml:"point,attr"`
		Verify  struct {
			Service int    `xml:"service,attr"`
			Account string `xml:"account,attr"`
			withAttribute
		} `xml:"verify"`
	}{
		Point: g.point,
	}
	req.Verify.Service = service
	req.Verify.Account = account
	req.Verify.Attributes = attrs

	resp := struct {
		XMLName xml.Name     `xml:"response"`
		Result  VerifyStatus `xml:"result"`
	}{}

	if err := g.callApi(&req, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

func (g *Client) GetStatus(id string) (*Status, error) {
	req := struct {
		XMLName xml.Name `xml:"request"`
		Point   int      `xml:"point,attr"`
		Status  struct {
			ID string `xml:"id,attr"`
		} `xml:"status"`
	}{
		Point: g.point,
	}
	req.Status.ID = id

	resp := struct {
		XMLName xml.Name `xml:"response"`
		Result  Status   `xml:"result"`
	}{}

	if err := g.callApi(&req, &resp); err != nil {
		return nil, err
	}
	if resp.Result.State == -2 {
		return nil, nil
	}
	return &resp.Result, nil
}

func (g *Client) Payment(id string, service int, amount gmfin.Amount, account string, opt *PaymentOptions) (*Status, error) {

	req := struct {
		XMLName xml.Name  `xml:"request"`
		Point   int       `xml:"point,attr"`
		Payment gmPayment `xml:"payment"`
	}{
		Point: g.point,
		Payment: gmPayment{
			ID:      id,
			Sum:     amount.InCents(),
			Account: account,
			Check:   1,
			Service: service,
			Source:  GmApiSourceGate,
			Date:    time.Now().Format(gmTimeFormat),
		},
	}

	opt.apply(&req.Payment)

	resp := struct {
		XMLName xml.Name `xml:"response"`
		Result  Status   `xml:"result"`
	}{}

	if err := g.callApi(&req, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

/*
func (g *Client) Payment(p *Payment) (*Status, error) {
	req := struct {
		XMLName xml.Name  `xml:"request"`
		Point   int       `xml:"point,attr"`
		Payment GmPayment `xml:"payment"`
	}{
		Point:   g.point,
		Payment: *p,
	}

	resp := struct {
		XMLName xml.Name `xml:"response"`
		Result  GmStatus `xml:"result"`
	}{}

	if err := g.callApi(&req, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}
*/

func (g *Client) Advanced(service int, fn string, attrs []Attribute) (*GmAdvanced, error) {
	req := struct {
		XMLName  xml.Name `xml:"request"`
		Point    int      `xml:"point,attr"`
		Advanced struct {
			Service   int         `xml:"service,attr"`
			Function  string      `xml:"function,attr"`
			Attribute []Attribute `xml:"attribute"`
		} `xml:"advanced"`
	}{
		Point: g.point,
	}
	req.Advanced.Service = service
	req.Advanced.Function = fn
	req.Advanced.Attribute = attrs

	resp := struct {
		XMLName xml.Name   `xml:"response"`
		Result  GmAdvanced `xml:"result"`
	}{}

	if err := g.callApi(&req, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

// internal methods

func (g *Client) callApi(request interface{}, response interface{}) error {

	callBody, err := xml.Marshal(request)
	if err != nil {
		return err
	}

	defer func(v string) {
		if g.logger != nil {
			go g.logger.Debug().Send(v)
		}
	}(string(callBody))

	httpResp, err := g.client.Post(g.url.String(), "text/xml", bytes.NewReader(callBody))
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	callBody = append(callBody, []byte("\n\n")...)
	callBody = append(callBody, respBody...)

	if strings.HasPrefix(string(respBody), "<error>") {
		v := struct {
			XMLName xml.Name `xml:"error"`
			Error   string   `xml:",chardata"`
		}{}
		err = xml.Unmarshal([]byte(respBody), &v)
		if err != nil {
			return err
		}
		return fmt.Errorf("GM SG error: %s", v.Error)
	}

	err = xml.Unmarshal(respBody, response)
	if err != nil {
		return err
	}

	return nil
}

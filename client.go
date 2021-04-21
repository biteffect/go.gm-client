package gmapi

import (
	"bytes"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client GM API client
type Client struct {
	url    url.URL
	point  int
	client *http.Client
	certs  *x509.CertPool
}

// API client public methods
// check balance
func (g *Client) Balance() (*GmBalance, error) {
	req := struct {
		XMLName xml.Name `xml:"request"`
		Point   int      `xml:"point,attr"`
		Balance string   `xml:"balance"`
	}{
		Point: g.point,
	}

	resp := struct {
		XMLName xml.Name  `xml:"response"`
		Balance GmBalance `xml:"balance"`
	}{}

	if err := g.callApi(&req, &resp); err != nil {
		return nil, err
	}

	return &resp.Balance, nil
}

func (g *Client) Verify(service int, account string, attrs []GmAttribute) (*GmVerify, error) {
	req := struct {
		XMLName xml.Name `xml:"request"`
		Point   int      `xml:"point,attr"`
		Verify  struct {
			Service   int           `xml:"service,attr"`
			Account   string        `xml:"account,attr"`
			Attribute []GmAttribute `xml:"attribute"`
		} `xml:"verify"`
	}{
		Point: g.point,
	}
	req.Verify.Service = service
	req.Verify.Account = account
	req.Verify.Attribute = attrs

	resp := struct {
		XMLName xml.Name `xml:"response"`
		Result  GmVerify `xml:"result"`
	}{}

	if err := g.callApi(&req, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

func (g *Client) Status(id string) (*GmStatus, error) {
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
		Result  GmStatus `xml:"result"`
	}{}

	if err := g.callApi(&req, &resp); err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

func (g *Client) Payment(p *GmPayment) (*GmStatus, error) {
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

func (g *Client) Advanced(service int, fn string, attrs []GmAttribute) (*GmAdvanced, error) {
	req := struct {
		XMLName  xml.Name `xml:"request"`
		Point    int      `xml:"point,attr"`
		Advanced struct {
			Service   int           `xml:"service,attr"`
			Function  string        `xml:"function,attr"`
			Attribute []GmAttribute `xml:"attribute"`
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

	httpBody, err := xml.Marshal(request)
	if err != nil {
		return err
	}

	httpResp, err := g.client.Post(g.url.String(), "text/xml", bytes.NewReader(httpBody))
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

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

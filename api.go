package gmapi

import (
	"crypto/tls"
	"errors"
	"github.com/nooize/go-assist/env"
	"log"
	"net/http"
	"net/url"
	"time"
)

// ClientInstance ClientInstance
var apiInst *Client

func init() {
	point := env.GetInt(EnvKeyGmApiPoint, 0)
	if point == 0 {
		return
	}

	url := env.GetUrl(EnvKeyGmApiUrl, DefaultGmApiUrl) // DefaultGmApiUrl
	cert, err := tls.LoadX509KeyPair(
		env.GetStr(EnvKeyGmCertPath, ""),
		env.GetStr(EnvKeyGmKeyPath, ""))
	if err != nil {
		log.Panicf("GM Api certificate eror : %s", err.Error())
	}

	apiInst, _ = NewGmClient(*url, point, cert)
}

// NewGmSG returns GM Server Gate Client
func NewGmClient(u url.URL, p int, cert tls.Certificate) (gmCl *Client, err error) {
	out := Client{
		url:   u,
		point: p,
		client: &http.Client{
			Timeout: 120 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					Renegotiation:      tls.RenegotiateOnceAsClient,
					Certificates:       []tls.Certificate{cert},
					InsecureSkipVerify: true,
				},
			}},
	}
	return &out, nil
}

func Balance() (*GmBalance, error) {
	if apiInst == nil {
		return nil, errors.New(NoDefaultApiDefined)
	}
	return apiInst.Balance()
}

func Verify(service int, account string, attrs []GmAttribute) (*GmVerify, error) {
	if apiInst == nil {
		return nil, errors.New(NoDefaultApiDefined)
	}
	return apiInst.Verify(service, account, attrs)
}

func Status(id string) (*GmStatus, error) {
	if apiInst == nil {
		return nil, errors.New(NoDefaultApiDefined)
	}
	return apiInst.Status(id)
}

func Payment(req *GmPayment) (*GmStatus, error) {
	if apiInst == nil {
		return nil, errors.New(NoDefaultApiDefined)
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return apiInst.Payment(req)
}

func Advanced(service int, fn string, attrs []GmAttribute) (*GmAdvanced, error) {
	if apiInst == nil {
		return nil, errors.New(NoDefaultApiDefined)
	}
	return apiInst.Advanced(service, fn, attrs)
}

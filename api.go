package gmapi

import (
	"crypto/tls"
	"errors"
	gmfin "github.com/biteffect/go.gm-fin"
	"github.com/rs/zerolog"
	"net/http"
	"net/url"
	"os"
	"time"
)

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
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.WarnLevel)
	out.logger = &logger
	return &out, nil
}

func SetLogger(l *zerolog.Logger) {
	apiInst.SetLogger(l)
}

func GetBalance() (*Balance, error) {
	if apiInst == nil {
		return nil, errors.New(NoDefaultApiDefined)
	}
	return apiInst.GetBalance()
}

func Verify(service int, account string, attrs []Attribute) (*VerifyStatus, error) {
	if apiInst == nil {
		return nil, errors.New(NoDefaultApiDefined)
	}
	return apiInst.Verify(service, account, attrs)
}

func GetStatus(id string) (*Status, error) {
	if apiInst == nil {
		return nil, errors.New(NoDefaultApiDefined)
	}
	return apiInst.GetStatus(id)
}

func Payment(id string, service int, amount gmfin.Amount, account string, opt *PaymentOptions) (*Status, error) {
	if apiInst == nil {
		return nil, errors.New(NoDefaultApiDefined)
	}
	return apiInst.Payment(id, service, amount, account, opt)
}

/*
func Advanced(service int, fn string, attrs []GmAttribute) (*GmAdvanced, error) {
	if apiInst == nil {
		return nil, errors.New(NoDefaultApiDefined)
	}
	return apiInst.Advanced(service, fn, attrs)
}
*/

package gmapi

const (
	GmTimeFormat = "2006-01-02T15:04:05Z0700"

	EnvKeyGmApiUrl   = "GM_API_URL"
	EnvKeyGmApiPoint = "GM_API_POINT"
	EnvKeyGmCertPath = "GM_API_CERTIFICATE_PATH"
	EnvKeyGmKeyPath  = "GM_API_KEY_PATH"

	DefaultGmApiUrl = "https://globalmoney.cash/external/extended-cert"

	NoDefaultApiDefined = "Unable to call : no dafualt api define"

	GmApiSourceGate = "GATE"
)

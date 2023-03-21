package gmapi

const (
	StatusNew      = 0
	StatusSuccess  = 60
	StatusError    = 80
	StatusNotFound = -2

	EnvKeyGmApiUrl   = "GM_API_URL"
	EnvKeyGmApiPoint = "GM_API_POINT"
	EnvKeyGmCertPath = "GM_API_CERTIFICATE_PATH"
	EnvKeyGmKeyPath  = "GM_API_KEY_PATH"

	DefaultGmApiUrl = "https://globalmoney.cash/external/extended-cert"

	NoDefaultApiDefined = "Unable to call : no default api define"

	GmApiSourceGate = "GATE"
)

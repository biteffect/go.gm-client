package gmapi

type GmBalance struct {
	Balance   int `xml:"balance,attr"`
	Overdraft int `xml:"overdraft,attr"`
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
	ID            int           `xml:"id,attr"`
	Sum           int           `xml:"sum,attr"`
	Check         int           `xml:"check,attr"`
	Service       int           `xml:"service,attr"`
	Source        string        `xml:"source,attr"`
	Account       string        `xml:"account,attr"`
	Date          string        `xml:"date,attr"`
	Delayed       int           `xml:"delayed,attr"`
	TerminalVPSID string        `xml:"terminal-vps-id,attr"`
	Attribute     []GmAttribute `xml:"attribute"`
}

type GmStatus struct {
	ID        int           `xml:"id,attr"`
	State     int           `xml:"state,attr"`
	Substate  int           `xml:"substate,attr"`
	Code      int           `xml:"code,attr"`
	Final     int           `xml:"final,attr"`
	Trans     int           `xml:"trans,attr"`
	Attribute []GmAttribute `xml:"attribute"`
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

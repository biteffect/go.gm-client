package gmapi

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

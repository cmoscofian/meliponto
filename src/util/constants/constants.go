package constants

// BaseURI for the ahgora server
const BaseURI string = "https://www.ahgora.com.br"

// DateLayout is the date format used for the ahgora server "yyyy-mm-dd"
const DateLayout string = "2006-01-02"

// DateInputLayout is the date format used to parse input from the user "d-m-yy"
const DateInputLayout string = "2-1-06"

// TimeLayout is the time format used for the ahgora server "HH:mm"
const TimeLayout string = "15:04"

// URI used for every specific call
const (
	SiteLoginURI        string = "/externo/login"
	GetMonthlyReportURI string = "/api-espelho/apuracao/%s"
	SitePunchURI        string = "/api-espelho/justificativas"
)

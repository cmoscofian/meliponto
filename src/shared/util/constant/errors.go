package constant

// Error messages used throughout the app
const (
	ClientError            string = "Client %s has not been initialized"
	EndAfterBeginDateError string = "End date provided must be after start date provided"
	InvalidDateError       string = "Unable to parse date provided (format: \"d-m-yy\")"
	InvalidLoginError      string = "Unable to login, check your RE on config and the plataform!"
	MissingDateError       string = "Date must be provided (format: \"d-m-yy\")"
	MissingDatesError      string = "Start and finish date arguments must be provided (format: \"d-m-yy\")"
	NoPreviousPunchesError string = "There are no punches to report"
	PastDateError          string = "Date provided must be in the past"
	PastDateTimeError      string = "Date and Time provided must be in the past"
	ServiceError           string = "Service %s has not been initialized"
	RestClientError        string = "Rest client has not been initialized"
	RestServiceError       string = "%s (%d) - Message: %s"
)

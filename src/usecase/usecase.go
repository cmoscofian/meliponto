package usecase

import (
	"encoding/json"
	"time"

	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

// The Usecase type implements a default structure for creating
// every punch possibility.
type Usecase struct {
	datetime       time.Time
	onGard         bool
	message        string
	time           string
	allowanceEnd   string
	allowanceStart string
}

// The Usecaser is an interface that defines all
// methods required for implementing every UC.
type Usecaser interface {
	Create() ([]byte, error)
}

// NewUsecase returns a pointer to a new Usecase given a message,
// date and onGard flag.\
func NewUsecase(msg string, date time.Time, onGard bool) *Usecase {
	return &Usecase{
		message:  msg,
		datetime: date,
		onGard:   onGard,
	}
}

// Create returns a []byte of a json body for the usecase u.
func (u *Usecase) Create() ([]byte, error) {
	time := u.time
	if time == "" {
		time = u.datetime.Format(constants.TimeLayout)
	}

	body := new(model.PunchRequest)
	body.Date = u.datetime.Format(constants.DateLayout)
	body.Justification = "outros"
	body.Message = u.message
	body.Type = model.RegularPunch
	body.Punch = &model.PunchField{
		Time: time,
	}

	if u.onGard {
		body.Allowance = &model.AllowanceField{
			Begin:  u.allowanceStart,
			End:    u.allowanceEnd,
			Period: "especifico",
			Reason: "Sobreaviso",
		}
		body.Type = model.AllowancePunch
		body.Punch = nil
	}

	return json.Marshal(body)
}

// SetTime sets the time field from the u entity based
// upon the time string argument provided.
func (u *Usecase) SetTime(time string) {
	u.time = time
}

// SetAllowance sets the fields from the u entity based
// for a valid allowance type punch (abono) based upon
// the startTime and endTime string arguments provided.
func (u *Usecase) SetAllowance(startTime, endTime string) {
	u.onGard = true
	u.allowanceStart = startTime
	u.allowanceEnd = endTime
}

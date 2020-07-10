package usecase

import (
	"encoding/json"
	"time"

	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

type Usecase struct {
	datetime       time.Time
	onGard         bool
	message        string
	time           string
	allowanceEnd   string
	allowanceStart string
}

type Usecaser interface {
	Create() ([]byte, error)
}

func NewUsecase(msg string, date time.Time, onGard bool) *Usecase {
	return &Usecase{
		message:  msg,
		datetime: date,
		onGard:   onGard,
	}
}

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

func (u *Usecase) SetTime(time string) {
	u.time = time
}

func (u *Usecase) SetAllowance(startTime, endTime string) {
	u.onGard = true
	u.allowanceStart = startTime
	u.allowanceEnd = endTime
}

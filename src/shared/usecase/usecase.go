package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

// The Key type is a custom string type
// to attach and encapsulate the get of usecases
type Key string

// Definition of all usecase keys
const (
	GardFirstDay Key = "gard_first_day"
	GardLastDay  Key = "gard_last_day"
	GardHoliday  Key = "gard_holiday"
	GardSaturday Key = "gard_saturday"
	GardSunday   Key = "gard_sunday"
	GardWeekday  Key = "gard_weekday"
	Regular      Key = "regular"
	Custom       Key = "custom"
)

var usecases = map[Key]domain.Usecase{
	GardFirstDay: &gardFirstDay{},
	GardLastDay:  &gardLastDay{},
	GardHoliday:  &gardHoliday{},
	GardSaturday: &gardSaturday{},
	GardSunday:   &gardSunday{},
	GardWeekday:  &gardWeekday{},
	Regular:      &regular{},
	Custom:       &custom{},
}

// Get returns a slice of pointers to PunchRequest entities and
// an error given a valid context and date.
// It is used to retrieve the usecases.
func (u Key) Get(ctx *entities.Context, date time.Time) []*entities.PunchRequest {
	return usecases[u].Get(ctx, date)
}

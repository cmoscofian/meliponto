package chain

import (
	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/chain/node"
)

// New returns a new chain of responsability
// based on a set of nodes and conditions
func New() domain.Node {
	regular := node.NewRegularWeekday()
	regularCustom := node.NewRegularForced()
	gardStart := node.NewGardFirstDay()
	gardEnd := node.NewGardLastDay()
	gardHoliday := node.NewGardHoliday()
	gardSaturday := node.NewGardSaturday()
	gardSunday := node.NewGardSunday()
	gardWeekday := node.NewGardWeekday()

	regular.SetNext(gardWeekday)
	gardWeekday.SetNext(gardSaturday)
	gardSaturday.SetNext(gardSunday)
	gardSunday.SetNext(gardHoliday)
	gardHoliday.SetNext(gardStart)
	gardStart.SetNext(gardEnd)
	gardEnd.SetNext(regularCustom)
	regularCustom.SetNext(nil)

	return regular
}

package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vitreuz/xtmg-ref/srv/models/constant"
)

type Filter struct {
	method constant.FilterMethod
	field  string
	value  string
}

func SelectFilter(field, value string) Filter {
	return Filter{method: constant.FilterSelect, field: field, value: value}
}

func ExcludeFilter(field, value string) Filter {
	return Filter{method: constant.FilterExclude, field: field, value: value}
}

func (f Filter) isSelect() bool {
	return f.method == constant.FilterSelect
}

func (f Filter) isExclude() bool {
	return f.method == constant.FilterExclude
}

type UnknownFilter string

func (e UnknownFilter) Error() string {
	return fmt.Sprintf("unexpected filter value: %s", string(e))
}

func (s Ship) AppendByFilters(ships []Ship, filters ...Filter) ([]Ship, []Warning) {
	if len(filters) == 0 {
		return append(ships, s), nil
	}

	warnings := []Warning{}
	for _, filter := range filters {
		selected, err := s.selectBy(filter)
		if err != nil {
			warnings = append(warnings, Warning{Error: err})
			continue
		}

		if filter.isSelect() && !selected {
			return ships, warnings
		}
		if filter.isExclude() && selected {
			return ships, warnings
		}
	}

	return append(ships, s), warnings
}

func (s Ship) selectBy(filter Filter) (bool, error) {
	switch filterBy := filter.field; filterBy {
	case "name":
		return s.selectByNameOrXWS(filter.value), nil
	case "faction":
		return s.selectByFaction(filter.value), nil
	case "attack", "agility", "hull", "shields":
		return s.selectByStat(filterBy, filter.value)
	case "action":
		return s.selectByAction(filter.value), nil
	case "maneuvers":
		return s.selectByManuevers(filter.value)
	case "size":
		return string(s.Size) == filter.value, nil
	case "firing_arcs":
		return s.selectByFiringArc(filter.value), nil
	case "xws":
		return strings.Contains(s.XWS, filter.value), nil
	}

	return filter.isSelect(), UnknownFilter(filter.field)
}

func (s Ship) selectByNameOrXWS(name string) bool {
	return strings.Contains(s.Name, name) || strings.Contains(s.XWS, name)
}

func (s Ship) selectByFaction(faction string) bool {
	for _, shipFaction := range s.Faction {
		if strings.Contains(shipFaction.Simplify(), faction) {
			return true
		}
		if strings.Contains(string(shipFaction.ToSquadronFaction()), faction) {
			return true
		}
	}

	return false
}

func (s Ship) selectByStat(stat, value string) (bool, error) {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return false, err
	}

	switch stat {
	case "attack":
		return s.Attack >= intVal, nil
	case "agility":
		return s.Agility >= intVal, nil
	case "hull":
		return s.Hull >= intVal, nil
	case "shields":
		return s.Shields >= intVal, nil
	}

	return false, nil
}

func (s Ship) selectByAction(action string) bool {
	for _, shipAction := range s.Actions {
		if shipAction.Simplify() == action {
			return true
		}
	}
	return false
}

func (s Ship) selectByManuevers(value string) (bool, error) {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return false, err
	}

	return s.ManeuversSummary()[0][0] >= intVal, nil
}

func (s Ship) selectByFiringArc(arc string) bool {
	for _, shipArc := range s.FiringArcs {
		if shipArc.Simplify() == arc {
			return true
		}
	}
	return false
}

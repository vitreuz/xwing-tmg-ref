package database_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	. "github.com/vitreuz/xtmg-ref/srv/database"
	"github.com/vitreuz/xtmg-ref/srv/models"
)

func TestReadShips(t *testing.T) {
	if !IS_DB_SET {
		t.Skip("data not set")
	}

	type checkOut func([]models.Ship, error) error
	check := func(fns ...checkOut) []checkOut { return fns }

	type filterFn func() models.Filter
	SelectShipByName := func(name string) filterFn {
		return func() models.Filter {
			return models.SelectFilter("name", name)
		}
	}
	filter := func(fns ...filterFn) []models.Filter {
		filters := []models.Filter{}
		for _, fn := range fns {
			filters = append(filters, fn())
		}

		return filters
	}

	expectFirstShip := func(expectedName string) checkOut {
		return func(ships []models.Ship, err error) error {
			if len(ships) < 1 {
				return fmt.Errorf(
					"expected to start with %s but got no ships",
					expectedName,
				)
			}
			if actualName := ships[0].Name; actualName != expectedName {
				return fmt.Errorf(
					"expected to start with %s, got %s",
					expectedName, actualName,
				)
			}
			return nil
		}
	}
	expectShipCount := func(expectedCount int) checkOut {
		return func(ships []models.Ship, err error) error {
			if count := len(ships); count != expectedCount {
				if count > 0 {
					return fmt.Errorf(
						"expected to have %d ships, got %d",
						expectedCount, count,
					)
				}
			}
			return nil
		}
	}
	hasNoError := func() checkOut {
		return func(ships []models.Ship, err error) error {
			if err != nil {
				return fmt.Errorf(
					"expected no error but got %v",
					err,
				)
			}
			return nil
		}
	}

	tests := [...]struct {
		name    string
		filters []models.Filter
		checks  []checkOut
	}{
		{
			"Simple scenario",
			filter(),
			check(
				expectFirstShip("X-wing"),
				expectShipCount(56),
				hasNoError(),
			),
		}, {
			"Simple select By Name",
			filter(SelectShipByName("tie")),
			check(
				expectFirstShip("TIE Fighter"),
				expectShipCount(14),
				hasNoError(),
			),
		}, {
			"Select using xws",
			filter(SelectShipByName("tief")),
			check(
				expectFirstShip("TIE Fighter"),
				expectShipCount(2),
				hasNoError(),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDB := createTempDB(t)
			defer os.Remove(tmpDB)
			db := createTestDB(t, tmpDB)

			ships, err := db.ReadShips(tt.filters...)
			for _, check := range tt.checks {
				if checkErr := check(ships, err); checkErr != nil {
					t.Error(checkErr)
				}
			}

		})
	}
}

func TestReadShipByXWS(t *testing.T) {
	if !IS_DB_SET {
		t.Skip("data not set")
	}

	type checkOut func(models.Ship, error) error
	check := func(fns ...checkOut) []checkOut { return fns }

	expectError := func(expectedError error) checkOut {
		return func(ship models.Ship, err error) error {
			if err == nil {
				return fmt.Errorf(
					"expected an error of type '%T' but got no error",
					expectedError,
				)
			}
			if reflect.TypeOf(err) != reflect.TypeOf(expectedError) {
				return fmt.Errorf(
					"expected an error of type '%T', got '%T'",
					expectedError, err,
				)
			}
			return nil
		}
	}
	expectShip := func(expectedName string) checkOut {
		return func(ship models.Ship, err error) error {
			if actualName := ship.Name; expectedName != actualName {
				return fmt.Errorf(
					"expected ship %s, got %s",
					expectedName, actualName,
				)
			}
			return nil
		}
	}

	tests := [...]struct {
		name   string
		ship   string
		checks []checkOut
	}{
		{
			"Simple scenario",
			"xwing",
			check(
				expectShip("X-wing"),
			),
		}, {
			"Unable to find resource",
			"some-fake-ship",
			check(
				expectError(UnableToLocateResourceError{}),
			),
		}, {
			"Search twice",
			"xwing",
			check(
				expectShip("X-wing"),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDB := createTempDB(t)
			defer os.Remove(tmpDB)
			db := createTestDB(t, tmpDB)

			ship, err := db.ReadShipByXWS(tt.ship)
			for _, check := range tt.checks {
				if err := check(ship, err); err != nil {
					t.Error(err)
				}
			}
		})
	}

	tmpDB := createTempDB(t)
	defer os.Remove(tmpDB)
	db := createTestDB(t, tmpDB)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ship, err := db.ReadShipByXWS(tt.ship)
			for _, check := range tt.checks {
				if err := check(ship, err); err != nil {
					t.Error(err)
				}
			}
		})
	}
}

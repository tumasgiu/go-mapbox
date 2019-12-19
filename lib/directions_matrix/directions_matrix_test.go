/**
 * go-mapbox Directions Module Tests
 * Wraps the mapbox directions API for server side use
 * See https://www.mapbox.com/api-documentation/#retrieve-a-matrix for API information
 *
 * https://github.com/tumasgiu/go-mapbox
 * Copyright 2017 Ryan Kurte
 */

package directionsmatrix

import (
	"os"
	"testing"
)

import (
	"github.com/tumasgiu/go-mapbox/lib/base"
)

func TestDirectionsMatrix(t *testing.T) {

	b, err := base.NewBase(os.Getenv("MAPBOX_TOKEN"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	Directionsmatrix := NewDirectionsMatrix(b)

	t.Run("Can Lookup Directions Matrix", func(t *testing.T) {
		var opts RequestOpts
		source := []string{"0", "1"}
		dest := []string{"all"}
		opts.SetSources(source)
		opts.SetDestinations(dest)

		locs := []base.Location{{37.752759, -122.467600}, {37.762819, -122.460304}, {37.758095, -122.442253}}

		res, err := Directionsmatrix.GetDirectionsMatrix(locs, RoutingCycling, &opts)
		if err != nil {
			t.Error(err)
		}

		if Codes(res.Code) != CodeOK {
			t.Errorf("Invalid response code: %s", res.Code)
		}

	})

}

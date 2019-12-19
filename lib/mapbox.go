/**
 * go-mapbox Mapbox API Modle
 * Wraps the mapbox APIs for golang server (or application) use
 * See https://www.mapbox.com/api-documentation/for API information
 *
 * https://github.com/tumasgiu/go-mapbox
 * Copyright 2017 Ryan Kurte
 */

package mapbox

import (
	"github.com/tumasgiu/go-mapbox/lib/base"
	"github.com/tumasgiu/go-mapbox/lib/directions"
	"github.com/tumasgiu/go-mapbox/lib/directions_matrix"
	"github.com/tumasgiu/go-mapbox/lib/geocode"
	"github.com/tumasgiu/go-mapbox/lib/map_matching"
	"github.com/tumasgiu/go-mapbox/lib/maps"
	"github.com/tumasgiu/go-mapbox/lib/styles"
)

// Mapbox API Wrapper structure
type Mapbox struct {
	base *base.Base
	// Maps allows fetching of tiles and tilesets
	Maps *maps.Maps
	// Geocode allows forward (by address) and reverse (by lat/lng) geocoding
	Geocode *geocode.Geocode
	// Directions generates directions between arbitrary points
	Directions *directions.Directions
	// Direction Matrix returns all travel times and ways points between multiple points
	DirectionsMatrix *directionsmatrix.DirectionsMatrix
	// MapMatching snaps inaccurate path tracked to a map to produce a clean path
	MapMatching *mapmatching.MapMatching
	Styles      *styles.Styles
}

// NewMapbox Create a new mapbox API instance
func New(token string) (*Mapbox, error) {
	m := &Mapbox{}

	// Create base instance
	base, err := base.NewBase(token)
	if err != nil {
		return nil, err
	}
	m.base = base

	// Bind modules
	m.Maps = maps.NewMaps(m.base)
	m.Geocode = geocode.NewGeocode(m.base)
	m.Directions = directions.NewDirections(m.base)
	m.DirectionsMatrix = directionsmatrix.NewDirectionsMatrix(m.base)
	m.MapMatching = mapmatching.NewMapMaptching(m.base)
	m.Styles = styles.NewStyles(m.base)

	return m, nil
}

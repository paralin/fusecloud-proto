package common

import (
	"fmt"
	"gopkg.in/dancannon/gorethink.v2/types"
)

/* Marshal a GeoLocation into RethinkDB Geo-spatial */

func (p GeoLocation) MarshalRQL() (interface{}, error) {
	return map[string]interface{}{
		"$reql_type$": "GEOMETRY",
		"coordinates": p.Coords(),
		"type":        "Point",
	}, nil
}

func (p *GeoLocation) UnmarshalRQL(data interface{}) error {
	g := &types.Geometry{}
	err := g.UnmarshalRQL(data)
	if err != nil {
		return err
	}
	if g.Type != "Point" {
		return fmt.Errorf("pseudo-type GEOMETRY object has type %s, expected type %s", g.Type, "Point")
	}

	p.Latitude = g.Point.Lat
	p.Longitude = g.Point.Lon

	return nil
}

func (p GeoLocation) Coords() interface{} {
	return []interface{}{p.Longitude, p.Latitude}
}

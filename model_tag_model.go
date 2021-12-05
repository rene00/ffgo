/*
 * Firefly III API v1.5.4
 *
 * This is the documentation of the Firefly III API. You can find accompanying documentation on the website of Firefly III itself (see below). Please report any bugs or issues. You may use the \"Authorize\" button to try the API below. This file was last generated on 2021-09-25T14:21:28+00:00 
 *
 * API version: 1.5.4
 * Contact: james@firefly-iii.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package ffgo
import (
	"time"
)

type TagModel struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The tag
	Tag string `json:"tag"`
	// The date to which the tag is applicable.
	Date time.Time `json:"date,omitempty"`
	Description string `json:"description,omitempty"`
	// Latitude of the tag's location, if applicable. Can be used to draw a map.
	Latitude float64 `json:"latitude,omitempty"`
	// Latitude of the tag's location, if applicable. Can be used to draw a map.
	Longitude float64 `json:"longitude,omitempty"`
	// Zoom level for the map, if drawn. This to set the box right. Unfortunately this is a proprietary value because each map provider has different zoom levels.
	ZoomLevel int32 `json:"zoom_level,omitempty"`
}

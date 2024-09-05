package events

import "errors"

var (
	NoCelebId     = errors.New("celebrity id required")
	NoTitle       = errors.New("title required")
	NoDescription = errors.New("description required")
	NoCity        = errors.New("city required")
	NoCountry     = errors.New("country required")
)

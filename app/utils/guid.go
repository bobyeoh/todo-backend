package utils

import (
	"github.com/rs/xid"
)

// GetGUID godoc
func GetGUID() string {
	guid := xid.New()
	return guid.String()
}

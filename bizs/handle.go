package bizs

import (
	"sample/dbs"
)

//
func HandleHealthCheck() string {
	if e := dbs.HealthCheck(); e != nil {
		return e.Error()
	}

	return "200"
}

// Package prototime returns the protosocial time.
package prototime

import (
	"fmt"
	"net/http"
	"time"
)

// Now returns the current time.
func Now(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now())
}

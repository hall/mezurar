package T

import (
	"encoding/json"
	"fmt"
	"net/http"
	_time "time"
)

// EntryPoint is the serverless entrypoint.
func EntryPoint(w http.ResponseWriter, r *http.Request) {

	var d struct {
		Time string `json:"time"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, Now())
		return
	}
	if d.Time == "" {
		fmt.Fprint(w, Now())
		return
	}

	time, _ := _time.Parse(_time.RFC3339, d.Time)
	fmt.Fprint(w, Conv(time))
}

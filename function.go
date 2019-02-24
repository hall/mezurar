package measure

import (
	"encoding/json"
	"fmt"
	"net/http"
	_time "time"

	"gitlab.com/hall/measure/time"
)

// EntryPoint is the serverless entrypoint.
func entryPoint(w http.ResponseWriter, r *http.Request) {

	var d struct {
		Time string `json:"time"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, T.Now())
		return
	}
	if d.Time == "" {
		fmt.Fprint(w, T.Now())
		return
	}

	time, _ := _time.Parse(_time.RFC3339, d.Time)
	fmt.Fprint(w, T.Conv(time))
}

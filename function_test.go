package T

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
	_time "time"
)

func TestEntryPoint(t *testing.T) {

	times := []struct {
		body string // input value
		want string // expected result
	}{
		// beginning of a day, contemporary year "zero"
		{`{"time": "0000-01-01T00:00:00Z"}`, "6930.0|0.000"},
		// end of a day and year
		{`{"time": "2000-12-31T23:59:59Z"}`, "7b18.b|26.bbb"},
		// halfway through a day and year
		{`{"time": "0396-07-01T12:30:30Z"}`, "7000.6|0.666"},
	}

	for _, time := range times {
		out := request(time.body, t)

		if got := string(out); got != time.want {
			t.Errorf("EntryPoint(%q) = %q (want %q)", time.body, got, time.want)
		}
	}

	// Nows is a list of inputs that should return the current time.
	nows := []string{`{"time": ""}`, "{}", ""}

	want := request("", t)
	for _, now := range nows {
		out := request(now, t)
		if string(out) != string(want) {
			t.Errorf("EntryPoint(%q) = %q (want %q)", now, out, want)
		}
	}

}

func TestNow(t *testing.T) {

	_now := Conv(_time.Now().UTC())
	now := Now()
	if now != _now {
		t.Errorf("Now(): expected %s, actual %s", _now, now)
	}
}

func request(body string, t *testing.T) []byte {
	req := httptest.NewRequest("GET", "/", strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	EntryPoint(rr, req)

	out, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}

	return out
}

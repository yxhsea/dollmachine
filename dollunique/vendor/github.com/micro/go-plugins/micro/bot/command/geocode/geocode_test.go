package geocode

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestGeocode(t *testing.T) {
	// skip on travis
	if tr := os.Getenv("TRAVIS"); len(tr) > 0 {
		return
	}

	testData := []struct {
		address  string
		response [2]string
	}{
		{"somerset house", [2]string{"51.51", "-0.12"}},
	}

	command := Geocode()

	for _, d := range testData {
		rsp, err := command.Exec("geocode", d.address)
		if err != nil {
			t.Fatal(err)
		}

		parts := strings.Split(string(rsp), ",")
		if len(parts) != 2 {
			t.Fatalf("Expected 2 parts, got %v", parts)
		}

		flat, _ := strconv.ParseFloat(parts[0], 64)
		flng, _ := strconv.ParseFloat(parts[1], 64)
		lat := fmt.Sprintf("%.2f", flat)
		lng := fmt.Sprintf("%.2f", flng)

		if (lat != d.response[0]) || (lng != d.response[1]) {
			t.Fatalf("Expected %v got %s,%s", d.response, lat, lng)

		}
	}
}

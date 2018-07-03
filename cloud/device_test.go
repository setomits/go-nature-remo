package natureremocloud

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"
)

func TestDevices(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/devices.json")
	if err != nil {
		t.Error(err)
	}

	var devices []*Device
	if err := json.Unmarshal(b, &devices); err != nil {
		t.Fatal(err)
	}

	if len(devices) != 2 {
		t.Errorf("There should be 2 devices, but %d", len(devices))
	}

	now := time.Now()
	for _, d := range devices {
		if !d.CreatedAt.Before(now) {
			t.Errorf("CreatedAt should be before %v, but %v", now, d.CreatedAt)
		}
		if !d.UpdatedAt.Before(now) {
			t.Errorf("UpdatedAt should be before %v, but %v", now, d.UpdatedAt)
		}

		if !d.NewestEvents.Temperature.CreatedAt.Before(now) {
			t.Errorf("Temperature.CreatedAt should be before %v, but %v",
				now, d.NewestEvents.Temperature.CreatedAt)
		}
	}
}

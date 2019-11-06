package fitdump

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

const sampleData = `
[{
  "logId" : 1501830069000,
  "weight" : 120.9,
  "bmi" : 18.93,
  "fat" : 9.277000427246094,
  "date" : "08/04/17",
  "time" : "07:01:09",
  "source" : "Aria"
},{
  "logId" : 1502616322000,
  "weight" : 122.7,
  "bmi" : 19.21,
  "fat" : 9.480999946594238,
  "date" : "08/13/17",
  "time" : "09:25:22",
  "source" : "Aria"
}]
`

var sampleExpect = WeightLog{
	WeightEntry{
		ID:         1501830069000,
		Weight:     120.9,
		BMI:        18.93,
		Fat:        9.277000427246094,
		RecordedAt: time.Date(2017, 8, 4, 7, 1, 9, 0, time.Local),
		Source:     "Aria",
	},
	WeightEntry{
		ID:         1502616322000,
		Weight:     122.7,
		BMI:        19.21,
		Fat:        9.480999946594238,
		RecordedAt: time.Date(2017, 8, 13, 9, 25, 22, 0, time.Local),
		Source:     "Aria",
	},
}

func TestUnmarshalJSON(t *testing.T) {
	var wl WeightLog
	err := json.Unmarshal([]byte(sampleData), &wl)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(wl, sampleExpect) {
		t.Errorf("want %v got %v", sampleExpect, wl)
	}
}

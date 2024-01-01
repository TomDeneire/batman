package batmantest

import (
	"testing"

	util "tomdeneire.be/batman/lib/util"
)

func TestBatteryInfo(t *testing.T) {
	state, percentage, err := util.BatteryInfo()

	if err != nil {
		t.Errorf("Could not get battery status: %v", err)
	}

	if (state != "Full") && (state != "Charging") && (state != "Discharging") {
		t.Errorf("Invalid battery state: %v", state)
	}

	if (percentage < 0) || (percentage > 100) {
		t.Errorf("Invalid battery percentage: %v", percentage)
	}
}

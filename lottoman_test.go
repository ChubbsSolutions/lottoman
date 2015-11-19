package main

import "testing"

var proposed = 7

func TestUnitMain(t *testing.T) {

}
func TestUnitIsNumGood(t *testing.T) {
	var s []int
	for i := range s {
		s[i] = i
		i++
	}
	if IsNumGood(s, proposed) != true {
		t.Error("Error with the proposed number.")
	}
}

func TestUnitHoosierLottoGet(t *testing.T) {
	res := hoosierLottoGet()
	if !(res[1] >= 0) {
		t.Error("Error getting Hoosier lotto numbers.")
	}
}

func TestUnitPowerballGet(t *testing.T) {
	res, pb := powerballGet()
	if !(res[1] > 0) || !(pb >= 0) {
		t.Error("Error getting Powerball numbers.")
	}
}

func TestUnitGenerateRandomNumber(t *testing.T) {
	proposed := generateRandomNumber(1, 25)
	if (proposed < 0) || (proposed > 25) {
		t.Error("Error generating number.")
	}
}

package word

import "testing"

func TestCountTransformationStep(t *testing.T) {
	basicTesting := func(t *testing.T, from, to string, expected int) {
		if output := CountTransformationStep(from, to); output != expected {
			t.Errorf(`
Test Failed with input : 
From 	: %s
To 		: %s
Expecting %v but got %v
		`, from, to, expected, output)
		}
	}
	basicTesting(t, "BAYAM", "AYAM", 1)
	basicTesting(t, "IKAN", "SAPI", 4)
	basicTesting(t, "HITAM", "PUTIH", 4)
}

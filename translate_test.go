package translate

import "testing"

func TestSplit(t *testing.T) {
	SplitMagentoCSV("/home/mihai/Documents/littlepatrick/transate/step2/1")
}

func TestCombine(t *testing.T) {
	Combine("/home/mihai/Desktop/step2/left", "/home/mihai/Desktop/step2/mid", "/home/mihai/Desktop/step2/right")
}

func TestNoDuplicates(t *testing.T) {
	NoDuplicates("/home/mihai/Documents/littlepatrick/transate/ro_RO.csv")
}

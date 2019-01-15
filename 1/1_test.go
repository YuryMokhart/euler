package main 

import(
"testing"
)

func Test1(t *testing.T) {
	var test = []struct{
		x int
		y int
		output int
	}{
		{2,10,20},
		{4,10,12},
		{5,10,5},
		{1,10,45},
	}

	for _, table := range test {
		testTotal := sum(table.x, table.y)
		if testTotal != table.output {
			t.Errorf("Incorrect answer, the sum of all the multiples of %d below %d EXPECTED to be %d, GOT %d", table.x, table.y, table.output, testTotal)
		}
	}
 }
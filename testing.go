package translate

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func LineEndings(left, mid string, delimiter []byte) {
	leftFile, _ := ioutil.ReadFile("left")
	midFile, _ := ioutil.ReadFile("mid")

	expr := regexp.MustCompile(`\n`)
	indexesL := expr.FindAllIndex(leftFile, -1)
	indexesM := expr.FindAllIndex(midFile, -1)

	if len(indexesL) != len(indexesM) {
	}

	for i := range indexesL {
		charL := indexesL[i][0] - 1
		charM := indexesM[i][0] - 1
		if leftFile[charL] != midFile[charM] {
			t.Errorf("pe randul %d %s != %s", i, string(leftFile[charL]), string(midFile[charM]))
		}
	}
}

func TestQuotes(t *testing.T) {
	leftFile, _ := ioutil.ReadFile("step2/left")
	midFile, _ := ioutil.ReadFile("step2/mid")

	rowL := make([]int, 0)
	rowM := make([]int, 0)
	row := make([]int, 0)
	noL := 0
	noM := 0

	for _, v := range leftFile {
		if v == '\n' {
			rowL = append(rowL, noL)
			noL = 0
		}
		if v == '\'' {
			noL++
		}
	}
	for _, v := range midFile {
		if v == '\n' {
			rowM = append(rowM, noM)
			noM = 0
		}
		if v == '\'' {
			noM++
		}
	}
	for i := range rowL {
		if rowL[i]-rowM[i] != 0 {
			row = append(row, i+1)
		}
	}
	fmt.Println(row)
	fmt.Println(len(row))
}


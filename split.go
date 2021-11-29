package translate

import (
	"io/ioutil"
	"os"
	"regexp"
)

// This function divides the content of the file in three different files.
// It thakes as argument the file to split which is the dictionary exported by magento.

func SplitMagentoCSV(fileName string) {
	//get the file content
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("file not found: " + fileName)
	}

	expr := regexp.MustCompile(`,(module|theme|lib),.+\n`)
	indexes := expr.FindAllIndex(fileContent, -1)

	//content of the 3 files
	leftContent := make([]byte, 0)
	midContent := make([]byte, 0)
	rightContent := make([]byte, 0)
	currentRow := 0
	midIndex := int((indexes[currentRow][0] / 2) + 1)

	for i, v := range fileContent {

		switch {
		case i < midIndex:
			leftContent = append(leftContent, v)
		case i == midIndex:
			leftContent = append(leftContent, byte('\n'))
			midContent = append(midContent, v)
		default:
			switch {
			case i < indexes[currentRow][0]:
				midContent = append(midContent, v)
			case i == indexes[currentRow][0]:
				midContent = append(midContent, v)
				midContent = append(midContent, byte('\n'))
			default:
				switch {
				case i < indexes[currentRow][1]:
					rightContent = append(rightContent, v)
				default:
					currentRow++
					midIndex = int(((indexes[currentRow][0] - indexes[currentRow-1][1]) / 2) + indexes[currentRow-1][1] + 1)
					leftContent = append(leftContent, v)
				}
			}

		}
	}
	mid, _ := os.OpenFile("mid", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	mid.WriteString(string(midContent))
	left, _ := os.OpenFile("left", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	left.WriteString(string(leftContent))
	right, _ := os.OpenFile("right", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	right.WriteString(string(rightContent))
}

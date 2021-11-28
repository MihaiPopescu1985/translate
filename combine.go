package translate

import (
	"io/ioutil"
	"os"
)

var il = 0
var im = 0
var ir = 0

func Combine(left, mid, right string) {

	leftFile, _ := ioutil.ReadFile(left)
	midFile, _ := ioutil.ReadFile(mid)
	rightFile, _ := ioutil.ReadFile(right)

	content := make([]byte, 0)

	var state int8 = 0
	// indexes

	for state < 3 {
		switch state {
		case 0:
			if leftFile[il] == '\n' {
				// will crush if the first character is '\n'
				if leftFile[il-1] == ',' {
					state = 1
					il++
					break
				}
			}
			content = append(content, leftFile[il])
			il++
		case 1:
			if midFile[im] == '\n' {
				// will crush if the first character is '\n'
				if midFile[im-1] == ',' {
					state = 2
					im++
					break
				}
			}
			content = append(content, midFile[im])
			im++
		case 2:
			if rightFile[ir] == '\n' {
				state = 0
			}
			content = append(content, rightFile[ir])
			ir++
			if ir >= len(rightFile) {
				state = 3
			}
		}
	}

	final, _ := os.OpenFile("step2/final", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	final.Write(content)
	final.Sync()
	final.Close()
}

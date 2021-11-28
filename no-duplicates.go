package magento

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

/*
* Removes duplicates entries from files. Usefull for magento translations.
* Takes 3 arguments: the left file, the mid file and the right file and combines them into result file.
 */
func RemoveDuplicates(left, mid, right string) {
	leftContent, _ := ioutil.ReadFile(left)
	midContent, _ := ioutil.ReadFile(mid)
	rightContent, _ := ioutil.ReadFile(right)

	leftRows := toRows(leftContent, []byte(",\n"))
	midRows := toRows(midContent, []byte(",\n"))
	rightRows := toRows(rightContent, []byte("\n"))

	set := make(map[string][]string)
	for i := range leftRows {
		set[leftRows[i]] = []string{midRows[i], rightRows[i]}
	}

	final, _ := os.OpenFile("result", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	for k, v := range set {
		k = strings.TrimSuffix(k, "\n")
		final.WriteString(k)
		v[0] = strings.TrimSuffix(v[0], "\n")
		final.WriteString(v[0])
		final.WriteString(v[1])
	}
	final.Sync()
	final.Close()
}

func toRows(content []byte, delimiter []byte) []string {
	row := make([]byte, 0)
	rows := make([]string, 0)

	expr := regexp.MustCompile(string(delimiter))
	indexes := expr.FindAllIndex(content, -1)
	j := 0

	for i, v := range content {
		row = append(row, v)
		if i == indexes[j][1] || v == content[len(content)-1] {
			rows = append(rows, string(row))
			row = make([]byte, 0)
			j++
		}
	}
	return rows
}

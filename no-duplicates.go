package translate

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func NoDuplicates(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic("file not found: " + file)
	}
	expr := regexp.MustCompile(`(module|theme|lib),.+\n`)
	indexes := expr.FindAllIndex(content, -1)

	ki := 0
	vi := 0
	set := make(map[string][]byte, 0)

	for i := range indexes {
		key := make([]byte, 0)
		val := make([]byte, 0)
		for ; ki < indexes[i][0]; ki++ {
			key = append(key, content[ki])
		}
		for vi = indexes[i][0]; vi < indexes[i][1]; vi++ {
			val = append(val, content[vi])
			ki++
		}
		set[string(key)] = val
	}
	final, _ := os.OpenFile("finalnodups3", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	for k, v := range set {
		final.WriteString(k)
		final.Write(v)
	}
	final.Sync()
	final.Close()
}

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

package gl

import (
	"bufio"
	"fmt"
	"strings"
)

type CompileError struct {
	Source string
	Log    string
	Type   uint32
}

func (c CompileError) Error() string {
	scanner := bufio.NewScanner(strings.NewReader(c.Source))
	var source []byte
	var i int

	for scanner.Scan() {
		newLine := fmt.Sprintf("%d: %s\n", i, scanner.Text())
		source = append(source, []byte(newLine)...)
		i++
	}
	return fmt.Sprintf("building shader type: %d\nsource:\n%s\nbuild log: %s\n", c.Type, source, c.Log)
}

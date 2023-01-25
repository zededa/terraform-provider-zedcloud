package resources

import (
	"fmt"
	"strings"
)

func formatTest(name, obj string) string {
	fmt.Println(obj)
	countLeadingSpaces := func(line string) int {
		return len(line) - len(strings.TrimSpace(line))
	}
	res := ""
	lines := strings.Split(obj, "\n")
	prevIndent := 0
	newName := name
	arrayIndex := 0
	isArray := false
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		currentIndent := countLeadingSpaces(line)
		if currentIndent < prevIndent {
			if !strings.HasPrefix(strings.TrimSpace(parts[0]), "-") {
				newName = name
				arrayIndex = 0
				isArray = false
			} else {
				parts[0] = strings.TrimLeft(parts[0], "- ")
				arrayIndex++
			}
		}
		prevIndent = currentIndent
		if len(parts) != 2 {
			newName += "." + strings.Title(strings.TrimRight(strings.TrimSpace(parts[0]), ":"))
			isArray = true
			continue
		}
		newLine := "if "
		newObj := newName
		if isArray {
			newObj += fmt.Sprintf("[%d]", arrayIndex)
		}
		newObj += "."
		newObj += strings.Title(strings.TrimSpace(parts[0]))
		newLine += newObj
		newLine += " != "
		newLine += "\""
		newValue := strings.TrimSpace(parts[1])
		newLine += newValue
		newLine += "\" {"
		newLine += "\n    return fmt.Errorf(\"expect "
		newLine += newObj
		newLine += " == "
		newLine += newValue
		newLine += " but got %+v\", "
		newLine += newObj
		newLine += ")\n}"

		newLine += "\n"

		res += newLine
	}
	return res

}

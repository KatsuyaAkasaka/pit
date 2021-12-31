package replace

import (
	"fmt"
	"strings"
)

func startProtoNumber(str string) int {
	if strings.HasPrefix(strings.TrimSpace(str), "enum") {
		return 0
	}
	return 1
}

func trimComment(line string) string {
	tabs := ""
	for i := range line {
		if fmt.Sprintf("%c", line[i]) == "\t" {
			tabs += "\t"
		} else if fmt.Sprintf("%c", line[i]) == " " {
			tabs += " "
		} else {
			break
		}
	}
	return tabs + strings.TrimSpace(strings.Split(line, "//")[0])
}

func startNestingSign(str string) bool {
	return strings.HasSuffix(str, "{")
}
func endNestingSign(str string) bool {
	return strings.HasSuffix(str, "}")
}

func getBlock(lines []string) ([]string, error) {
	score := 0
	dst := []string{}
	for i, line := range lines {
		line = trimComment(lines[i])
		dst = append(dst, lines[i])
		if startNestingSign(line) {
			score++
		}
		if endNestingSign(line) {
			score--
		}
		if score == 0 {
			return dst, nil
		}
	}
	return nil, fmt.Errorf("failed to find block. Did you select collect {} block?")
}

func addProtoNumber(line *string, num int) {
	if strings.HasPrefix(*line, "reserved") ||
		strings.HasPrefix(*line, "//") ||
		strings.TrimSpace(*line) == "" {
		return
	}
	trimCommentLine := trimComment(*line)
	comment := ""
	if strs := strings.Split(*line, "//"); len(strs) >= 2 {
		comment = "//" + strings.Join(strs[1:], "//")
	}

	if strings.HasSuffix(trimCommentLine, "=") {
		*line = fmt.Sprintf("%s %d;", trimCommentLine, num)
	} else {
		*line = fmt.Sprintf("%s = %d;", trimCommentLine, num)
	}
	*line = fmt.Sprintf("%s %s", *line, comment)
}

func exec(block []string, checkBlockIndex, protoNumIndex, totalLen int) ([]string, error) {
	trimCommentLine := trimComment(block[checkBlockIndex])
	if startNestingSign(trimCommentLine) {
		if len(block) == totalLen && checkBlockIndex == 0 {
			return exec(block, 1, startProtoNumber(trimCommentLine), totalLen)
		}
		smallBlock, err := getBlock(block[checkBlockIndex:])
		if err != nil {
			return nil, err
		}
		res, err := exec(smallBlock, 1, startProtoNumber(trimCommentLine), totalLen)
		if err != nil {
			return nil, err
		}
		block = append(
			block[:checkBlockIndex],
			append(
				res,
				block[checkBlockIndex+len(smallBlock):]...,
			)...,
		)
		checkBlockIndex += len(res) - 1
		protoNumIndex--
	} else if endNestingSign(trimCommentLine) {
		return block, nil
	} else {
		addProtoNumber(&block[checkBlockIndex], protoNumIndex)
	}
	return exec(block, checkBlockIndex+1, protoNumIndex+1, totalLen)
}

func (v *Variable) protoNum() (string, error) {
	lines := strings.Split(v.Option.InputText, "\n")
	if len(lines) < 2 {
		return "", fmt.Errorf("needs at least 2 lines")
	}
	res, err := exec(lines, 0, 0, len(lines))
	if err != nil {
		return "", err
	}
	return strings.Join(res, "\n"), nil
}

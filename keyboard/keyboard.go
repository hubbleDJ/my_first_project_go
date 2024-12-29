package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFloat64(terminalText string) (float64, error) {
	fmt.Print(terminalText)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0.0, err
	}

	input = strings.TrimSpace(input)
	returnFloat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0.0, err
	}

	return returnFloat, nil
}

package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNumber(msg string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	in, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	val, err := strconv.Atoi(strings.TrimSpace(in))
	if err != nil {
		return 0, err
	}
	return val, nil
}

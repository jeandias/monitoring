package cmd

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func GetOption(stdin io.Reader) (int, error) {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit program")

	reader := bufio.NewReader(stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	option, err := strconv.Atoi(input)

	return option, err
}

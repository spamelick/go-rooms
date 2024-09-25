package ihelp

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// Ввод значений в терминале.
func Input(promt string, r *bufio.Reader) string {
	fmt.Print(promt)
	answer, _ := r.ReadString('\n')

	return strings.TrimSpace(answer)
}

func Radio(label string, opts []string) string {
	var res string
	prompt := &survey.Select{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

// func multiselect(label string, opts []string) []string {
// 	var res []string
// 	prompt := &survey.MultiSelect{
// 		Message: label,
// 		Options: opts,
// 	}
// 	survey.AskOne(prompt, &res)

// 	return res
// }

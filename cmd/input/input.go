package input

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

// String 输入字符串
func String(label, value string) string {
	prompt := promptui.Prompt{
		Label:     label,
		Default:   value,
		AllowEdit: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return ""
	}

	return result

}

// Confirm 确认输入
func Confirm(label string) bool {

	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		return false
	}

	if result == "Y" || result == "y" {
		return true
	}

	return false

}

// 选择输入
func SelectWithAdd(label string, items *[]string) int {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:     label,
			Items:     *items,
			AddLabel:  "新增",
			IsVimMode: true,
			HideHelp:  true,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			*items = append(*items, result)
		} else {
			break
		}
		fmt.Println(err)
	}

	return index
}

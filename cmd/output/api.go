package output

import (
	"apimake/pkg/entity"
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
)

func RequiredContent(isRequired int) string {
	if isRequired == 0 {
		return "否"
	}
	return "是"
}

// Tips 提示
func Tips(message string) {
	_, _ = color.New(color.FgCyan).Add(color.Bold).
		Println(message)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v)
}
func ShowTitle(title string) {
	fmt.Println()
	_, _ = color.New(color.FgCyan).Add(color.Bold).
		Println(fmt.Sprintf(">%s", title))
}

func ShowHttpUrl(api entity.Api) {
	ShowTitle("请求地址")
	color.Red(api.HttpUri)
}

func ShowHttpMethod(api entity.Api) {
	ShowTitle("请求方式")
	color.Yellow(api.HttpMethod)
}

func ShowDescription(api entity.Api) {
	ShowTitle("备注说明")
	color.White(api.Description)
}

func ShowHeaderParams(api entity.Api) {
	ShowTitle("Headers:")
	header := []string{"标签", "是否必填", "内容"}
	data := [][]string{
	}

	for _, item := range api.HeaderParams.Items {
		data = append(data, []string{
			item.Tag,
			RequiredContent(item.IsRequired),
			item.Content,
		})
	}

	outputTable(header, data)
}

func ShowRequestParams(api entity.Api) {
	ShowTitle("请求参数:")
	header := []string{"参数名称", "是否必填", "类型", "说明", "内容"}
	data := [][]string{
	}

	for _, item := range api.RequestParams.Items {
		data = append(data, []string{
			item.Name,
			RequiredContent(item.IsRequired),
			item.Type,
			item.Description,
			item.ExampleValue,
		})
	}

	outputTable(header, data)
}

func ShowResponseParams(api entity.Api) {
	ShowTitle("响应参数:")
	header := []string{"参数名称", "是否必含", "类型", "说明", "内容"}
	data := [][]string{
	}

	for _, item := range api.ResponseParams.Items {
		data = append(data, []string{
			item.Name,
			RequiredContent(item.IsRequired),
			item.Type,
			item.Description,
			item.ExampleValue,
		})
	}

	outputTable(header, data)
}

func outputTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, v := range data {
		table.Append(v)
	}
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.Render() // Send output
}

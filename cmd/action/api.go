package action

import (
	"apimake/cmd/flag"
	"apimake/cmd/input"
	"apimake/cmd/output"
	"apimake/pkg/dao"
	"apimake/pkg/entity"
	"apimake/pkg/enum"
	"apimake/pkg/errors"
	"fmt"
	"github.com/marcusolsson/tui-go"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
)

// UpdateApi 更新
func UpdateApi(ctx *cli.Context) error {
	if flag.ApiId < 0 {
		return errors.InvalidParam()
	}

	item := dao.Api().Get(flag.ApiId)
	if item == nil {
		return errors.NotFound()
	}

	switch flag.UpdateType {
	case enum.UpdateTypeApi:
		return updateApi(flag.ApiId)
	case enum.UpdateTypeHeader:
		return updateHeader(item)
	case enum.UpdateTypeRequest:
		return updateRequest(item)
	case enum.UpdateTypeResponse:
		return updateResponse(item)
	}

	return fmt.Errorf("invalid type")
}

func updateApi(id int) error {
	item := dao.Api().Get(id)
	if item == nil {
		return fmt.Errorf("data not found")
	}
	name := tui.NewEntry()
	name.SetText(item.Name)
	description := tui.NewEntry()
	description.SetText(item.Description)
	httpUrl := tui.NewEntry()
	httpUrl.SetText(item.HttpUri)
	httpMethod := tui.NewEntry()
	httpMethod.SetText(item.HttpMethod)

	form := tui.NewGrid(0, 0)
	form.AppendRow(tui.NewLabel("接口名称:"), name)
	form.AppendRow(tui.NewLabel("请求地址:"), httpUrl)
	form.AppendRow(tui.NewLabel("请求方法:"), httpMethod)
	form.AppendRow(tui.NewLabel("备注说明:"), description)

	status := tui.NewStatusBar("Ready.")

	save := tui.NewButton("[保存]")

	cancel := tui.NewButton("[取消]")

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, save),
		tui.NewPadder(1, 0, cancel),
	)

	window := tui.NewVBox(
		tui.NewPadder(12, 2, tui.NewLabel("更新接口")),
		tui.NewPadder(0, 1, form),
		tui.NewPadder(0, 4, buttons),
	)

	wrapper := tui.NewVBox(
		window,
		tui.NewSpacer(),
	)
	content := tui.NewHBox(wrapper)

	root := tui.NewVBox(
		content,
		status,
	)

	tui.DefaultFocusChain.Set(name, httpUrl, httpMethod, description, save, cancel)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	save.OnActivated(func(b *tui.Button) {

		item.Name = name.Text()
		item.Description = description.Text()
		item.HttpMethod = httpMethod.Text()
		item.HttpUri = httpUrl.Text()

		if err := dao.Api().Update(id, item); err != nil {
			log.Fatalf("更新接口失败:%v", err)
		}
		status.SetText("已保存.")

		ui.Quit()

	})
	cancel.OnActivated(func(b *tui.Button) {
		ui.Quit()
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func updateHeader(item *entity.Api) error {
	var items []string
	for _, item := range item.HeaderParams.Items {
		items = append(items, item.Tag)
	}

	index := input.SelectWithAdd("选择参数", &items)

	var header entity.ApiHeaderParam

	if len(item.HeaderParams.Items) < index+1 {
		// 新增
		item.HeaderParams.Items = append(item.HeaderParams.Items, header)
		header.Tag = items[index]
	} else {
		header = item.HeaderParams.Items[index]
		header.Tag = input.String("参数名称", header.Tag)
	}

	if input.Confirm("是否必填") {
		header.IsRequired = enum.BoolTrue
	} else {
		header.IsRequired = enum.BoolFalse
	}
	header.Content = input.String("内容", header.Content)

	item.HeaderParams.Items[index] = header

	if !input.Confirm("确认修改") {
		return fmt.Errorf("已取消")
	}
	if err := dao.Api().Update(flag.ApiId, item); err != nil {
		log.Fatalf("更新接口失败:%v", err)
	}

	return nil
}

func updateRequest(item *entity.Api) error {
	var items []string
	for _, item := range item.RequestParams.Items {
		items = append(items, item.Name)
	}

	index := input.SelectWithAdd("选择参数", &items)

	var request entity.ApiRequestParam

	if len(item.RequestParams.Items) < index+1 {
		// 新增
		item.RequestParams.Items = append(item.RequestParams.Items, request)
		request.Name = items[index]
	} else {
		request = item.RequestParams.Items[index]
		request.Name = input.String("参数名称", request.Name)
	}

	request.Type = input.String("参数类型", request.Type)
	if input.Confirm("是否必填") {
		request.IsRequired = enum.BoolTrue
	} else {
		request.IsRequired = enum.BoolFalse
	}
	request.Description = input.String("参数说明", request.Description)
	request.ExampleValue = input.String("参数示例", request.ExampleValue)

	item.RequestParams.Items[index] = request

	if !input.Confirm("确认修改") {
		return fmt.Errorf("已取消")
	}
	if err := dao.Api().Update(flag.ApiId, item); err != nil {
		log.Fatalf("更新接口失败:%v", err)
	}

	return nil
}

func updateResponse(item *entity.Api) error {
	var items []string
	for _, item := range item.ResponseParams.Items {
		items = append(items, item.Name)
	}

	index := input.SelectWithAdd("选择参数", &items)

	var response entity.ApiResponseParam

	if len(item.ResponseParams.Items) < index+1 {
		// 新增
		item.ResponseParams.Items = append(item.ResponseParams.Items, response)
		response.Name = items[index]
	} else {
		response = item.ResponseParams.Items[index]
		response.Name = input.String("参数名称", response.Name)

	}

	response.Type = input.String("参数类型", response.Type)
	if input.Confirm("是否必含") {
		response.IsRequired = enum.BoolTrue
	} else {
		response.IsRequired = enum.BoolFalse
	}
	response.Description = input.String("参数说明", response.Description)
	response.ExampleValue = input.String("参数示例", response.ExampleValue)

	item.ResponseParams.Items[index] = response

	if !input.Confirm("确认修改") {
		return fmt.Errorf("已取消")
	}
	if err := dao.Api().Update(flag.ApiId, item); err != nil {
		log.Fatalf("更新接口失败:%v", err)
	}

	return nil
}

// CreateApi 创建接口
func CreateApi(ctx *cli.Context) error {
	name := tui.NewEntry()
	description := tui.NewEntry()
	httpUrl := tui.NewEntry()
	httpMethod := tui.NewEntry()

	form := tui.NewGrid(0, 0)
	form.AppendRow(tui.NewLabel("接口名称:"), name)
	form.AppendRow(tui.NewLabel("请求地址:"), httpUrl)
	form.AppendRow(tui.NewLabel("请求方法:"), httpMethod)
	form.AppendRow(tui.NewLabel("备注说明:"), description)

	save := tui.NewButton("[保存]")
	cancel := tui.NewButton("[取消]")

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, save),
		tui.NewPadder(1, 0, cancel),
	)

	window := tui.NewVBox(
		tui.NewPadder(12, 2, tui.NewLabel("创建接口信息")),
		tui.NewPadder(0, 1, form),
		tui.NewPadder(0, 4, buttons),
	)

	wrapper := tui.NewVBox(
		window,
		tui.NewSpacer(),
	)
	content := tui.NewHBox(wrapper)

	root := tui.NewVBox(
		content,
	)

	tui.DefaultFocusChain.Set(name, httpUrl, httpMethod, description, save, cancel)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	save.OnActivated(func(b *tui.Button) {

		api := entity.Api{
			Name:        name.Text(),
			Description: description.Text(),
			HttpMethod:  httpMethod.Text(),
			HttpUri:     httpUrl.Text(),
		}

		if err := dao.Api().Create(api); err != nil {
			output.Fatalf("创建接口失败:%v", err)
		}

		ui.Quit()

	})
	cancel.OnActivated(func(b *tui.Button) {
		ui.Quit()
	})

	if err := ui.Run(); err != nil {
		return err
	}

	return nil
}

// ShowApi 展示接口详情
func ShowApi(ctx *cli.Context) error {
	fmt.Println()

	api := dao.Api().Get(flag.ApiId)
	if api == nil {
		return fmt.Errorf("not found")
	}
	output.ShowTitle(api.Name)
	output.ShowHttpUrl(*api)
	output.ShowHttpMethod(*api)
	output.ShowDescription(*api)
	output.ShowHeaderParams(*api)
	output.ShowRequestParams(*api)
	output.ShowResponseParams(*api)

	return nil
}

// ListApi 接口列表
func ListApi(ctx *cli.Context) error {
	data := [][]string{}

	items := dao.Api().GetItems()
	for index, item := range items {
		data = append(data, []string{
			strconv.Itoa(index + 1),
			item.Name,
			item.HttpUri,
			item.HttpMethod,
			item.Description,
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"编号", "接口名称", "地址", "请求方式", "备注说明"})

	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.BgBlackColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
	return nil
}

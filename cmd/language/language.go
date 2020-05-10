package language

const (
	ZH = "zh"
	EN = "en"
)

var (
	apiUsageItems = map[string]string{
		ZH: "创建接口信息",
		EN: "Create api",
	}
)

func getItem(items map[string]string, lang string) string {
	if item, ok := items[lang]; ok {
		return item
	}

	return items[EN]
}

func GetCreateApiUsage(lang string) string {
	return getItem(apiUsageItems, lang)
}

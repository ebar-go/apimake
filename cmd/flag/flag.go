package flag

import (
	"apimake/cmd/language"
	"apimake/pkg/constant"
	"github.com/urfave/cli/v2"
)

var Lang string
var UpdateType string
var ApiId int
var FilePath string

var (
	VersionFlag = &cli.StringFlag{
		Name:  "version",
		Value: "1.0",
		Usage: "show client version",
	}

	FileFlag = &cli.StringFlag{
		Name:        "file",
		Aliases:     []string{"f"},
		Value:       constant.DefaultStorage,
		DefaultText: constant.DefaultStorage,
		Usage:       "read/write from `FILE`",
		Destination: &FilePath,
	}

	LanguageFlag = &cli.StringFlag{
		Hidden:      true,
		Name:        "language",
		Aliases:     []string{"l"},
		DefaultText: language.EN,
		Destination: &Lang,
		Value:       language.EN,
		Usage:       "Show help with `LANGUAGE` like: zh, en",
	}

	UpdateTypeFlag = &cli.StringFlag{
		Name:        "type",
		Usage:       "choose `TYPE` to update:api,header,request,response",
		Value:       "api",
		DefaultText: "api",
		Destination: &UpdateType,
	}

	ApiIdFlag = &cli.IntFlag{
		Name:        "id",
		Value:       0,
		Usage:       "api primary key,ex: 1",
		Required:    true,
		Destination: &ApiId,
	}
)

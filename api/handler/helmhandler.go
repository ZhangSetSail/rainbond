package handler

import (
	api_model "github.com/goodrain/rainbond/api/model"
	"github.com/goodrain/rainbond/api/util"
)

type HelmHandler interface {
	CommandHelm(command string) (*api_model.HelmCommandRet, *util.APIHandleError)
	AddHelmRepo(helmRepo api_model.CheckHelmApp) error
	CheckHelmApp(checkHelmApp api_model.CheckHelmApp) (string, *util.APIHandleError)
	GetChartInformation(chart api_model.ChartInformation) (*[]api_model.HelmChartInformation, *util.APIHandleError)
}

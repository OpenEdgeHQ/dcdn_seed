// Code generated by hertz generator. DO NOT EDIT.

package dcdn_seed

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	dcdn_seed "seed.manager/biz/handler/dcdn_seed"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_seed_manager := root.Group("/seed_manager", _seed_managerMw()...)
		{
			_device := _seed_manager.Group("/device", _deviceMw()...)
			_device.GET("/task", append(_getdownloadtaskMw(), dcdn_seed.GetDownloadTask)...)
			{
				_report := _device.Group("/report", _reportMw()...)
				_report.POST("/all", append(_reportseedallMw(), dcdn_seed.ReportSeedAll)...)
			}
		}
		{
			_sdk := _seed_manager.Group("/sdk", _sdkMw()...)
			_sdk.GET("/list", append(_queryfidpeerMw(), dcdn_seed.QueryFidPeer)...)
		}
	}
}
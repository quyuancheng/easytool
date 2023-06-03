package controller

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"net/http"
	v1 "tool/pkg/controller/v1"
	"tool/pkg/usecase"
)

// ShellController -
type ShellController struct {
	ShellUcase usecase.ShellUsecase `inject:""`
}

// NewShellController creates a new ShellController.
func NewShellController() *ShellController {
	return &ShellController{}
}

func (m *ShellController) DisablePrefix() bool {
	return true
}

// WebService -
func (m *ShellController) WebService(ws *restful.WebService) {
	tags := []string{"scan"}

	ws.Route(ws.GET("/shell").To(m.GetShell).
		Doc("获取终端信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page", "页码").DataType("integer")).
		Param(ws.QueryParameter("pageSize", "每页的大小. -1 表示不限制").DataType("integer")).
		Param(ws.QueryParameter("startTime", "起始时间").DataType("string")).
		Param(ws.QueryParameter("endTime", "截止时间").DataType("string")).
		Param(ws.QueryParameter("query", "模糊搜索").DataType("string")).
		Returns(200, "ok", v1.Shell{}))
}

func (m *ShellController) GetShell(r *restful.Request, w *restful.Response) {
	w.WriteHeaderAndEntity(http.StatusOK, &v1.Response{
		Message: "请求成功",
		Code:    http.StatusOK,
	})
}

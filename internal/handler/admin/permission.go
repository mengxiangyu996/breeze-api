package admin

import (
	"breeze-api/internal/service"
	"breeze-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// 权限请求
type Permission struct{}

// 创建权限
func (*Permission) Create(ctx *fiber.Ctx) error {

	type request struct {
		Name      string `json:"name"`
		GroupName string `json:"groupName"`
		Path      string `json:"path"`
		Method    string `json:"method"`
		Status    int    `json:"status"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.Path == "" || req.Method == "" {
		return response.Error(ctx, "参数错误")
	}

	permission := (&service.Permission{}).DetailByPathWithMethod(req.Path, req.Method)
	if permission.Id > 0 {
		return response.Error(ctx, "权限已存在")
	}

	if err := (&service.Permission{}).Create(&service.Permission{
		Name:      req.Name,
		GroupName: req.GroupName,
		Path:      req.Path,
		Method:    req.Method,
		Status:    req.Status,
	}); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新权限
func (*Permission) Update(ctx *fiber.Ctx) error {

	type request struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		GroupName string `json:"groupName"`
		Path      string `json:"path"`
		Method    string `json:"method"`
		Status    int    `json:"status"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.Id <= 0 || req.Path == "" || req.Method == "" {
		return response.Error(ctx, "参数错误")
	}

	permission := (&service.Permission{}).DetailByPathWithMethod(req.Path, req.Method)
	if permission.Id > 0 && permission.Id != req.Id {
		return response.Error(ctx, "权限已存在")
	}

	if err := (&service.Permission{}).Update(&service.Permission{
		Id:        req.Id,
		Name:      req.Name,
		GroupName: req.GroupName,
		Path:      req.Path,
		Method:    req.Method,
		Status:    req.Status,
	}); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除权限
func (*Permission) Delete(ctx *fiber.Ctx) error {

	type request struct {
		Id int `json:"id"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.Permission{}).Delete(req.Id); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 权限列表
func (*Permission) Page(ctx *fiber.Ctx) error {

	page := ctx.QueryInt("page", 1)
	size := ctx.QueryInt("size", 10)

	type request struct {
		Name      string `query:"name"`
		GroupName string `query:"groupName"`
		Path      string `query:"path"`
		Method    string `query:"method"`
	}

	var req request

	if err := ctx.QueryParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	list, count := (&service.Permission{}).Page(page, size, req.Name, req.GroupName, req.Path, req.Method)

	return response.Success(ctx, "成功", map[string]interface{}{
		"list":  list,
		"count": count,
	})
}

// 权限详情
func (*Permission) Detail(ctx *fiber.Ctx) error {

	id := ctx.QueryInt("id")
	if id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	permission := (&service.Permission{}).Detail(id)

	return response.Success(ctx, "成功", map[string]interface{}{
		"permission": permission,
	})
}

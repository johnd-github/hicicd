package controllers

import (
	"github.com/hidevopsio/hiboot/pkg/starter/web"
	"net/http"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hicicd/pkg/orch/k8s"
)

// Operations about object
type ConfigMapController struct {
	BaseController
}

func init() {
	web.Add(new(ConfigMapController))
}


func (c *ConfigMapController) Before(ctx *web.Context) {
	c.BaseController.Before(ctx)
}

func (c *ConfigMapController) PostAdd(ctx *web.Context) {
	log.Debug("ConfigMap  add:{}")
	var configMap k8s.ConfigMaps
	err := ctx.RequestBody(&configMap)
	if err != nil {
		ctx.ResponseError(err.Error(), http.StatusUnavailableForLegalReasons)
		return
	}
	config := k8s.NewConfigMaps(configMap.Name, configMap.Namespace, configMap.Data)
	co, e := config.Create()
	log.Debug("create debug",co)
	if e != nil {
		log.Error("create configMap err:", e)
		ctx.ResponseError(err.Error(), http.StatusServiceUnavailable)
		return
	}
	ctx.ResponseBody("success", nil)
	return
}

func (c *ConfigMapController) PostDelete(ctx *web.Context)  {
	log.Debug("ConfigMap  Delete:{}")
	var configMap k8s.ConfigMaps
	err := ctx.RequestBody(&configMap)
	if err != nil {
		ctx.ResponseError(err.Error(), http.StatusRequestedRangeNotSatisfiable)
		return
	}
	err = configMap.Delete()
	if err != nil {
		ctx.ResponseError( err.Error(), http.StatusRequestedRangeNotSatisfiable)
		return
	}
	message := "success"
	ctx.ResponseBody(message, nil)
}

func (c *ConfigMapController) PostGet(ctx *web.Context)  {
	log.Debug("ConfigMap  get:{}")
	var configMap k8s.ConfigMaps
	err := ctx.RequestBody(&configMap)
	if err != nil {
		ctx.ResponseError(err.Error(), http.StatusRequestedRangeNotSatisfiable)
		return
	}
	co, err := configMap.Get()
	if err != nil {
		ctx.ResponseError( err.Error(), http.StatusRequestedRangeNotSatisfiable)
		return
	}
	message := "success"
	ctx.ResponseBody(message, co)
}

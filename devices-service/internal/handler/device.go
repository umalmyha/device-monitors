package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
	"github.com/umalmyha/device-monitors/devices-service/internal/model"
	"github.com/umalmyha/device-monitors/devices-service/internal/service"
)

type Path struct {
	ID uuid.UUID `uri:"id" binding:"required,uuid"`
}

type DeviceService interface {
	FindAll(ctx context.Context, qr model.GetAllDevicesQuery) ([]*model.Device, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.Device, error)
	Create(ctx context.Context, nd service.CreateDevice) (*model.Device, error)
	Update(ctx context.Context, id uuid.UUID, dvc service.UpdateDevice) (*model.Device, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type DeviceHandler struct {
	deviceSrv DeviceService
}

func NewDeviceHandler(deviceSrv DeviceService) *DeviceHandler {
	return &DeviceHandler{deviceSrv: deviceSrv}
}

func (h *DeviceHandler) FindAll(c *gin.Context) {
	var qr model.GetAllDevicesQuery
	if err := c.BindQuery(&qr); err != nil {
		return
	}

	devices, err := h.deviceSrv.FindAll(c.Request.Context(), qr)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, devices)
}

func (h *DeviceHandler) FindByID(c *gin.Context) {
	var path Path
	if err := c.BindUri(&path); err != nil {
		return
	}

	dvc, err := h.deviceSrv.FindByID(c.Request.Context(), path.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dvc)
}

func (h *DeviceHandler) Create(c *gin.Context) {
	var nd service.CreateDevice
	if err := c.BindJSON(&nd); err != nil {
		return
	}

	dvc, err := h.deviceSrv.Create(c.Request.Context(), nd)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, dvc)
}

func (h *DeviceHandler) Update(c *gin.Context) {
	type Update struct {
		Path
		service.UpdateDevice
	}

	var upd Update
	if err := c.Bind(&upd); err != nil {
		return
	}

	dvc, err := h.deviceSrv.Update(c.Request.Context(), upd.ID, upd.UpdateDevice)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dvc)
}

func (h *DeviceHandler) Delete(c *gin.Context) {
	var path Path
	if err := c.BindUri(&path); err != nil {
		return
	}

	if err := h.deviceSrv.Delete(c.Request.Context(), path.ID); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

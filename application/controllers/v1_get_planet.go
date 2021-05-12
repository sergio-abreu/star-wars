package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/application/planets/get_planet"
	"github.com/sirupsen/logrus"
	"net/http"
)

func V1GetPlanet(ctrl *get_planet.GetPlanetController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var query get_planet.GetPlanetQuery
		err := ctx.ShouldBindUri(&query)
		if err != nil {
			abortWithMessage(ctx, http.StatusBadRequest, "invalid input")
			return
		}
		aPlanet, err := ctrl.GetPlanet(query)
		if get_planet.ErrNotFound(err) {
			abortWithMessage(ctx, http.StatusNotFound, errors.Wrap(err, "failed to get planet").Error())
			return
		}
		if err != nil {
			logrus.WithError(err).Error("failed to get planet")
			abortWithMessage(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}
		ctx.JSON(http.StatusOK, aPlanet)
	}
}

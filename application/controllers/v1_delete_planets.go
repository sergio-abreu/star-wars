package controllers

import (
	"github.com/gin-gonic/gin"
	get_planet "github.com/sergio-vaz-abreu/star-wars/modules/world/application/planets/delete_planet"
	"github.com/sirupsen/logrus"
	"net/http"
)

func V1DeletePlanets(ctrl *get_planet.DeletePlanetController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var command get_planet.DeletePlanetCommand
		err := ctx.ShouldBindUri(&command)
		if err != nil {
			abortWithMessage(ctx, http.StatusBadRequest, "invalid input")
			return
		}
		err = ctrl.DeletePlanet(command)
		if err != nil {
			logrus.WithError(err).Error("failed to delete planet")
			abortWithMessage(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}

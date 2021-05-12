package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/application/planets/get_planets"
	"github.com/sirupsen/logrus"
	"net/http"
)

func V1GetPlanets(ctrl *get_planets.GetPlanetsController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allPlanets, err := ctrl.GetPlanets()
		if err != nil {
			logrus.WithError(err).Error("failed to get planets")
			abortWithMessage(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}
		ctx.JSON(http.StatusOK, allPlanets)
	}
}

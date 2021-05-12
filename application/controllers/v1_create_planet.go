package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sergio-vaz-abreu/star-wars/modules/world/application/planets/create_planet"
	"github.com/sirupsen/logrus"
	"net/http"
)

func V1CreatePlanet(ctrl *create_planet.CreatePlanetController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var command create_planet.CreatePlanetCommand
		err := ctx.BindJSON(&command)
		if err != nil {
			abortWithMessage(ctx, http.StatusBadRequest, "invalid input")
			return
		}
		aPlanet, err := ctrl.CreatePlanet(command)
		if create_planet.ErrInvalidInput(err) {
			abortWithMessage(ctx, http.StatusPreconditionFailed, err.Error())
			return
		}
		if err != nil {
			logrus.WithError(err).Error("failed to create planet")
			abortWithMessage(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}
		ctx.JSON(http.StatusCreated, aPlanet)
	}
}

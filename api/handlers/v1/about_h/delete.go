package aboutsh

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	errorh "gitlab.com/burxondvv/blog-cicd-gitlab/api/handlers/v1/error_h"
	"gitlab.com/burxondvv/blog-cicd-gitlab/api/models"
	"gitlab.com/burxondvv/blog-cicd-gitlab/storage/postgres/aboutrepo"
)

// @Summary		Delete about
// @Tags        About
// @Description	Here about can be deleted.
// @Security    BearerAuth
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  models.Response
// @Failure     default {object}  FullResponse
// @Router		/about/{id} [DELETE]
func (h *AboutHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if errorh.HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi()") {
		return
	}

	err = h.postgres.About().Delete(context.Background(), &aboutrepo.DeleteReq{Id: id})
	if errorh.HandleDatabaseLevelWithMessage(c, h.log, err, "h.postgres.About().Delete()") {
		return
	}

	c.JSON(http.StatusOK, models.Response{
		ErrorCode:    errorh.ErrorSuccessCode,
		ErrorMessage: "Successfully deleted",
	})
}

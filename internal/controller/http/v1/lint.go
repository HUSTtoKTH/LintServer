package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/HUSTtoKTH/lintserver/internal/entity"
	"github.com/HUSTtoKTH/lintserver/internal/usecase"
	"github.com/HUSTtoKTH/lintserver/pkg/comerr"
	"github.com/HUSTtoKTH/lintserver/pkg/logger"
	"github.com/gin-gonic/gin"
)

type lintRoutes struct {
	t usecase.Lint
	l logger.Interface
}

func newLintRoutes(handler *gin.RouterGroup, t usecase.Lint, l logger.Interface) {
	r := &lintRoutes{t, l}

	h := handler.Group("/lint")
	{
		h.POST("/upload", r.upload)
		h.GET("/rule/:project_id", r.getRule)
	}
}

type uploadRequest struct {
	ProjectId      int64  `json:"project_id"   binding:"required"  example:"1"`
	OrganizationId int64  `json:"organization_id"   binding:"required"  example:"1"`
	RuleYml        string `json:"rule_yml"   binding:"required"  example:"json string"`
}

// upload TODO
// @Summary     create or update project's lint rule
// @Description create or update project's lint rule
// @ID          upload
// @Tags  	    lint
// @Accept      json
// @Produce     json
// @Param       request body uploadRequest true "Upload Rule"
// @Success     200 {object} response
// @Failure     500 {object} response
// @Router      /lint/upload [post]
func (r *lintRoutes) upload(c *gin.Context) {
	var request uploadRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - upload")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	err := r.t.Upload(c.Request.Context(), entity.Lint{
		ProjectId:      request.ProjectId,
		OrganizationId: request.OrganizationId,
		Rule:           request.RuleYml,
	}, c.GetHeader("Token"))
	if err != nil {
		r.l.Error(err, "http - v1 - upload")
		if errors.Is(err, comerr.ErrPermission) {
			errorResponse(c, http.StatusForbidden, "permission denied")
		} else if errors.Is(err, comerr.ErrUnauthorized) {
			errorResponse(c, http.StatusUnauthorized, "unauthorized")
		} else {
			errorResponse(c, http.StatusInternalServerError, "database problems")
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// getRule TODO
// @Summary     get project's lint rule
// @Description get project's lint rule
// @ID          getRule
// @Tags  	    lint
// @Accept      json
// @Produce     json
// @Param       project_id   path   int  true  "Project ID"
// @Success     200 {object} entity.Lint
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /lint/rule/{project_id} [get]
func (r *lintRoutes) getRule(c *gin.Context) {
	id := c.Param("project_id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		r.l.Error(err, "http - v1 - getRule")
		errorResponse(c, http.StatusBadRequest, "invalid project id")
		return
	}
	rule, err := r.t.GetRule(c.Request.Context(), int64(aid), c.GetHeader("Token"))
	if err != nil {
		r.l.Error(err, "http - v1 - getRule")
		if errors.Is(err, comerr.ErrPermission) {
			errorResponse(c, http.StatusForbidden, "permission denied")
		} else if errors.Is(err, comerr.ErrUnauthorized) {
			errorResponse(c, http.StatusUnauthorized, "unauthorized")
		} else if errors.Is(err, comerr.ErrNoRecord) {
			errorResponse(c, http.StatusNotFound, "no rule found")
		} else {
			errorResponse(c, http.StatusInternalServerError, "database problems")
		}
		return
	}

	c.JSON(http.StatusOK, rule)
}

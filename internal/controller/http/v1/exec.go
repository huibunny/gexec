package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gexec/internal/usecase"
	"gexec/pkg/logger"
)

type execRoutes struct {
	t usecase.Exec
	l logger.Interface
}

func newExecRoutes(handler *gin.RouterGroup, t usecase.Exec, l logger.Interface) {
	r := &execRoutes{t, l}

	h := handler.Group("/exec")
	{
		h.POST("/save", r.save)
		h.POST("/query", r.query)
	}
}

type saveResponse struct {
	ErrCode int `json:"errcode" example:"0"`
}

type doSaveRequest struct {
	Values [][][]interface{} `json:"values" binding:"required"  example:"[[[Alice,1,21]], [Shanghai,2000]]"`
}

type queryResponse struct {
	ErrCode int                    `json:"errcode" example:"0"`
	Result  map[string]interface{} `json:"result" example:"{name:Alice,sex:1,age:21}"`
}

type doQueryRequest struct {
	Condition map[string]interface{} `json:"values" binding:"required"  example:"{name:Alice,sex:1,age:21}"`
}

// @Summary     Save
// @Description Save system
// @ID          Save
// @Tags  	    Save
// @Accept      json
// @Produce     json
// @Param       request body doLoginRequest true "Save System"
// @Success     200 {object} loginResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /exec/Save [post]
func (r *execRoutes) save(c *gin.Context) {
	var request doSaveRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - save")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	c.JSON(http.StatusOK, saveResponse{ErrCode: 0})
}

// @Summary     Query
// @Description Query system
// @ID          Query
// @Tags  	    Query
// @Accept      json
// @Produce     json
// @Param       request body doLoginRequest true "Query System"
// @Success     200 {object} loginResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /exec/Query [post]
func (r *execRoutes) query(c *gin.Context) {
	var request doQueryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - query")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	c.JSON(http.StatusOK, queryResponse{ErrCode: 0})
}

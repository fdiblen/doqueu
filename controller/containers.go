package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fdiblen/doqueu/httputil"
	"github.com/fdiblen/doqueu/model"
)

// ListContainers godoc
// @Summary      List containers
// @Description  get containers
// @Tags         containers
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Container
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /containers [get]
func (c *Controller) ListContainers(ctx *gin.Context) {
	containers, err := model.ContainersAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, containers)
}

// ShowContainer godoc
// @Summary      Show a container
// @Description  get container by ID
// @ID           get-container-by-id
// @Tags         containers
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Container ID"
// @Success      200  {object}  model.Container
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /containers/{id} [get]
func (c *Controller) ShowContainer(ctx *gin.Context) {
	id := ctx.Param("id")
	// bid, err := strconv.Atoi(id)
	// if err != nil {
	// 	httputil.NewError(ctx, http.StatusBadRequest, err)
	// 	return
	// }

	container, err := model.ContainerOne(id)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, container)
}

// RunContainer godoc
// @Summary      Run a container
// @Description  run a container by image name
// @Tags         containers
// @Accept       multipart/form-data
// @Produce      json
// @Param        imagename formData string  true  "Docker image name"
// @Param        command   formData string  false  "Command"
// @Success      200      {object}  model.Container
// @Failure      400      {object}  httputil.HTTPError
// @Failure      404      {object}  httputil.HTTPError
// @Failure      500      {object}  httputil.HTTPError
// @Router       /containers/run [post]
func (c *Controller) RunContainer(ctx *gin.Context) {
	imagename := ctx.PostForm("imagename")
	command := ctx.PostForm("command")

	container_out, err := model.ContainerStart(imagename, command)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, container_out)
}

// StopContainer godoc
// @Summary      Stop a container
// @Description  stop a container by id
// @Tags         containers
// @Accept       multipart/form-data
// @Produce      json
// @Param        id 	  formData string  true  "container id"
// @Success      200      {object}  model.Container
// @Failure      400      {object}  httputil.HTTPError
// @Failure      404      {object}  httputil.HTTPError
// @Failure      500      {object}  httputil.HTTPError
// @Router       /containers/stop [post]
func (c *Controller) StopContainer(ctx *gin.Context) {
	id := ctx.PostForm("id")

	container_out, err := model.ContainerEnd(id)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, container_out)
}

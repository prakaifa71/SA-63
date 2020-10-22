package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/prakaifa21/app/ent"
   "github.com/prakaifa21/app/ent/patientroom"
   "github.com/gin-gonic/gin"
)
 
// PatientroomController defines the struct for the patientroom controller
type PatientroomController struct {
   client *ent.Client
   router gin.IRouter
}
// CreatePatientroom handles POST requests for adding patientroom entities
// @Summary Create patientroom
// @Description Create patientroom
// @ID create-patientroom
// @Accept   json
// @Produce  json
// @Param patientroom body ent.Patientroom true "Patientroom entity"
// @Success 200 {object} ent.Patientroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrooms [post]
func (ctl *PatientroomController) CreatePatientroom(c *gin.Context) {
    obj := ent.Patientroom{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "patientroom binding failed",
        })
        return
    }
  
    pa, err := ctl.client.Patientroom.
    Create().
    SetTyperoom(obj.Typeroom).
    Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{
            "error": "saving failed",
        })
        return
    }
  
    c.JSON(200, pa)
 }
// GetPatientroom handles GET requests to retrieve a patientroom entity
// @Summary Get a patientroom entity by ID
// @Description get patientroom by ID
// @ID get-patientroom
// @Produce  json
// @Param id path int true "Patientroom ID"
// @Success 200 {object} ent.Patientroom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrooms/{id} [get]
func (ctl *PatientroomController) GetPatientroom(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    pa, err := ctl.client.Patientroom.
        Query().
        Where(patientroom.IDEQ(int(id))).
        Only(context.Background())
    if err != nil {
        c.JSON(404, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    c.JSON(200, pa)
 }
// ListPatientroom handles request to get a list of patientroom entities
// @Summary List patientroom entities
// @Description list patientroom entities
// @ID list-patientroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Patientroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrooms [get]
func (ctl *PatientroomController) ListPatientroom(c *gin.Context) {
    limitQuery := c.Query("limit")
    limit := 10
    if limitQuery != "" {
        limit64, err := strconv.ParseInt(limitQuery, 10, 64)
        if err == nil {limit = int(limit64)}
    }
  
    offsetQuery := c.Query("offset")
    offset := 0
    if offsetQuery != "" {
        offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
        if err == nil {offset = int(offset64)}
    }
  
    patientrooms, err := ctl.client.Patientroom.
        Query().
        Limit(limit).
        Offset(offset).
        All(context.Background())
        if err != nil {
        c.JSON(400, gin.H{"error": err.Error(),})
        return
    }
  
    c.JSON(200, patientrooms)
 }
// DeletePatientroom handles DELETE requests to delete a patientroom entity
// @Summary Delete a patientroom entity by ID
// @Description get patientroom by ID
// @ID delete-patientroom
// @Produce  json
// @Param id path int true "Patientroom ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrooms/{id} [delete]
func (ctl *PatientroomController) DeletePatientroom(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    err = ctl.client.Patientroom.
        DeleteOneID(int(id)).
        Exec(context.Background())
    if err != nil {
        c.JSON(404, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
 }
// UpdatePatientroom handles PUT requests to update a patientroom entity
// @Summary Update a patientroom entity by ID
// @Description update patientroom by ID
// @ID update-patientroom
// @Accept   json
// @Produce  json
// @Param id path int true "Patientroom ID"
// @Param patientroom body ent.Patientroom true "Patientroom entity"
// @Success 200 {object} ent.Patientroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrooms/{id} [put]
func (ctl *PatientroomController) UpdatePatientroom(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    obj := ent.Patientroom{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "patientroom binding failed",
        })
        return
    }
    obj.ID = int(id)
    pa, err := ctl.client.Patientroom.
        UpdateOne(&obj).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{"error": "update failed",})
        return
    }
  
    c.JSON(200, pa)
 }
// NewPatientroomController creates and registers handles for the patientroom controller
func NewPatientroomController(router gin.IRouter, client *ent.Client) *PatientroomController {
    pac := &PatientroomController{
        client: client,
        router: router,
    }
    pac.register()
    return pac
 }
  
 // InitPatientroomController registers routes to the main engine
 func (ctl *PatientroomController) register() {
    patientrooms := ctl.router.Group("/patientrooms")
  
    patientrooms.GET("", ctl.ListPatientroom)
  
    // CRUD
    patientrooms.POST("", ctl.CreatePatientroom)
    patientrooms.GET(":id", ctl.GetPatientroom)
    patientrooms.PUT(":id", ctl.UpdatePatientroom)
    patientrooms.DELETE(":id", ctl.DeletePatientroom)
 }
      
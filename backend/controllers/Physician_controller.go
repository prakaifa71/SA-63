package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/prakaifa21/app/ent"
   "github.com/prakaifa21/app/ent/physician"
   "github.com/gin-gonic/gin"
)
 
// PhysicianController defines the struct for the physician controller
type PhysicianController struct {
   client *ent.Client
   router gin.IRouter
}



// CreatePhysician handles POST requests for adding physician entities
// @Summary Create physician
// @Description Create physician
// @ID create-physician
// @Accept   json
// @Produce  json
// @Param physician body ent.Physician true "Physician entity"
// @Success 200 {object} ent.Physician
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /physicians [post]
func (ctl *PhysicianController) CreatePhysician(c *gin.Context) {
    obj := ent.Physician{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "physician binding failed",
        })
        return
    }
  
    ph, err := ctl.client.Physician.
            Create().
            SetPhysicianname(obj.Physicianname).
            SetPhysicianemail(obj.Physicianemail).
            SetPassword(obj.Password).
            Save(context.Background())
            
    if err != nil  {
        c.JSON(400, gin.H{
            "error": "saving failed",
        })
        return
    }
  
    c.JSON(200, ph)
 }
// GetPhysician handles GET requests to retrieve a physician entity
// @Summary Get a physician entity by ID
// @Description get physician by ID
// @ID get-physician
// @Produce  json
// @Param id path int true "Physician ID"
// @Success 200 {object} ent.Physician
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /physicians/{id} [get]
func (ctl *PhysicianController) GetPhysician(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    ph, err := ctl.client.Physician.
        Query().
        Where(physician.IDEQ(int(id))).
        Only(context.Background())
    if err != nil {
        c.JSON(404, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    c.JSON(200, ph)
 }
// ListPhysician handles request to get a list of physician entities
// @Summary List physician entities
// @Description list physician entities
// @ID list-physician
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Physician
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /physicians [get]
func (ctl *PhysicianController) ListPhysician(c *gin.Context) {
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
  
    physicians, err := ctl.client.Physician.
        Query().
        Limit(limit).
        Offset(offset).
        All(context.Background())
        if err != nil {
        c.JSON(400, gin.H{"error": err.Error(),})
        return
    }
  
    c.JSON(200, physicians)
 }
// DeletePhysician handles DELETE requests to delete a physician entity
// @Summary Delete a physician entity by ID
// @Description get physician by ID
// @ID delete-physician
// @Produce  json
// @Param id path int true "Physician ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /physicians/{id} [delete]
func (ctl *PhysicianController) DeletePhysician(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    err = ctl.client.Physician.
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
// UpdatePhysician handles PUT requests to update a physician entity
// @Summary Update a physician entity by ID
// @Description update physician by ID
// @ID update-physician
// @Accept   json
// @Produce  json
// @Param id path int true "Physician ID"
// @Param physician body ent.Physician true "Physician entity"
// @Success 200 {object} ent.Physician
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /physicians/{id} [put]
func (ctl *PhysicianController) UpdatePhysician(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    obj := ent.Physician{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "physician binding failed",
        })
        return
    }
    obj.ID = int(id)
    ph, err := ctl.client.Physician.
        UpdateOne(&obj).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{"error": "update failed",})
        return
    }
  
    c.JSON(200, ph)
 }
// NewPhysicianController creates and registers handles for the physician controller
func NewPhysicianController(router gin.IRouter, client *ent.Client) *PhysicianController {
    phc := &PhysicianController{
        client: client,
        router: router,
    }
    phc.register()
    return phc
 }
  
 // InitPhysicianController registers routes to the main engine
 func (ctl *PhysicianController) register() {
    physicians := ctl.router.Group("/physicians")
  
    physicians.GET("", ctl.ListPhysician)
  
    // CRUD
    physicians.POST("", ctl.CreatePhysician)
    physicians.GET(":id", ctl.GetPhysician)
    physicians.PUT(":id", ctl.UpdatePhysician)
    physicians.DELETE(":id", ctl.DeletePhysician)
 }
      
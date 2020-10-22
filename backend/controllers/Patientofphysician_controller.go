package controllers

import (
	"context"
	"fmt"
	"strconv"
	

	"github.com/prakaifa21/app/ent"
	"github.com/prakaifa21/app/ent/patientroom"
	"github.com/prakaifa21/app/ent/patient"
	"github.com/prakaifa21/app/ent/physician"
	"github.com/gin-gonic/gin"
)

// PatientofphysicianController defines the struct for the patientofphysician controller
type PatientofphysicianController struct {
	client *ent.Client
	router gin.IRouter
}

// Patientofphysician defines the struct for the patientofphysician
type Patientofphysician struct {
	PatientID 	int
	PhysicianID int
	RoomID     	int
	Ailment  	string
}

// CreatePatientofphysician handles POST requests for adding patientofphysician entities
// @Summary Create patientofphysician
// @Description Create patientofphysician
// @ID create-patientofphysician
// @Accept   json
// @Produce  json
// @Param patientofphysician body Patientofphysician true "Patientofphysician entity"
// @Success 200 {object} ent.Patientofphysician
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientofphysicians [post]
func (ctl *PatientofphysicianController) CreatePatientofphysician(c *gin.Context) {
	obj := Patientofphysician{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "patientofphysician binding failed",
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.PatientID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	ph, err := ctl.client.Physician.
		Query().
		Where(physician.IDEQ(int(obj.PhysicianID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "physician   not found",
		})
		return
	}

	pa, err := ctl.client.Patientroom.
		Query().
		Where(patientroom.IDEQ(int(obj.RoomID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "room not found",
		})
		return
	}
	

	pp, err := ctl.client.Patientofphysician.
		Create().
		SetIdpatient(p).
		SetIdphysicianid(ph).
		SetRoomid(pa).
		SetAilment(obj.Ailment).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pp)
}

// ListPatientofphysician handles request to get a list of patientofphysician entities
// @Summary List patientofphysician entities
// @Description list patientofphysician entities
// @ID list-patientofphysician
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Patientofphysician
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientofphysicians [get]
func (ctl *PatientofphysicianController) ListPatientofphysician(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	patientofphysicians, err := ctl.client.Patientofphysician.
		Query().
		WithIdpatient().
		WithIdphysicianid().
		WithRoomid().
		Limit(limit).
		Offset(offset).
		
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, patientofphysicians)
}

// DeletePatientofphysician handles DELETE requests to delete a patientofphysician entity
// @Summary Delete a patientofphysician entity by ID
// @Description get patientofphysician by ID
// @ID delete-patientofphysician
// @Produce  json
// @Param id path int true "Patientofphysician ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientofphysicians/{id} [delete]
func (ctl *PatientofphysicianController) DeletePatientofphysician(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Patientofphysician.
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

// NewPatientofphysicianController creates and registers handles for the patientofphysician controller
func NewPatientofphysicianController(router gin.IRouter, client *ent.Client) *PatientofphysicianController {
	drc := &PatientofphysicianController{
		client: client,
		router: router,
	}
	drc.register()
	return drc
}

// InitUserController registers routes to the main engine
func (ctl *PatientofphysicianController) register() {
	patientofphysicians := ctl.router.Group("/patientofphysicians")

	patientofphysicians.GET("", ctl.ListPatientofphysician)

	// CRUD
	patientofphysicians.POST("", ctl.CreatePatientofphysician)
	patientofphysicians.DELETE(":id", ctl.DeletePatientofphysician)
}
package main 
import (
   "context"
   "log"
 
   "github.com/prakaifa21/app/controllers"
   _ "github.com/prakaifa21/app/docs"
   "github.com/prakaifa21/app/ent"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)


type Patients struct {
	Patient []Patient
}

type Patient struct {
	Patientname  string
	
}

type Physicians struct {
	Physician []Physician
}

type Physician struct {
    Physicianname  string
    Physicianemail string
    Password       string
	
}

type Patientrooms struct {
	Patientroom []Patientroom
}

type Patientroom struct {
	Typeroom  string
	
}


// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/
 
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
 
// @host localhost:8080
// @BasePath /api/v1
 
// @securityDefinitions.basic BasicAuth
 
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
 
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
   router := gin.Default()
   router.Use(cors.Default())
 
   client, err := ent.Open("sqlite3", "file:ent.db?&cache=shared&_fk=1")
   if err != nil {
       log.Fatalf("fail to open sqlite3: %v", err)
   }
   defer client.Close()
 
   if err := client.Schema.Create(context.Background()); err != nil {
       log.Fatalf("failed creating schema resources: %v", err)
   }
 
   v1 := router.Group("/api/v1")
   controllers.NewPatientofphysicianController(v1, client)
   controllers.NewPatientController(v1, client)
   controllers.NewPhysicianController(v1, client)
   controllers.NewPatientroomController(v1, client)
 
   patients := Patients{
        Patient: []Patient{
                  Patient{"Thanabodee Petchrey"},
                  Patient{"Poomin Pimpimai"}, 
                  Patient{"Chanwit Kaewkasi"},
                  Patient{"Apichaya Thadee"},
                  Patient{"Supaporn Boonin"},
                  Patient{"Naruephon Phanthet"},
                  Patient{"Wirat Nimpao"},
                  Patient{"Jurarat Homsai"},
                  Patient{"Tiwakorn Sasen"},
                  Patient{"Konrawit Kongsri"},

        },
    }

    for _, p := range patients.Patient {
		client.Patient.
			Create().
			SetPatientname(p.Patientname).
			Save(context.Background())
	}

    
   patientrooms := Patientrooms{
        Patientroom: []Patientroom{
               Patientroom{"ห้องเดี่ยว"},
               Patientroom{"ห้องรวม"}, 
               

            },
    }

    for _, pa := range patientrooms.Patientroom {
        client.Patientroom.
            Create().
            SetTyperoom(pa.Typeroom).
            Save(context.Background())
 }

    
   physicians := Physicians{
        Physician: []Physician{
               Physician{"Prakaifa kummungkun","prakaifa@gmail.com","dsfgdh123"},
               Physician{"Noppawan Khunprawong","Noppawan@gmail.com","dsdh123"}, 
               Physician{"Nantiya Prabmeechai","Nantiya@gmail.com","dsfgh123"},
               Physician{"Pradchaya Srisan","Pradchaya@gmail.com","dsfgdh3"},
               Physician{"Parinthon Khotyotee","Parinthon@gmail.com","dsfg3"},
               Physician{"Pongsit Rimsakul","Pongsit@gmail.com","dsf123"},
               Physician{"Supakorn Riyapa","Supakorn@gmail.com","dsfgd3"},
               Physician{"Nattapong Poonperm","Nattapong@gmail.com","dsh123"},
               Physician{"Kittipong Phanugram","Kittipong@gmail.com","dsfg23"},
               Physician{"Subancha Sawaddee","Subancha@gmail.com","dh123"},
        },
 }

    for _, ph := range physicians.Physician {
        client.Physician.
            Create().
            SetPhysicianname(ph.Physicianname).
            SetPhysicianemail(ph.Physicianemail).
            SetPassword(ph.Password).
            Save(context.Background())
 }






   
   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   router.Run()
}

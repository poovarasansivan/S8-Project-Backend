package main

import (
	"fmt"
	"server/config"
	"server/handlers"
	"server/middleware"
	"server/routes"
	"server/routes/fullstack"
	"server/routes/home"
	"server/routes/inputdata"
	"server/routes/psdata"
	"server/routes/student"
	"server/routes/technical"

	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config.ConnectDB()
	defer config.Database.Close()
	router := mux.NewRouter()

	router.HandleFunc("/", routes.Sample).Methods("GET")
	router.HandleFunc("/signin", handlers.Signin).Methods("POST")

	protected := router.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// Admin and Faculty Pages API Endpoints
	protected.HandleFunc("/home", home.GetAdminHomeData).Methods("GET")
	protected.HandleFunc("/admin-home-academic-data", home.GetAdminHomeAcademicData).Methods("GET")
	protected.HandleFunc("/admin-home-fullstack-data", home.GetAdminHomeFullStackData).Methods("GET")
	protected.HandleFunc("/get-individual-rollno", fullstack.GetAllStudentsRollno).Methods("GET")
	protected.HandleFunc("/get-inidividual-fullstack-card/{roll_no}", fullstack.GetIndividualStudentFullStackData).Methods("GET")
	protected.HandleFunc("/get-inidividual-projectdetails/{roll_no}", fullstack.GetIndividualStudentProjectDetails).Methods("GET")
	protected.HandleFunc("/get-inidividual-review/{roll_no}", fullstack.GetIndividualStudentReviewData).Methods("GET")
	protected.HandleFunc("/get-admin-pscard-data/{roll_no}", psdata.GetAdminPsCardData).Methods("GET")
	protected.HandleFunc("/get-psdata-rollno", psdata.GetPSDataRollNo).Methods("GET")
	protected.HandleFunc("/get-individual-psslotdata/{roll_no}", psdata.GetIndividualStudentPSSlotData).Methods("GET")
	protected.HandleFunc("/get-psslotdata-completion/{roll_no}", psdata.GetFacultyViewPSLevelCompletion).Methods("GET")
	protected.HandleFunc("/get-psprofiledata/{roll_no}", psdata.GetStudentAcademicDataForFacultyView).Methods("GET")
	protected.HandleFunc("/add-students-profile-data", home.AddStudentsProfileData).Methods("POST")
	protected.HandleFunc("/get-students-profile-data", inputdata.GetStudentsProfileData).Methods("GET")
	protected.HandleFunc("/get-students-academic-data", inputdata.GetAcademicData).Methods("GET")
	protected.HandleFunc("/get-psslot-data", inputdata.GetPsSlotData).Methods("GET")
	protected.HandleFunc("/get-fullstackproject-data", inputdata.GetFullStackProjectData).Methods("GET")
	protected.HandleFunc("/get-fullstackreview-data", inputdata.GetFullStackReviewData).Methods("GET")
	protected.HandleFunc("/get-codingplatform-data/{roll_no}", technical.GetCodingPlatformsUsername).Methods("GET")
	protected.HandleFunc("/get-student-rollno", technical.GetStudentRollno).Methods("GET")
	protected.HandleFunc("/add-suggestion", home.InsertSuggestion).Methods("POST")
	protected.HandleFunc("/get-suggestion/{roll_no}", home.GetIndividualSuggestion).Methods("GET")
	protected.HandleFunc("/update-suggestion", home.UpdateSuggestionStatus).Methods("PUT")
	protected.HandleFunc("/get-studentprofiledata/{roll_no}", student.GetStudentProfileCardData).Methods("GET")
	protected.HandleFunc("/get-studentcarddata/{roll_no}", student.GetStudentCardData).Methods("GET")
	protected.HandleFunc("/get-placementdetails/{roll_no}", student.GetStudentPlacementDetails).Methods("GET")
	protected.HandleFunc("/get-studentsgpascore/{roll_no}", student.GetStudentSGPAScore).Methods("GET")
	protected.HandleFunc("/get-profiledetails/{roll_no}", student.GetStudentPersonalDetails).Methods("GET")
	protected.HandleFunc("/get-educationdetails/{roll_no}", student.GetStudentEducationProfile).Methods("GET")
    protected.HandleFunc("/get-pscarddata/{roll_no}", psdata.GetStudentPSCardData).Methods("GET")
	protected.HandleFunc("/get-pscard/{roll_no}", psdata.GetStudentPSData).Methods("GET")
	protected.HandleFunc("/get-pointhistory/{roll_no}", psdata.GetSudentPsPointHistory).Methods("GET")
    protected.HandleFunc("/get-pschartdata/{roll_no}", psdata.GetPSlevelsData).Methods("GET")
    protected.HandleFunc("/get-fullstackcarddata/{roll_no}", fullstack.GetStudentFullStackCardData).Methods("GET")
    protected.HandleFunc("/get-fullstackproject-details/{roll_no}", fullstack.GetProjectDetails).Methods("GET")
    protected.HandleFunc("/get-fullstackproject-currentlevel/{roll_no}", fullstack.GetFullStackCurrentlevel).Methods("GET")


	c := cors.AllowAll()
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", c.Handler(router))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

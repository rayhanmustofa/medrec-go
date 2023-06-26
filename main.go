package main

import (
	"log"
	"net/http"

	"github.com/your-username/patient-app/pkg/handlers"
)

func main() {
	patientHandler := handlers.NewPatientHandler()

	http.HandleFunc("/patients", patientHandler.ListPatientsHandler)
	http.HandleFunc("/patients/new", patientHandler.CreatePatientHandler)
	http.HandleFunc("/patients/edit", patientHandler.EditPatientHandler)
	http.HandleFunc("/patients/delete", patientHandler.DeletePatientHandler)

	log.Println("Server is running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

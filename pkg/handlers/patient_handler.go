package handlers

import (
	patient "command-line-argumentsC:\\Users\\Rayhan\\OneDrive\\Dokumen\\_dev\\medrec-go\\pkg\\patient\\patient.go"
	"net/http"
	"text/template"
)

type PatientHandler struct {
	Patients []patient.Patient
	Tpl      *template.Template
}

func NewPatientHandler() *PatientHandler {
	return &PatientHandler{
		Tpl: template.Must(template.ParseFiles("templates/index.html", "templates/patient.html")),
	}
}

func (ph *PatientHandler) ListPatientsHandler(w http.ResponseWriter, r *http.Request) {
	err := ph.Tpl.ExecuteTemplate(w, "index.html", ph.Patients)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ph *PatientHandler) CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}

func (ph *PatientHandler) EditPatientHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}

func (ph *PatientHandler) DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}

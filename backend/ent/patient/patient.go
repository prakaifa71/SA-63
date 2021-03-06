// Code generated by entc, DO NOT EDIT.

package patient

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPatientname holds the string denoting the patientname field in the database.
	FieldPatientname = "patientname"

	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "Patient"

	// Table holds the table name of the patient in the database.
	Table = "patients"
	// PatientTable is the table the holds the Patient relation/edge.
	PatientTable = "patientofphysicians"
	// PatientInverseTable is the table name for the Patientofphysician entity.
	// It exists in this package in order to avoid circular dependency with the "patientofphysician" package.
	PatientInverseTable = "patientofphysicians"
	// PatientColumn is the table column denoting the Patient relation/edge.
	PatientColumn = "patient_patient"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldPatientname,
}

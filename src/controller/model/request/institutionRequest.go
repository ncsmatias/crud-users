package request

type InstitutionRequest struct {
	InstitutionType string `json:"institution_type" binding:"required,oneof=company university"`
	Name            string `json:"name" binding:"required"`
	Phone           string `json:"phone" binding:"max=128"`
	ZipCode         string `json:"zip_code"`
	Street          string `json:"street"`
	Number          string `json:"number"`
	Neighborhood    string `json:"neighborhood"`
	City            string `json:"city"`
	State           string `json:"state"`
	UF              string `json:"UF"`
	Country         string `json:"country"`
	CountryCode     string `json:"country_code"`
}

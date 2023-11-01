package request

type AddressRequest struct {
	Street       string `json:"street" binding:"required"`
	Neighborhood string `json:"neighborhood" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	ZipCode      string `json:"zip_code" binding:"required"`
	Notes        string `json:"notes"`
}

package track

import "fmt"

type TrackingStatus string

const (
	PrisEnCharge        TrackingStatus = "PRIS_EN_CHARGE"
	EnLivraison         TrackingStatus = "EN_LIVRAISON"
	Expedie             TrackingStatus = "EXPEDIE"
	ARetirer            TrackingStatus = "A_RETIRER"
	TriEffectue         TrackingStatus = "TRI_EFFECTUE"
	Distribue           TrackingStatus = "DISTRIBUE"
	Livre               TrackingStatus = "LIVRE"
	DestinataireInforme TrackingStatus = "DESTINATAIRE_INFORME"
	RetourDestinataire  TrackingStatus = "RETOUR_DESTINATAIRE"
	Erreur              TrackingStatus = "ERREUR"
	Inconnu             TrackingStatus = "INCONNU"
)

type Response struct {
	Code    string         `json:"code"`
	Date    string         `json:"date,omitempty"` // DD/MM/YYYY
	Status  TrackingStatus `json:"date,omitempty"`
	Message string         `json:"message"`
	Link    string         `json:"string,omitempty"`
	Type    string         `json:"string,omitempty"`
}

type ErrorCode string

const (
	ErrForbidden ErrorCode = "FORBIDDEN"
	ErrServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"
)

type ErrorResponse struct {
	Code ErrorCode `json:"code"`
	Message string `json:"message"`
}

func (ec ErrorResponse) Error() string {
	return fmt.Sprintf("Remote returned error with code \"%s\" and message : %s", ec.Code, ec.Message)
}
package track

import (
	"path"
	"github.com/aabizri/laposte"
	"net/http"
	"encoding/json"
	"errors"
)

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
	/*
	En cas de succès: le code donné
	En cas d'échec: le code de l'erreur
	 */
	Code    string         `json:"code"`
	Date    string         `json:"date,omitempty"` // DD/MM/YYYY
	Status  TrackingStatus `json:"date,omitempty"`
	Message string         `json:"message"`
	Link    string         `json:"string,omitempty"`
	Type    string         `json:"string,omitempty"`
}

var (
	ErrNotFound error = errors.New("resource not found")
	ErrBadCode  error = errors.New("bad code")
)

func (cl *Client) Track(id string) (*Response, error) {
	endpoint := "https://" + path.Join(laposte.APIHost, root, version, id)

	// Create the request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Authenticate
	cl.auth(req)

	// Execute
	httpResp, err := cl.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Parse
	dec := json.NewDecoder(httpResp.Body)
	resp := &Response{}
	err = dec.Decode(resp)
	if err != nil {
		return nil, err
	}

	// Switch on responses
	switch httpResp.StatusCode {
	case 400:
		return resp, ErrBadCode
	case 404:
		return resp, ErrNotFound
	}

	// Else, we're done
	return resp, nil
}

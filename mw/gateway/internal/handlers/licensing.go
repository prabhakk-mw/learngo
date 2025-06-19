package handlers

import "net/http"

// LicensingService APIs
//
//	@Summary		Lists the licensing options available to use for an environment.
//	@Description	List of available licensing options.
//	@Tags			Licensing
//	@Param			env_id	path	string	true	"Environment ID"
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/licensing/options/{env_id} [get]
func (handlers *Handlers) AvailableLicensingOptions(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}

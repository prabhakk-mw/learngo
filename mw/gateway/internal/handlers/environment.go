package handlers

import "net/http"

// EnvironmentService APIs
//
//	@Summary		Lists the MATLAB Installations found on the device.
//	@Description	Use this to find the MATLABs available to use.
//	@Tags			Environment
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/environment/list [get]
func (handlers *Handlers) ListMATLABs(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}

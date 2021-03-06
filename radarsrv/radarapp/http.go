package radarapp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO: Development only code - should only use RPC services in final thing
func NewRadarService() http.ServeMux {
	router := http.ServeMux{}
	router.Handle(
		"/",
		http.FileServer(
			http.Dir("./radarapp/static")))

	router.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		p, err := GlobalGameState.GetPlayerPositions()
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.WriteHeader(http.StatusOK)
		j, err := json.Marshal(*p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintf(w, string(j))
	})

	return router
}

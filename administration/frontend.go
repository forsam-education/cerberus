package administration

import (
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"github.com/gobuffalo/packr/v2"
	"net/http"
)

type spaHandler struct {
	box       *packr.Box
	indexPath string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	utils.LogVerbose(fmt.Sprintf("Received request on Web Administration route %s from %s", r.URL.Path, r.RemoteAddr), nil)

	if !h.box.Has(r.URL.Path) || r.URL.Path == "/" {
		indexContent, err := h.box.Find(h.indexPath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if _, err := w.Write(indexContent); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	http.FileServer(h.box).ServeHTTP(w, r)
}

package httperror

import (
	"fmt"
	"net/http"
	"strings"
)

// HandleMethod handles a method error.
func HandleMethod(w http.ResponseWriter, err Method) {
	w.Header().Set("Allow", strings.Join(err.Allowed, ", "))
	error := fmt.Sprintf("405 Method Not Allowed\n\n%v", err)
	http.Error(w, error, http.StatusMethodNotAllowed)
}

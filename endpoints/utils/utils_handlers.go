package utils
import "net/http"

type Service struct {
}

func (s Service) GetAllNamespaces(w http.ResponseWriter, r *http.Request) {
  defaultResponse(w, http.StatusOK, []byte("{\"Said\":\"Hello\"}"))
}

func (s Service) GetNamespaceById(w http.ResponseWriter, r *http.Request) {
  defaultResponse(w, http.StatusOK, []byte("{\"Said\":\"Hello\"}"))
}

func defaultResponse(w http.ResponseWriter, code int, resp []byte) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(resp)
}

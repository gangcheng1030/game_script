package chaojidou

import (
	"fmt"
	"net/http"
)

type LeaderHandler struct {
}

func (lh *LeaderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	addr := r.Form.Get("addr")
	fmt.Printf("register : %s \n", addr)
	if len(addr) > 0 && !lh.contains(Follwers, addr) {
		Follwers = append(Follwers, addr)
	}

	w.WriteHeader(http.StatusOK)
}

func (lh *LeaderHandler) contains(followers []string, addr string) bool {
	for _, follower := range followers {
		if follower == addr {
			return true
		}
	}

	return false
}

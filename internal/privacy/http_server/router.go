package http_server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"lider/internal/privacy"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(wl privacy.WhiteUserList) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	//r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)

	r.Get("/whitelist/{tgId}", whiteListUserCheck(wl))
	r.Put("/whitelist", whiteListAdd(wl))
	return r
}

func whiteListAdd(wl privacy.WhiteUserList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data []byte = make([]byte, 64)
		usdata := make([]byte, 0, 64)
		var err error
		for {
			_, err = r.Body.Read(data)
			if err != nil {
				break
			}
			usdata = append(usdata, data...)
		}
		if err != io.EOF {
			goto END
		}

		{
			var out privacy.WhiteListEntry
			err = json.Unmarshal(usdata, &out)
			if err != nil {
				goto END
			}

			err = wl.Add(out)
			if err != nil {
				goto END
			}
			return
		}
	END:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//TODO log
			fmt.Println(err)
			return
		}
	}
}

func whiteListUserCheck(wl privacy.WhiteUserList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tgId := chi.URLParam(r, "tgId")
		userId, err := strconv.ParseInt(tgId, 10, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		wlEntry, err := wl.FindByID(privacy.UserID(userId))
		if err != nil {
			if err == privacy.ErrNotFound {
				w.WriteHeader(http.StatusNotFound)
			} else {
				// TODO Log error
				w.WriteHeader(http.StatusInternalServerError)

			}
		}
		jsn, err := json.Marshal(wlEntry)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO Log error
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsn)
	}
}

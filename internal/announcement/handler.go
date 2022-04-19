package announcement

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	announcementsURL = "/api/v1/announcements"
	announcementURL  = "/api/v1/announcements/"
)

type (
	Handler struct {
		announcement Service
	}

	reqAnnouncementByID struct {
		ID      int64    `form:"uuid"`
		Options []string `form:"fields"`
	}

	reqListAnnouncement struct {
		Page   int    `form:"page"`
		Limit  int    `form:"limit"`
		Sort   string `form:"sort"`
		SortBy string `form:"sortby"`
	}
)

func NewHandler(service Service) *Handler {
	return &Handler{
		announcement: service,
	}
}

// Register registered http routes for services
func (h *Handler) Register(mux *http.ServeMux) {
	mux.Handle("/", h)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		r.ParseForm()
		if r.Form.Get("uuid") != "" {
			h.getAnnouncement(w, r)
			return
		}

		h.getListAnnouncement(w, r)
		return

	case "POST":
		if r.URL.Path == announcementsURL {
			h.createAnnouncement(w, r)
			return
		}

	case "PATCH":
		if r.URL.Path == announcementURL {
			h.addImage(w, r)
			return
		}
	}

	http.NotFound(w, r)

	return
}

// createAnnouncement trie create announcement
func (h *Handler) createAnnouncement(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	add, err := h.announcement.Add(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(int(add))))
}

func (h *Handler) getAnnouncement(w http.ResponseWriter, r *http.Request) {
	q := &reqAnnouncementByID{}
	err := parseForm(q, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	announcementByUID, err := h.announcement.GetById(q.ID, q.Options)
	if err != nil {
		if err.Error() == "id not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(announcementByUID)
}

func (h *Handler) getListAnnouncement(w http.ResponseWriter, r *http.Request) {
	q := &reqListAnnouncement{}
	err := parseForm(q, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	list, err := h.announcement.GetList(q.Page, q.Limit, q.SortBy, q.Sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(list)
}

func (h *Handler) addImage(w http.ResponseWriter, r *http.Request) {
	id, path, err := fileSave(r)
	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	bytes, err := h.announcement.AddImage(id, path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

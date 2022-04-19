package announcement

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-playground/form/v4"
	json "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

func fileSave(r *http.Request) (int, string, error) {
	// left shift 32 << 20 which results in 32*2^20 = 33554432
	// x << y, results in x*2^y
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return 0, "", err
	}
	idStr := r.Form.Get("uuid")
	iaAn := r.Form.Get("path")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, "", err
	}

	f, h, err := r.FormFile("file")
	if err != nil {
		return 0, "", err
	}
	defer f.Close()
	path := filepath.Join(".", "files", idStr)
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + iaAn + filepath.Ext(h.Filename)
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, "", err
	}
	defer file.Close()
	// Copy the file to the destination path
	_, err = io.Copy(file, f)
	if err != nil {
		return 0, "", err
	}
	return id, fullPath, nil
}

// parseForm tries to parse and decode specified data.
func parseForm(v interface{}, r *http.Request) error {
	// try to parse the http request form data
	if err := r.ParseForm(); err != nil {
		return err
	}
	if len(r.Form) > 0 {
		if err := form.NewDecoder().Decode(v, r.Form); err != nil {
			return err
		}
	}

	return nil
}

func parseHTTPRequest(v interface{}, r *http.Request) (interface{}, error) {
	// make sure the http request is not empty
	if r.ContentLength == 0 {
		return nil, errors.New("empty request")
	}

	// try to parse the http request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body")
	}

	if len(body) > 0 {
		if err = json.ConfigFastest.Unmarshal(body, &v); err != nil {
			return nil, errors.Wrap(err, "unmarshal body")
		}
	} else {
		// try to parse the http request form
		if err = r.ParseForm(); err != nil {
			return nil, errors.Wrap(err, "parse form")
		}
		if len(r.Form) > 0 {
			if err = form.NewDecoder().Decode(v, r.Form); err != nil {
				return nil, errors.Wrap(err, "decode form")
			}
		}
	}

	return v, nil
}

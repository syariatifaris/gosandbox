package handler

import (
	"io"
	"net/http"

	"encoding/json"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/syariatifaris/gosandbox/core/config"
)

const (
	OperationSuccess = "Success"
	OperationFailed  = "Failed"
)

type THandler interface {
	Name() string
	RegisterHandlers(router *mux.Router)
}

type MuxServer interface {
	http.Handler
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

type baseHandler struct {
	config *config.ConfigurationData
	db     interface{}
}

type apiResult struct {
	//data  interface{} `json:"data"`
	error string `json:"error"`
}

func NewResult(d interface{}, err error) apiResult {
	return apiResult{
		//data:  d,
		error: err.Error(),
	}
}

type APIResult struct {
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error_message"`
}

func render(w http.ResponseWriter, data interface{}, err error) {
	if data == nil {
		data = ""
	}

	errMessage := ""
	if err != nil {
		errMessage = err.Error()
	}

	d, err := json.Marshal(APIResult{
		Data:         data,
		ErrorMessage: errMessage,
	})

	if err != nil {
		w.Write([]byte(fmt.Sprintf("service unavailable, err: %s", err.Error())))
	}

	w.Write(d)
}

func getPostData(r *http.Request, v interface{}) error {
	if r.Method == http.MethodPost {
		//body, err := ioutil.ReadAll(r.Body)
		//fmt.Println(string(body))
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("method %s not allowed %d", r.Method, http.StatusMethodNotAllowed)
	}

	return nil
}

func parseData(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	return err
}

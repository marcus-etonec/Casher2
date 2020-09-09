package endpoints

import (
	"encoding/json"
	"net/http"

	structs "etoneclab.com/basic-server/data"
	"github.com/gorilla/mux"
)

// Definition of some static error
const EverythingOk uint = 0
const MissingParameters uint = 1
const RecordNotFound uint = 2
const WrongParameter uint = 3
const RecordNotAdded uint = 4
const Timeout uint = 5

// Definition of the ApiResponseWriter
type ApiResponseWriter struct {
	http.ResponseWriter
}

type Parameters struct {
	Token string
	Id    string
}

type BaseResponse struct {
	ID   uint
	Data interface{}
}

// This function writes an error to the corresponding responsewriter.
// Should be improved by adding log functionality and other header specific functionality.
func (w ApiResponseWriter) WriteError(error string) {
	w.Write([]byte(error))
}

// This function sets the common headers for all the responses.
func (w ApiResponseWriter) SetHeaders() {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

// Format the response to a more digestible response for our API : Error code, and Response...
func PrepareResponse(ID uint, data interface{}) BaseResponse {
	var response BaseResponse
	response.Data = data
	response.ID = ID
	return response
}

// NotFound returns a record
func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotFound)
}

// Home Self explantory function which returns a plain text as JSON.
func Home(w ApiResponseWriter, r *http.Request, params *Parameters) {
	var test = structs.Transaction{}
	var testarray = []structs.Transaction{}
	test.Id = 0
	test.Name = "John Doe"
	test.Registered = "Guest"
	test.Status = "Pending"
	testarray = append(testarray, test)
	response := PrepareResponse(EverythingOk, testarray)
	json.NewEncoder(w).Encode(response)
}

// Function to check a valid token to respond as API
func checkValidToken(token string) bool {
	return true
}

// Function to check permission to respond to API
func checkPermissions(token string, permissions []string) bool {
	return true
}

//Reads all parameters from the endpoint request, and puts them in the Parameters structure.
func SetParameters(r *http.Request, params *Parameters) {
	vars := mux.Vars(r)
	// SAMPLE id as parameter
	v := vars["id"]
	params.Id = v
}

// MakeRequest This function parses parameters, creates a ApiResponseWriter and call the function registered
func MakeRequest(handler func(ApiResponseWriter, *http.Request, *Parameters), w http.ResponseWriter, r *http.Request, public bool, permissions []string) {
	var params Parameters
	var wloc ApiResponseWriter

	token := r.Header.Get("X-Token")

	if !public {
		ret := checkValidToken(token)
		if ret == false {
			http.Error(w, "", http.StatusNotFound)
			return
		}

		ret = checkPermissions(token, permissions)
		if ret == false {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
	}

	wloc.ResponseWriter = w
	params.Token = token
	SetParameters(r, &params)
	wloc.SetHeaders()
	handler(wloc, r, &params)
}

// RegisterCall Creates all aliases of the endpoints
// mxRouter : the router gorilla mux
// URLRndPoint: the endpoint to repond to
// FunctionEndPoint: the local function to execute when endpoint is hit
// method: the type of method to respond to (GET, PUT etc.)
// public: if public is true, there are no checks on tokens or permissions, i.e. LOGIN or REGISTER are public endpoints
// permissions: an array of strings defining the permissions to allow the endpoint to respond
func RegisterCall(muxRouter *mux.Router, URLEndPoint string, FunctionEndPoint func(ApiResponseWriter, *http.Request, *Parameters), method string, public bool, permissions []string) {
	muxRouter.HandleFunc(URLEndPoint, func(w http.ResponseWriter, r *http.Request) {
		MakeRequest(FunctionEndPoint, w, r, public, permissions)
	}).Methods(method, "OPTIONS")
}

func PrepareEndpoints(muxRouter *mux.Router) {
	muxRouter.NotFoundHandler = http.HandlerFunc(NotFound)
	RegisterCall(muxRouter, "/", Home, "GET", true, nil)
}

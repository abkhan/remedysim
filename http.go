package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const surl = ":8088"

var currentInc int = 2345
var currentWLG int = 6745

func nonServe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p><h1>No server at /</h1>")
}

func incidents(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Incident:\n  Incoming Headers: \n%+v\n", r.Header)
	fmt.Printf("   Incoming Body: \n%s\n", reqBody)

	xid := fmt.Sprintf("INC000000%d", currentInc)
	currentInc++

	locstr := fmt.Sprintf("/enterprise/service-management/v1/incidents/%s", xid)
	respuid := r.Header["Content-Tracking-Id"]
	xrid := time.Now().String() + "_" + respuid[0][5:16]

	w.Header().Set("location", locstr)
	w.Header().Set("server", "Microsoft-IIS/10.0")
	w.Header().Set("x-id", xid)
	w.Header().Set("x-request-id", xrid)
	w.Header().Set("x-tracking-id", respuid[0])
	w.Header().Set("x-uri-documentation", "https://api.wecenergygroupstg.com/enterprise/service-management/swagger/ui/index#/IncidentV1")
	w.Header().Set("x-powered-by", "ASP.NET")
	w.Header().Set("date", time.Now().String())
}

func workdetails(w http.ResponseWriter, r *http.Request) {
	uvars := mux.Vars(r)

	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Work Details:\n   Incoming Headers: \n%+v", r.Header)
	fmt.Printf("\n   Work details Vars: \n%+v\n", uvars)
	fmt.Printf("   Work details Body: \n%s\n", reqBody)

	xid := fmt.Sprintf("WLG000000%d", currentWLG)
	currentWLG++

	locstr := fmt.Sprintf("/enterprise/service-management/v1/incidents/INCxyz/work-details/%s", xid)
	respuid := r.Header["Content-Tracking-Id"]
	xrid := time.Now().String() + "_" + respuid[0][5:16]

	w.Header().Set("location", locstr)
	w.Header().Set("server", "Microsoft-IIS/10.0")
	w.Header().Set("x-request_id", xrid)
	w.Header().Set("x-tracking-id", respuid[0])
	w.Header().Set("x-uri-documentation", "https://api.wecenergygroupstg.com/enterprise/service-management/swagger/ui/index#/IncidentV1")
	w.Header().Set("x-powered-by", "ASP.NET")
	w.Header().Set("date", time.Now().String())
}

func main() {
	grr := mux.NewRouter()

	grr.HandleFunc("/incidents", incidents)
	grr.HandleFunc("/incidents/INC{inc:[0-9]+}/work-details", workdetails)

	grr.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.Handle("/", grr)

	fmt.Printf("Started for incidents and work-details\n")
	fmt.Printf(" - waiting on %s\n=>> \n", surl)
	log.Fatal(http.ListenAndServe(surl, nil))
}

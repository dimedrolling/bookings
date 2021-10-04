package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/dimedrolling/bookings/internal/config"
	"github.com/dimedrolling/bookings/internal/models"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	//what im going to put in session
	gob.Register(models.Reservation{})
	//change this to true while in production
	testApp.InProduction = false

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infolog

	errlog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errlog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = testApp.InProduction

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWrite struct{}

func (tw *myWrite) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWrite) WriteHeader(i int) {

}

func (tw *myWrite) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

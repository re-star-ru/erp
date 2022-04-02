package pkg

import (
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"net/http"
	"net/url"
	"runtime"
	"strings"
)

var (
	ErrWrongInput = errors.New("wrong input")
)

// SendErrorJSON makes {error: blah, message: blah, code: 42} json body and responds with error code
func SendErrorJSON(w http.ResponseWriter, r *http.Request, httpStatusCode int, err error, details string) {
	log.Warn().Err(err).Msg(errDetailsMsg(r, httpStatusCode, details))
	render.Status(r, httpStatusCode)
	render.JSON(w, r, JSON{"error": err.Error(), "message": details})
}

func errDetailsMsg(r *http.Request, httpStatusCode int, details string) string {
	uinfoStr := ""

	//if sess, err := session.GetSession(r); err != nil {
	//	uinfoStr = sess.Username + "/" + strconv.Itoa(int(sess.UserID)) + " - "
	//} // todo: add session
	q := r.URL.String()
	if qun, e := url.QueryUnescape(q); e == nil {
		q = qun
	}

	srcFileInfo := ""
	if pc, file, line, ok := runtime.Caller(2); ok {
		fnameElems := strings.Split(file, "/")
		funcNameElems := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		srcFileInfo = fmt.Sprintf("[%s:%d %s]", strings.Join(fnameElems[len(fnameElems)-3:], "/"),
			line, funcNameElems[len(funcNameElems)-1])
	}

	return fmt.Sprintf("%s - %d - %s%s - %s",
		details, httpStatusCode, uinfoStr, q, srcFileInfo)
}

type JSON map[string]interface{}

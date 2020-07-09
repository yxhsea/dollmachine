package ff_http

import (
	"io/ioutil"
	"strings"
	"net/http"
	log "github.com/sirupsen/logrus"
)

type Context struct {}

func NewContext() *Context{
	return &Context{}
}

func (ctx *Context) PostArgs(r *http.Request) map[string]string{
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Failure to read body data. Error : %s", err.Error())
	}
	strMap := make(map[string]string)
	strArr := strings.Split(string(body), "&")
	for _, v := range strArr {
		vArr := strings.Split(v, "=")
		strMap[vArr[0]] = vArr[len(vArr)-1]
	}
	return strMap
}

func (ctx *Context) QueryArgs(r *http.Request) map[string]string{
	strMap := r.URL.Query()
	queryMap := make(map[string]string)
	for k, v := range strMap {
		queryMap[k] = v[0]
	}
	return queryMap
}
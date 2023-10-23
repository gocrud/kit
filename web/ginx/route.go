package ginx

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type HandleFunc func(*gin.Context) Result

func Handle(h HandleFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ret := h(c)
		raw := ret.(*result)
		for _, f := range filter.handles {
			f(raw)
		}
		if raw == nil {
			c.Status(204)
			return
		}
		c.JSON(raw.getStatus(), raw.getBody())
	}
}

type FilterResult interface {
	Body() any
	Status() int
	SetBody(value any)
	SetStatus(code int)
	GetError() error
}

type FilterHandle func(FilterResult)

type Filter struct {
	mutex   sync.Mutex
	handles []FilterHandle
}

func (f *Filter) setHandle(handles ...FilterHandle) {
	f.mutex.Lock()
	f.handles = append(f.handles, handles...)
	f.mutex.Unlock()
}

var filter = &Filter{}

func SetFilter(handles ...FilterHandle) {
	filter.setHandle(handles...)
}

package test

import(
	"github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/render"
)
// to be deleted
func Test_Render_JSON(t *testing.T) {
	m := martini.Classic()
	m.Get("/foobar", func(r render.Render) {
		
	}
}


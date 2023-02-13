package web

import (
	"net/http"
	"strings"
	"text/template"
)

type HandlerFunc func(*Context)

type (
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}

	Engine struct {
		*RouterGroup
		router        *router
		groups        []*RouterGroup     // store all groups
		htmlTemplates *template.Template // for html render
		funcMap       template.FuncMap   // for html render
	}
)

func New()*Engine{
	engine:=&Engine{router:newRouter()}
	engine.RouterGroup=&RouterGroup{engine:engine}
	engine.groups=[]*RouterGroup{engine.RouterGroup}

	return engine
}

func Default()*Engine{
	engine:=New()
	engine.Use(Logger(),Recovery())

	return engine
}

func (group *RouterGroup)Group(prefix string)*RouterGroup{
	engine:=group.engine
	newGroup:=&RouterGroup{
		prefix:group.prefix+prefix,
		parent:group,
		engine:engine,
	}
	engine.groups=append(engine.groups, newGroup)
	return newGroup
}

func(group *RouterGroup)Use(middlewares ...HandlerFunc){
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouterGroup)addRoute(method string,comp string,handler HandlerFunc){
  pattern:=group.prefix+comp
	group.engine.router.addRoute(method,pattern,handler)
}

func (group *RouterGroup)GET(pattern string,handler HandlerFunc){
	group.addRoute("GET",pattern,handler)
}

func (group *RouterGroup)POST(pattern string,handler HandlerFunc){
	group.addRoute("POST",pattern,handler)
}

func (engine *Engine)SetFuncMap(funcMap template.FuncMap){
	engine.funcMap=funcMap
}

func (engine *Engine)LoadHTMLGlob(pattern string){
	engine.htmlTemplates=template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine)ServeHTTP(writer http.ResponseWriter,request *http.Request){
	var middlewares []HandlerFunc
	for _,group:=range engine.groups{
		if strings.HasPrefix(request.URL.Path,group.prefix){
			middlewares=append(middlewares, group.middlewares...)
		}
	}
	context:=newContext(writer,request)
	context.handlers=middlewares
	context.engine=engine
	engine.router.handle(context)
}
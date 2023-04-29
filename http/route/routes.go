package route

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sabahtalateh/toolbox/http/verbs"
)

type HttpRoute struct {
	path           *routePath
	headerBindings map[string]*paramBinding
	pathBindings   map[string]*paramBinding
	queryBindings  map[string]*paramBinding
	funcDefs       []codeLocation
}

type routePath struct {
	value        string
	normal       string
	placeholders []string
	pathError    error
	defs         []codeLocation
}

type paramBinding struct {
	routeParam string
	funcParam  string
	defs       []codeLocation
}

type codeLocation struct {
	file string
	line int
}

// Routes map[GET] -> map["/some/url"] -> ...
var Routes = map[verbs.Verb]map[string]*HttpRoute{
	verbs.GET:     {},
	verbs.HEAD:    {},
	verbs.POST:    {},
	verbs.PUT:     {},
	verbs.DELETE:  {},
	verbs.CONNECT: {},
	verbs.OPTIONS: {},
	verbs.TRACE:   {},
	verbs.PATCH:   {},
}

func Http(verb verbs.Verb, path string) *HttpRoute {
	var loc codeLocation
	_, loc.file, loc.line, _ = runtime.Caller(1)

	normPath := normalizePath(path)
	route, ok := Routes[verb][normPath]
	if !ok {
		p := routePath{
			value:        path,
			normal:       normPath,
			placeholders: placeholders(path),
		}
		p.pathError = p.validatePath(loc)

		route = &HttpRoute{
			path:           &p,
			headerBindings: map[string]*paramBinding{},
			pathBindings:   map[string]*paramBinding{},
			queryBindings:  map[string]*paramBinding{},
		}
		Routes[verb][normPath] = route
	}

	route.path.defs = append(route.path.defs, loc)

	return route
}

func (h *HttpRoute) Func(_ any) *HttpRoute {
	var loc codeLocation
	_, loc.file, loc.line, _ = runtime.Caller(1)

	h.funcDefs = append(h.funcDefs, loc)

	return h
}

func (h *HttpRoute) BindHeaderParam(headerParam string, funcParam string) *HttpRoute {
	var loc codeLocation
	_, loc.file, loc.line, _ = runtime.Caller(1)

	param, ok := h.headerBindings[headerParam]
	if !ok {
		param = &paramBinding{routeParam: headerParam, funcParam: funcParam}
		h.headerBindings[headerParam] = param
	}

	param.defs = append(param.defs, loc)

	return h
}

func (h *HttpRoute) BindPathParam(pathParam string, funcParam string) *HttpRoute {
	var loc codeLocation
	_, loc.file, loc.line, _ = runtime.Caller(1)

	param, ok := h.pathBindings[pathParam]
	if !ok {
		param = &paramBinding{routeParam: pathParam, funcParam: funcParam}
		h.pathBindings[pathParam] = param
	}

	param.defs = append(param.defs, loc)

	return h
}

func (h *HttpRoute) BindQueryParam(queryParam string, funcParam string) *HttpRoute {
	var loc codeLocation
	_, loc.file, loc.line, _ = runtime.Caller(1)

	param, ok := h.queryBindings[queryParam]
	if !ok {
		param = &paramBinding{routeParam: queryParam, funcParam: funcParam}
		h.queryBindings[queryParam] = param
	}

	param.defs = append(param.defs, loc)

	return h
}

func placeholders(path string) []string {
	var placeholders []string

	open := -1
	for i, x := range path {
		if x == '{' {
			open = i
		}
		if x == '}' {
			if open != -1 {
				placeholders = append(placeholders, path[open+1:i])
				open = -1
			}
		}
	}

	return placeholders
}

func (p *routePath) validatePath(loc codeLocation) error {
	for _, placeholder := range p.placeholders {
		if placeholder == "{}" {
			return makeError(loc, "empty placeholder `{}` not allowed")
		}
	}

	for i := 0; i < len(p.placeholders); i++ {
		for j := i + 1; j < len(p.placeholders); j++ {
			if p.placeholders[i] == p.placeholders[j] {
				return makeError(loc, "placeholder duplication `{%s}` not allowed", p.placeholders[i])
			}
		}
	}

	return nil
}

func normalizePath(path string) string {
	for _, p := range placeholders(path) {
		path = strings.Replace(path, fmt.Sprintf("{%s}", p), "{}", 1)
	}

	return path
}

func makeError(loc codeLocation, message string, args ...string) error {
	var argsAny []any
	for _, arg := range args {
		argsAny = append(argsAny, arg)
	}

	return fmt.Errorf(message+fmt.Sprintf("\n\tat %s:%d", loc.file, loc.line), argsAny...)
}

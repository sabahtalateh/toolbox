package main

//
// import (
// 	"context"
// 	"net/http"
//
// 	"github.com/sabahtalateh/toolbox"
// 	"github.com/sabahtalateh/toolbox/di/inject"
// 	"github.com/sabahtalateh/toolbox/http/request"
// 	"github.com/sabahtalateh/toolbox/http/response"
// 	"github.com/sabahtalateh/toolbox/http/route"
// 	"github.com/sabahtalateh/toolbox/http/verbs"
// )
//
// func main() {
// 	a := toolbox.TB
// 	println(a)
// }
//
// func (s *Srv1) Hello() {
//
// }
//
// type Hello interface {
// 	Hello()
// }
//
// type Conf struct {
// 	V1 string
// 	V2 string
// }
//
// type Conf1 Conf
// type Conf2 Conf
//
// func init() {
// 	route.Http(verbs.GET, "/item/{id1}/{id1}/{id3}").
// 		Func(Handle).
// 		Func(Handle2).
// 		BindPathParam("a", "b").
// 		BindPathParam("a", "A")
//
// 	route.Http(verbs.GET, "/item/{hello}/{world}/{!}")
//
// 	// component.Component[Srv1](NewSrv1).Name("SRV111")
// }
//
// type Srv1 struct {
// 	srv2  *Srv2
// 	srv22 *Srv2
// 	ss    []*Srv2
// 	conf1 string
// }
//
// func NewSrv1(
// 	srv2 *inject.Struct[Srv2],
// 	srv22 *inject.Struct[Srv2],
// 	srvcs *inject.StructList[Srv2],
// 	hello *inject.Impl[Hello],
// 	hellos *inject.ImplList[Hello],
// 	conf1 *inject.Config[Conf1],
// 	conf2 *inject.Config[Conf2],
// ) *Srv1 {
// 	hh := hellos.Get()
// 	for _, h := range hh {
// 		h.Hello()
// 	}
// 	return &Srv1{
// 		srv2:  srv2.Get(),
// 		srv22: srv22.ByName("SRV2222"),
// 		ss:    srvcs.ByName("SRV111", "SRV222"),
// 		conf1: conf1.Get().V1,
// 	}
// }
//
// func Start(srv1 *Srv1, ctx context.Context) {
//
// }
//
// type Srv2 struct {
// }
//
// func NewSrv22() *Srv2 {
// 	return &Srv2{}
// }
//
// func NewSrv2() *Srv2 {
// 	return &Srv2{}
// }
//
// type Item struct {
// 	Id     string
// 	Name   string
// 	Amount int
// }
//
// type Err struct {
// }
//
// func Handle(
// 	r *http.Request,
// 	w http.ResponseWriter,
// 	srv1 inject.Struct[Srv1],
// 	item request.Body[Item],
// ) (
// 	*response.Ok[Item],
// 	*response.NotFound[Err],
// 	*response.InternalServerError[Err],
// ) {
// 	i := Item{}
// 	return response.NewOk(&i), nil, nil
// }
//
// func Handle2(
// 	userId request.Header[string],
// 	id request.PathParam[string],
// 	from request.QueryParam[string],
// 	to request.QueryParam[string],
// 	srv1 inject.Struct[Srv1],
// ) (
// 	*response.Ok[Item],
// 	*response.NotFound[Err],
// 	*response.InternalServerError[Err],
// ) {
// 	i := Item{}
// 	return response.NewOk(&i), nil, nil
// }

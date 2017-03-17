// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tPostController struct {}
var PostController tPostController


func (_ tPostController) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("PostController.Index", args).Url
}

func (_ tPostController) Show(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("PostController.Show", args).Url
}

func (_ tPostController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("PostController.Create", args).Url
}

func (_ tPostController) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("PostController.Update", args).Url
}

func (_ tPostController) Delete(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("PostController.Delete", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}



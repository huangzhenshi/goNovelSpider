package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["novelProject/controllers:MainController"] = append(beego.GlobalControllerRouter["novelProject/controllers:MainController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["novelProject/controllers:MainController"] = append(beego.GlobalControllerRouter["novelProject/controllers:MainController"],
        beego.ControllerComments{
            Method: "Download",
            Router: `/download`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["novelProject/controllers:MainController"] = append(beego.GlobalControllerRouter["novelProject/controllers:MainController"],
        beego.ControllerComments{
            Method: "ReadNovel",
            Router: `/read`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

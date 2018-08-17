package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["airTicket/controllers:AirController"] = append(beego.GlobalControllerRouter["airTicket/controllers:AirController"],
		beego.ControllerComments{
			Method: "AirAdd",
			Router: `/api/airAdd`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:AirController"] = append(beego.GlobalControllerRouter["airTicket/controllers:AirController"],
		beego.ControllerComments{
			Method: "AirDelete",
			Router: `/api/airDelete`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:AirController"] = append(beego.GlobalControllerRouter["airTicket/controllers:AirController"],
		beego.ControllerComments{
			Method: "AirLoad",
			Router: `/api/airQuery`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:AirController"] = append(beego.GlobalControllerRouter["airTicket/controllers:AirController"],
		beego.ControllerComments{
			Method: "AirUpdate",
			Router: `/api/airUpdate`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:CarController"] = append(beego.GlobalControllerRouter["airTicket/controllers:CarController"],
		beego.ControllerComments{
			Method: "CarTicket",
			Router: `/api/AddBuyCar`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:CarController"] = append(beego.GlobalControllerRouter["airTicket/controllers:CarController"],
		beego.ControllerComments{
			Method: "CarDel",
			Router: `/api/DelBuyCar`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:CarController"] = append(beego.GlobalControllerRouter["airTicket/controllers:CarController"],
		beego.ControllerComments{
			Method: "GetAllCar",
			Router: `/api/GetBuyCar`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:CarController"] = append(beego.GlobalControllerRouter["airTicket/controllers:CarController"],
		beego.ControllerComments{
			Method: "GetAllOrder",
			Router: `/api/GetOrder`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:CarController"] = append(beego.GlobalControllerRouter["airTicket/controllers:CarController"],
		beego.ControllerComments{
			Method: "BuyTicket",
			Router: `/api/ResBuyCar`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:MainController"] = append(beego.GlobalControllerRouter["airTicket/controllers:MainController"],
		beego.ControllerComments{
			Method: "GetAllFlight",
			Router: `/api/getFlight`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:MainController"] = append(beego.GlobalControllerRouter["airTicket/controllers:MainController"],
		beego.ControllerComments{
			Method: "AdLogin",
			Router: `/login`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:MainController"] = append(beego.GlobalControllerRouter["airTicket/controllers:MainController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:MainController"] = append(beego.GlobalControllerRouter["airTicket/controllers:MainController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/updateUser`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:MainController"] = append(beego.GlobalControllerRouter["airTicket/controllers:MainController"],
		beego.ControllerComments{
			Method: "UsLogin",
			Router: `/userLogin`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:TicketController"] = append(beego.GlobalControllerRouter["airTicket/controllers:TicketController"],
		beego.ControllerComments{
			Method: "DYTicket",
			Router: `/api/ticketDY`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["airTicket/controllers:TicketController"] = append(beego.GlobalControllerRouter["airTicket/controllers:TicketController"],
		beego.ControllerComments{
			Method: "QueryTicket",
			Router: `/api/ticketQuery`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

}

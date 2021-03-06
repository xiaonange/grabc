package route

import (
	. "github.com/codyi/grabc/views/layout"
)

type Index struct {
	BaseTemplate
}

func (this *Index) Html() string {
	html := `
    <div class="box box-info">
        <div class="box-body">
            <table class="table route_warp">
                <tbody>
                    <tr>
                        <td style="width: 40%">
                            <select multiple="" size="20" id="select_route_not_add">
                            </select>
                        </td>
                        <td>
                            <div>
                                <button id="btn_add_route" class="btn btn-primary">>>添加</button>
                            </div>
                            <div style="margin-top: 15px">
                                <button id="btn_remove_route" class="btn btn-danger"><<删除</button>
                            </div>
                        </td>
                        <td style="width: 40%">
                            <select multiple="" size="20" id="select_route_add">
                            </select>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
	<script>
        var add_routes = new Array();    //已添加路由
        var not_add_routes = new Array(); //未添加路由

		{{range $index,$route := .notAddRoutes}}
		not_add_routes.push("{{$route}}")
        {{end}}
        {{range $index,$route := .addRoutes}}
        add_routes.push("{{$route}}")
		{{end}}		
		
		$(function(){
			$.showSelectOption("#select_route_add", add_routes);
			$.showSelectOption("#select_route_not_add", not_add_routes);

			//添加路由
		    $.addRoute();
		    //删除路由
		    $.removeRoute();
		});
		
		//添加路由
		$.addRoute = function () {
		    $("#btn_add_route").click(function () {
		        var select_routes = $("#select_route_not_add").val();
		
		        if (select_routes.length > 0) {
		            $("#btn_add_route").attr("disabled","disabled");
		
		            $(select_routes).each(function (index, value) {
		                $.ajax({
		                    type:"post",
		                    url:"/route/ajaxadd",
		                    data:{route:value},
		                    dataType:"json",
		                    async:false,
		                    success:function (response) {
		                        if (response.Code == 200) {
		                            add_routes.push(value);
		                            not_add_routes = $.removeItem(value, not_add_routes);
		                            $("#select_route_not_add option[value='"+value+"']").remove();
		                            $.showSelectOption("#select_route_add", add_routes);
		                        }
		                    }
		                });
		            });
		
		            $("#btn_add_route").removeAttr("disabled");
		        }
		    });
		};
		
		//删除路由
		$.removeRoute = function () {
		    $("#btn_remove_route").click(function () {
		        var select_routes = $("#select_route_add").val();
		
		        if (select_routes.length > 0) {
		            $("#btn_remove_route").attr("disabled","disabled");
		
		            $(select_routes).each(function (index, value) {
		                $.ajax({
		                    type:"post",
		                    url:"/route/ajaxremove",
		                    data:{route:value},
		                    dataType:"json",
		                    async:false,
		                    success:function (response) {
		                        if (response.Code == 200) {
		                            not_add_routes.push(value)
		                            add_routes = $.removeItem(value, add_routes);
		                            $("#select_route_add option[value='"+value+"']").remove();
		                            $.showSelectOption("#select_route_not_add", not_add_routes);
		                        }
		                    }
		                });
		            });
		
		            $("#btn_remove_route").removeAttr("disabled");
		        }
		    });
		};
    </script>
`
	return this.DealHtml(html)
}

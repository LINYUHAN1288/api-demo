package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	//"testing"
)

type HttplibController struct {
	beego.Controller
}
// URLMapping...
func (c *HttplibController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// Get...
// 获取原接口返回的数据
func (c *HttplibController) Get() {
	type1 := c.Ctx.Input.Param(":type") // 接收用户数据
	req := httplib.Get("http://api.tkjidi.com/getGoodsLink?appkey=d46dc32ae787795284fe0a2d8b427248"+"&type="+type1)
	str, err := req.String()
 
    if err != nil {
        panic(err)
    }
	c.Ctx.WriteString(str)
}

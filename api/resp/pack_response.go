package resp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type None struct {
	Readme  string `json:"_README!!!!" example:"如果您看到这行文字，表示这里的示例 data 字段实际是 null，这里的内容是为上层封装做注释。"`
	Code    string `json:"code" example:"0-成功，1-失败"`
	Success string `json:"success" example:"成功会显示成 true，失败会被设置为 false"`
	Msg     string `json:"msg" example:"msg 字段显示相应错误的信息，如数据库查询失败、参数校验失败等"`
	Data    string `json:"data" example:"实际是 null"`
}

type Pack[D any] struct {
	context          *gin.Context
	httpResponseCode int

	Code    uint   `json:"code" example:"0"`
	Success bool   `json:"success" example:"true"`
	Msg     string `json:"msg" example:"msg"`
	Data    D      `json:"data"`
}

func SuccessPack[D any](c *gin.Context) *Pack[D] {
	return &Pack[D]{
		context:          c,
		httpResponseCode: 200,

		Code:    0,
		Msg:     "OK",
		Success: true,
	}
}

func ErrorPack[D any](c *gin.Context) *Pack[D] {
	return &Pack[D]{
		context:          c,
		httpResponseCode: 400,

		Code:    1,
		Msg:     "ERROR!",
		Success: false,
	}
}

func (p *Pack[D]) WithMessage(message string, a ...interface{}) *Pack[D] {
	if len(a) == 0 {
		p.Msg = message
	} else {
		p.Msg = fmt.Sprintf(message, a)
	}
	return p
}

func (p *Pack[D]) WithData(data D) *Pack[D] {
	p.Data = data
	return p
}

func (p *Pack[D]) WithErrorCode(code uint) *Pack[D] {
	p.Code = code
	return p
}

func (p *Pack[D]) WithHttpResponseCode(responseCode int) *Pack[D] {
	p.httpResponseCode = responseCode
	return p
}

func (p *Pack[D]) Responds() {
	p.context.Header("x-sg-err-code", strconv.Itoa(int(p.Code)))
	errDesc := ""
	if !p.Success {
		errDesc = p.Msg
	}
	p.context.Header("x-sg-err-desc", errDesc)
	p.context.Header("x-sg-biz-code", strconv.Itoa(int(p.Code)))
	p.context.Header("x-sg-biz-desc", p.Msg)
	p.context.JSON(p.httpResponseCode, p)
}

func ErrorMessageStatus(c *gin.Context, message string, status int) {
	ErrorPack[any](c).WithMessage(message).WithHttpResponseCode(status).Responds()
}

func Error(c *gin.Context, status int, message string, err error) {
	logUUID := ResponseLogging(err)
	msg := fmt.Sprintf("[%v]%v", logUUID, message)
	ErrorPack[any](c).WithMessage(msg).WithHttpResponseCode(status).Responds()
}

func SuccessMessage(c *gin.Context, message string) {
	SuccessPack[any](c).WithMessage(message).Responds()
}

func SuccessWithData[D any](c *gin.Context, data D) {
	SuccessPack[D](c).WithData(data).Responds()
}

func SuccessMessageData[D any](c *gin.Context, message string, data D) {
	SuccessPack[D](c).WithMessage(message).WithData(data).Responds()
}

func SuccessMessageDataCode[D any](c *gin.Context, message string, data D, httpResponseCode int) {
	SuccessPack[D](c).WithMessage(message).WithData(data).WithHttpResponseCode(httpResponseCode).Responds()
}

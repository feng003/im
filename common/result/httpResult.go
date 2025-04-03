package result

import (
	"fmt"
	"net/http"
	"strings"
	"im/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

// http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()

		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				} else {
					errmsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusBadRequest, Error(errcode, errmsg))
	}
}

// 授权的http方法
func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【GATEWAY-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusUnauthorized, Error(errcode, errmsg))
	}
}

func WechatHttpResult(r *http.Request, w http.ResponseWriter, resp string, err error) {

	if err == nil {
		//httpx.WriteJson(w, http.StatusOK, resp)
		fmt.Fprintf(w, resp)
	} else {
		r := ErrorAlipay(err.Error())
		httpx.WriteJson(w, http.StatusBadRequest, r)
	}
}

// spi
func AlipayHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		httpx.WriteJson(w, http.StatusOK, resp)
	} else {
		r := ErrorAlipay(err.Error())
		httpx.WriteJson(w, http.StatusBadRequest, r)
	}
}

// zft
func ZftHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		_, err = fmt.Fprintf(w, "success")
	} else {
		_, err = fmt.Fprintf(w, "fail")
	}
}

func AuthResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	page := `<!DOCTYPE html>
				<html lang="en">
				<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>授权成功</title>
				<style>
					body {
						margin: 0;
						padding: 0;
						display: flex;
						justify-content: center;
						align-items: center;
						height: 100vh;
						font-family: Arial, sans-serif;
					}
					.message {
						text-align: center;
						padding: 20px;
						font-size: 26px;
						border-radius: 5px;
					}
				</style>
				</head>
				<body>
					<div class="message">
						授权成功！
					</div>
				</body>
				</html>
				`

	_, err = fmt.Fprintf(w, page)
}

func QrcodePage(r *http.Request, w http.ResponseWriter, resp interface{}, qrcodeUrl string, err error) {
	fmt.Printf("qrcodeUrl %s and resp %+v \n", qrcodeUrl, resp)
	qrcodePage := strings.Replace(`<!DOCTYPE html>
                <html lang="en">
                  <head>
                   <meta charset="UTF-8" />
                   <meta name="viewport" content="width=device-width, initial-scale=1.0" />
                   <title>二维码</title>
                   <style>
                     .box {
                      width: 100%;
                      height: 98vh; 
                      display: flex;
                      justify-content: center; 
                      align-items: center;
                     }
                     img {
                      width: 360px;
                      height: 450px;
                     }
                   </style>
                  </head>
                  <body>
                   <div class="box">
                     <img src="%s" alt="" />
                   </div>
                  </body>
                </html>`, "%s", qrcodeUrl, 1)
	_, err = fmt.Fprintf(w, qrcodePage)
}

// http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.REUQEST_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.REUQEST_PARAM_ERROR, errMsg))
}

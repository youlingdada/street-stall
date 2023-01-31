package router

var accessUrls map[string]string

func init() {
	accessUrls = make(map[string]string, 10)

	accessUrls["/user/login"] = "/user/login"
	accessUrls["/user/register"] = "/user/register"
	accessUrls["/user/sendRegCode"] = "/user/sendRegCode"
	accessUrls["/user/updatePwd"] = "/user/updatePwd"
	accessUrls["/user/sendModifyCode"] = "/user/sendModifyCode"

	accessUrls["/swagger/doc.json"] = "/swagger/doc.json"

	accessUrls["/utils/encodeRsa"] = "/utils/encodeRsa"
	accessUrls["/utils/decodeRsa"] = "/utils/decodeRsa"
	accessUrls["/utils/getPublicRsaKey"] = "/utils/getPublicRsaKey"

}

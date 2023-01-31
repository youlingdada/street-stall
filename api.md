## 摊位后端api


#### 全局参数
* baseUrl
```shell
http://youlingdada.top:8080
```
* 返回参数格式(eg /user/login)

```json
{
    "code":"200", // 状态码
    "message": "login seccessful", // 提示信息，调试信息
    "data":"token" // 返回的数据
}
```
* res code 返回数据的状态码
```javascript
NO_AUTH = 403  // 用户未登录
REQ_SUCCESS = 200 // 相关的操作成功 到达预期的效果
REQ_FAIL = 400 // 相关的操作失败 未达到预期效果
```


#### 用户模块

##### 登录
* url
```shell
/user/login
```
* http method
```shell
POST
```
* params
```json
{
    "username":"yxd", // 用户名
    "password":"rsa encry password", // 使用rsa 公钥加密后的密文 
}
```

* res
```json
{
    "code":"200", // 200 -> login success 400 -> login fail
    "message": "login seccessful",
    "data":"token" // 用户jwt token 数据
}
```

##### 注册
* url
```shell
/url/register
```
* http method
```
POST
```
* params
```json
{
    "username":"yxd",
    "password":"rsa encry password",
    "email":"youlingdada@163.com", // 邮箱
    "code":"email verify code" // 邮箱验证码 通过/user/getEmailcode 获取
}
```

##### 获取邮箱验证码


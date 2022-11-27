package models

import (
	"k8s.io/api/networking/v1beta1"
)

//选项，显示用（列表)
type IngressOptions struct {
	IsCros bool
	IsRewrite bool
	IsAuth bool
	IsRatelimit bool
}
// 列表用
type IngressModel struct {
	Name string
	NameSpace string
	CreateTime string
	Host string
	Options IngressOptions

}
type GenAuth struct {
	Ns string `json:"ns"`
	Sname string `json:"sname"`  // secretname
	Uname string `json:"Uname"` // 用户名
	Upwd string `json:"upwd"` //用户密码
}
///下面的模型是提交（新增或修改用)
 // path 配置
type IngressPath struct {
	Path string  `json:"path"`
	SvcName string `json:"svc_name"`
	Port string `json:"port"`

}

// 规则
type  IngressRules struct {
	Host string `json:"host"`
	Paths []*IngressPath `json:"paths"`
}

//提交Ingress 对象
type  IngressPost struct{
	Name string
	Namespace string
	Rules []*IngressRules
	Annotations map[string]string //标签  之前是字符串，现在改成map.对应前端是 对象形式
	TLS []v1beta1.IngressTLS
}


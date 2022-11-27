package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8sapi/src/helpers"
	"sigs.k8s.io/yaml"

	"k8sapi/src/models"
	"k8sapi/src/services"
)

type IngressCtl struct{
	IngressMap *services.IngressMapStruct `inject:"-"`
	IngressService *services.IngressService `inject:"-"`
	Client *kubernetes.Clientset  `inject:"-"`
}
func NewIngressCtl() *IngressCtl{
  return &IngressCtl{}
}
func(*IngressCtl)  Name() string{
	 return "IngressCtl"
}

func(this *IngressCtl) PostIngress(c *gin.Context) goft.Json{
	postModel:=&models.IngressPost{}
	goft.Error(c.BindJSON(postModel))
	if c.Query("update")!=""{  //如果有query参数，且有值。（值随便是什么) 代表是更新
		goft.Error(this.IngressService.PostIngress(postModel,true ))
	}else{

		goft.Error(this.IngressService.PostIngress(postModel,false ))
	}


	return gin.H{
		"code":20000,
		"data":postModel,
	}
}
//DELETE /ingress?ns=xx&name=xx
func(this *IngressCtl) RmIngress(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	goft.Error(this.Client.NetworkingV1beta1().Ingresses(ns).
		Delete(c,name,v1.DeleteOptions{}))
	return gin.H{
		"code":20000,
		"data":"OK",
	}
}
//快捷生成 basic auth 密文
func(this *IngressCtl) GenAuthFileOrMap(c *gin.Context) goft.Json{
	auth:=&models.GenAuth{}
	goft.Error(c.ShouldBindJSON(auth))
	t:=c.DefaultQuery("t","1") // 1是auth-file  2是 auth-map
	//为了简单，目前只支持单个用户名和密码生成 ，如果要扩展 自行修改secret内容
	secret:=corev1.Secret{}
	secret.Name=auth.Sname
	secret.Namespace=auth.Ns

	if t=="1"{
		secret.Data= map[string][]byte{
			"auth":[]byte(auth.Uname+":"+helpers.HashApr1(auth.Upwd)),
		}
	}else{
		secret.Data= map[string][]byte{
			auth.Uname:[]byte(helpers.HashApr1(auth.Upwd)),
		}
	}
	//调用k8s client 创建出 密文
	_,err:=this.Client.CoreV1().Secrets(auth.Ns).Create(c,&secret,v1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code":20000,
		"data":"success",

	}
}
func(this *IngressCtl) ListAll(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	return gin.H{
		"code":20000,
		 "data":this.IngressService.ListIngress(ns), //暂时 不分页

	}
}
//吧ingress对象转化为yaml
func(this *IngressCtl) GetIngressForYaml(c *gin.Context) goft.Json{

	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	ingress,err:=this.Client.NetworkingV1beta1().Ingresses(ns).Get(c,name,v1.GetOptions{})

	goft.Error(err)

	b,err:=yaml.Marshal(ingress)
	goft.Error(err)
	return gin.H{
		"code":20000,
		"data":string(b), //暂时 不分页
	}
}
//加载详细的ingress
func(this *IngressCtl) LoadIngressDetail(c *gin.Context) goft.Json{
	ns:=c.Param("ns")
	name:=c.Param("name")
	ingress:=this.IngressMap.Get(ns,name)// 原生对象
	if ingress==nil{
		panic(fmt.Sprintf("no such ingress %s/%s",ns,name))
	}
	return gin.H{
		"code":20000,
		"data":ingress,
	}

}

func(this *IngressCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/ingress",this.ListAll)
	goft.Handle("GET","/ingress/:ns/:name",this.LoadIngressDetail)



	goft.Handle("DELETE","/ingress",this.RmIngress)
	goft.Handle("POST","/ingress",this.PostIngress)

	goft.Handle("POST","/ingressauth",this.GenAuthFileOrMap) //快捷 创建 basic 密文
	goft.Handle("GET","/ingressyaml",
		this.GetIngressForYaml)
}

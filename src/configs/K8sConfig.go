package configs

import (
	tektonVersiond "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"k8sapi/pkg/kubefed"
	"k8sapi/pkg/rbac"
	"k8sapi/pkg/tekton"
	"k8sapi/src/models"
	"k8sapi/src/services"
	"log"
)

type K8sConfig struct {
	DepHandler                *services.DepHandler           `inject:"-"`
	RsHandler                 *services.RsHandler            `inject:"-"`
	PodHandler                *services.PodHandler           `inject:"-"`
	NsHandler                 *services.NsHandler            `inject:"-"`
	EventHandler              *services.EventHandler         `inject:"-"`
	IngressHandler            *services.IngressHandler       `inject:"-"`
	ServiceHandler            *services.ServiceHandler       `inject:"-"`
	SecretHandler             *services.SecretHandler        `inject:"-"`
	ConfigMapHandler          *services.ConfigMapHandler     `inject:"-"`
	NodeHandler               *services.NodeMapHandler       `inject:"-"`
	RoleHander                *rbac.RoleHander               `inject:"-"`
	RoleBindingHander         *rbac.RoleBindingHander        `inject:"-"`
	SaHandler                 *rbac.SaHander                 `inject:"-"`
	ClusterRoleHandler        *rbac.ClusterRoleHandler       `inject:"-"`
	ClusterRoleBindingHandler *rbac.ClusterRoleBindingHander `inject:"-"`

	TaskHandler        *tekton.TaskHandler        `inject:"-"`
	TaskRunHandler     *tekton.TaskRunHandler     `inject:"-"`
	PipelineHandler    *tekton.PipelineHandler    `inject:"-"`
	PipelineRunHandler *tekton.PipelineRunHandler `inject:"-"`

	FdeployHandler *kubefed.FdeployHandler `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

//初始化 系统 配置
func (*K8sConfig) InitSysConfig() *models.SysConfig {
	b, err := ioutil.ReadFile("app.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config := &models.SysConfig{}
	err = yaml.Unmarshal(b, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
func (*K8sConfig) K8sRestConfig() *rest.Config {
	//config, err := clientcmd.BuildConfigFromFlags("","config" )

	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	config.Insecure = true

	if err != nil {
		log.Fatal(err)
	}
	return config
}

//初始化client-go客户端
func (this *K8sConfig) InitClient() *kubernetes.Clientset {
	//config:=&rest.Config{
	//	Host:"http://124.70.204.12:8009",
	//}
	//rest.Config{
	//	TLSClientConfig:
	//}
	c, err := kubernetes.NewForConfig(this.K8sRestConfig())
	c.RESTClient().GetRateLimiter()
	if err != nil {
		log.Fatal(err)
	}
	return c
}

//初始化tekton客户端
func (this *K8sConfig) InitTektonClient() *tektonVersiond.Clientset {
	c, err := tektonVersiond.NewForConfig(this.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// metric客户端
func (this *K8sConfig) InitMetricClient() *versioned.Clientset {
	c, err := versioned.NewForConfig(this.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return c
}

//初始化Informer
func (this *K8sConfig) InitInformer() informers.SharedInformerFactory {
	fact := informers.NewSharedInformerFactory(this.InitClient(), 0)

	depInformer := fact.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(this.DepHandler)

	rsInformer := fact.Apps().V1().ReplicaSets()
	rsInformer.Informer().AddEventHandler(this.RsHandler)

	podInformer := fact.Core().V1().Pods() //监听pod
	podInformer.Informer().AddEventHandler(this.PodHandler)

	serviceInformer := fact.Core().V1().Services() //监听service
	serviceInformer.Informer().AddEventHandler(this.ServiceHandler)

	nsInformer := fact.Core().V1().Namespaces() //监听namespace
	nsInformer.Informer().AddEventHandler(this.NsHandler)

	eventInformer := fact.Core().V1().Events() //监听event
	eventInformer.Informer().AddEventHandler(this.EventHandler)

	IngressInformer := fact.Networking().V1beta1().Ingresses() //监听Ingress
	IngressInformer.Informer().AddEventHandler(this.IngressHandler)

	SecretInformer := fact.Core().V1().Secrets() //监听Secret
	SecretInformer.Informer().AddEventHandler(this.SecretHandler)

	ConfigMapInformer := fact.Core().V1().ConfigMaps() //监听Configmap
	ConfigMapInformer.Informer().AddEventHandler(this.ConfigMapHandler)

	NodeInformer := fact.Core().V1().Nodes()
	NodeInformer.Informer().AddEventHandler(this.NodeHandler)

	RolesInformer := fact.Rbac().V1().Roles()
	RolesInformer.Informer().AddEventHandler(this.RoleHander)

	RolesBindingInformer := fact.Rbac().V1().RoleBindings()
	RolesBindingInformer.Informer().AddEventHandler(this.RoleBindingHander)

	SaInformer := fact.Core().V1().ServiceAccounts()
	SaInformer.Informer().AddEventHandler(this.SaHandler)

	fact.Rbac().V1().ClusterRoles().Informer().AddEventHandler(this.ClusterRoleHandler)
	fact.Rbac().V1().ClusterRoleBindings().Informer().AddEventHandler(this.ClusterRoleBindingHandler)
	fact.Start(wait.NeverStop)
	return fact
}

//动态informers处理
var taskResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "tasks", Version: "v1beta1"}
var pipelineResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "pipelines", Version: "v1beta1"}
var piperunlineResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "pipelineruns", Version: "v1beta1"}
var taskrunResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "taskruns", Version: "v1beta1"}

///////////////////kubefed相关
var fdeploy = schema.GroupVersionResource{Group: "types.kubefed.io", Resource: "federateddeployments", Version: "v1beta1"}

func (this *K8sConfig) InitGenericInformer() dynamicinformer.DynamicSharedInformerFactory {
	client, err := dynamic.NewForConfig(this.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	di := dynamicinformer.NewDynamicSharedInformerFactory(client, 0)

	taskInformer := di.ForResource(taskResource)
	taskInformer.Informer().AddEventHandler(this.TaskHandler)

	di.ForResource(taskrunResource).Informer().AddEventHandler(this.TaskRunHandler)
	di.ForResource(pipelineResource).Informer().AddEventHandler(this.PipelineHandler)
	di.ForResource(piperunlineResource).Informer().AddEventHandler(this.PipelineRunHandler)

	di.ForResource(fdeploy).Informer().AddEventHandler(this.FdeployHandler)

	di.Start(wait.NeverStop)
	return di
}

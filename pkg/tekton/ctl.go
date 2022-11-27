package tekton

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	tektonVersiond "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/models"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
)
//程序员在囧途(www.jtthink.com)咨询群：98514334
type TektonCtl struct {
	TektonService *TektonService `inject:"-"`
	Client *tektonVersiond.Clientset `inject:"-"`
}
func NewTektonCtl() *TektonCtl {
	return &TektonCtl{}
}
/////task 相关
//本课程来自 程序员在囧途(www.jtthink.com)咨询群：98514334
func(this *TektonCtl) LoadTask(ctx *gin.Context) goft.Json{

	corev1.Pod{}

	ns:=ctx.Param("ns")
	name:=ctx.Param("name")

//	v1beta12.EventListener{}
//gitlab.Interceptor{}


	task,err:=this.Client.TektonV1beta1().Tasks(ns).Get(ctx,name,v1.GetOptions{})
	goft.Error(err)
	return gin.H{
		"code":20000,
		"data":task,
	}
}
func(this *TektonCtl) TaskList(ctx *gin.Context) goft.Json{
	ns:=ctx.DefaultQuery("ns","default")
	  return gin.H{
		"code":20000,
		"data":this.TektonService.LoadTask(ns),
	}
}
func(this *TektonCtl) RmTask(ctx *gin.Context) goft.Json{
	ns:=ctx.Param("ns")
	name:=ctx.Param("name")
	goft.Error(this.Client.TektonV1beta1().Tasks(ns).Delete(ctx,name,v1.DeleteOptions{}))
	return gin.H{
		"code":20000,
		"data":"success",
	}
}
//80  k8sconfig 第80 行 就看到 client 如何创建
func(this *TektonCtl) SaveTask(ctx *gin.Context) goft.Json{
	task:=&v1beta1.Task{}
	goft.Error(ctx.ShouldBindJSON(task))
	update:=ctx.Query("update")
	if update!=""{ //代表是修改
		_,err:=this.Client.TektonV1beta1().Tasks(task.Namespace).Update(ctx,task,v1.UpdateOptions{})
		goft.Error(err)
	}else{
		_,err:=this.Client.TektonV1beta1().Tasks(task.Namespace).Create(ctx,task,v1.CreateOptions{})
		goft.Error(err)
	}

	return gin.H{
		"code":20000,
		"data":"success",
	}
}


// taskrun 相关
//本课程来自 程序员在囧途(www.jtthink.com)咨询群：98514334
func(this *TektonCtl) LoadTaskRun(ctx *gin.Context) goft.Json{
	ns:=ctx.Param("ns")
	name:=ctx.Param("name")
	task,err:=this.Client.TektonV1beta1().TaskRuns(ns).Get(ctx,name,v1.GetOptions{})
	goft.Error(err)
	return gin.H{
		"code":20000,
		"data":task,
	}
}
func(this *TektonCtl) TaskRunList(ctx *gin.Context) goft.Json{
	ns:=ctx.DefaultQuery("ns","default")
	return gin.H{
		"code":20000,
		"data":this.TektonService.LoadTaskRun(ns),
	}
}
func(this *TektonCtl) RmTaskRun(ctx *gin.Context) goft.Json{
	ns:=ctx.Param("ns")
	name:=ctx.Param("name")
	goft.Error(this.Client.TektonV1beta1().TaskRuns(ns).Delete(ctx,name,v1.DeleteOptions{}))
	return gin.H{
		"code":20000,
		"data":"success",
	}
}
//80  k8sconfig 第80 行 就看到 client 如何创建
func(this *TektonCtl) SaveTaskRun(ctx *gin.Context) goft.Json{
	task:=&v1beta1.TaskRun{}
	goft.Error(ctx.ShouldBindJSON(task))
	update:=ctx.Query("update")
	if update!=""{ //代表是修改
		_,err:=this.Client.TektonV1beta1().TaskRuns(task.Namespace).Update(ctx,task,v1.UpdateOptions{})
		goft.Error(err)
	}else{
		_,err:=this.Client.TektonV1beta1().TaskRuns(task.Namespace).Create(ctx,task,v1.CreateOptions{})
		goft.Error(err)
	}

	return gin.H{
		"code":20000,
		"data":"success",
	}
}


///pipeline相关
//本课程来自 程序员在囧途(www.jtthink.com)咨询群：98514334
func(this *TektonCtl)  PipelineList(ctx *gin.Context) goft.Json{
	//v1beta1.Pipeline{}
	ns:=ctx.DefaultQuery("ns","default")
	return gin.H{
		"code":20000,
		"data":this.TektonService.LoadPipeline(ns),
	}

}
func(this *TektonCtl) RmPipeline(ctx *gin.Context) goft.Json{
	ns:=ctx.Param("ns")
	name:=ctx.Param("name")
	goft.Error(this.Client.TektonV1beta1().Pipelines(ns).Delete(ctx,name,v1.DeleteOptions{}))
	return gin.H{
		"code":20000,
		"data":"success",
	}
}
func(this *TektonCtl) SavePipeline(ctx *gin.Context) goft.Json{
	pipeline:=&v1beta1.Pipeline{}
	goft.Error(ctx.ShouldBindJSON(pipeline))
	update:=ctx.Query("update")
	if update!=""{ //代表是修改
		_,err:=this.Client.TektonV1beta1().Pipelines(pipeline.Namespace).Update(ctx,pipeline,v1.UpdateOptions{})
		goft.Error(err)
	}else{
		_,err:=this.Client.TektonV1beta1().Pipelines(pipeline.Namespace).Create(ctx,pipeline,v1.CreateOptions{})
		goft.Error(err)
	}

	return gin.H{
		"code":20000,
		"data":"success",
	}
}
func(this *TektonCtl) LoadPipeline(ctx *gin.Context) goft.Json{

	ns:=ctx.Param("ns")
	name:=ctx.Param("name")
	task,err:=this.Client.TektonV1beta1().Pipelines(ns).Get(ctx,name,v1.GetOptions{})
	goft.Error(err)
	return gin.H{
		"code":20000,
		"data":task,
	}
}



//pipelinerun 相关
//本课程来自 程序员在囧途(www.jtthink.com)咨询群：98514334
func(this *TektonCtl)  PipelineRunList(ctx *gin.Context) goft.Json{

	ns:=ctx.DefaultQuery("ns","default")
	list:=this.TektonService.LoadPipelineRun(ns)
	fmt.Println(list[0].Status.PipelineResults)

	return gin.H{
		"code":20000,
		"data":this.TektonService.LoadPipelineRun(ns),
	}

}
func(this *TektonCtl) RmPipelineRun(ctx *gin.Context) goft.Json{
	ns:=ctx.Param("ns")
	name:=ctx.Param("name")
	goft.Error(this.Client.TektonV1beta1().PipelineRuns(ns).Delete(ctx,name,v1.DeleteOptions{}))
	return gin.H{
		"code":20000,
		"data":"success",
	}
}
func(this *TektonCtl) SavePipelineRun(ctx *gin.Context) goft.Json{
	pipeline:=&v1beta1.PipelineRun{}

	goft.Error(ctx.ShouldBindJSON(pipeline))
	update:=ctx.Query("update")
	if update!=""{ //代表是修改
		_,err:=this.Client.TektonV1beta1().PipelineRuns(pipeline.Namespace).Update(ctx,pipeline,v1.UpdateOptions{})
		goft.Error(err)
	}else{
		_,err:=this.Client.TektonV1beta1().PipelineRuns(pipeline.Namespace).Create(ctx,pipeline,v1.CreateOptions{})
		goft.Error(err)
	}

	return gin.H{
		"code":20000,
		"data":"success",
	}
}
func(this *TektonCtl) LoadPipelineRun(ctx *gin.Context) goft.Json{

	ns:=ctx.Param("ns")
	name:=ctx.Param("name")
	task,err:=this.Client.TektonV1beta1().PipelineRuns(ns).Get(ctx,name,v1.GetOptions{})
	goft.Error(err)
	return gin.H{
		"code":20000,
		"data":task,
	}
}
func(*TektonCtl) Name() string{
	return "TektonCtl"
}
func(this *TektonCtl)  Build(goft *goft.Goft){

	//task 相关路由
	goft.Handle("GET","/tekton/tasks/:ns/:name",this.LoadTask)
	goft.Handle("GET","/tekton/tasks",this.TaskList)
	goft.Handle("POST","/tekton/tasks",this.SaveTask)
	goft.Handle("DELETE","/tekton/tasks/:ns/:name",this.RmTask)

	//taskrun 相关路由
	goft.Handle("GET","/tekton/taskruns/:ns/:name",this.LoadTaskRun)
	goft.Handle("GET","/tekton/taskruns",this.TaskRunList)
	goft.Handle("POST","/tekton/taskruns",this.SaveTaskRun)
	goft.Handle("DELETE","/tekton/taskruns/:ns/:name",this.RmTaskRun)



	//pipeline相关路由
	goft.Handle("GET","/tekton/pipelines/:ns/:name",this.LoadPipeline)
	goft.Handle("GET","/tekton/pipelines",this.PipelineList)
	goft.Handle("DELETE","/tekton/pipelines/:ns/:name",this.RmPipeline)
	goft.Handle("POST","/tekton/pipelines",this.SavePipeline)

	//pipelinerun 相关路由
	goft.Handle("GET","/tekton/pipelineruns/:ns/:name",this.LoadPipelineRun)
	goft.Handle("GET","/tekton/pipelineruns",this.PipelineRunList)
	goft.Handle("DELETE","/tekton/pipelineruns/:ns/:name",this.RmPipelineRun)
	goft.Handle("POST","/tekton/pipelineruns",this.SavePipelineRun)
}
//程序员在囧途(www.jtthink.com)咨询群：98514334
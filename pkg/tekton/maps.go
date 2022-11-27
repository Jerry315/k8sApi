package tekton

import (
	"fmt"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

	"sort"
	"sync"
)
//Task相关
type V1Task []*v1beta1.Task
func(this V1Task) Len() int{
	return len(this)
}
func(this V1Task) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this V1Task) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}
type TaskMapStruct struct {
	data sync.Map   // [ns string] []*v1beta1.Task
}
func(this *TaskMapStruct) Add(item *v1beta1.Task){
	if list,ok:=this.data.Load(item.Namespace);ok{
		list=append(list.([]*v1beta1.Task),item)
		this.data.Store(item.Namespace,list)
	}else{
		this.data.Store(item.Namespace,[]*v1beta1.Task{item})
	}
}
func(this *TaskMapStruct) Update(item *v1beta1.Task) error {
	if list,ok:=this.data.Load(item.Namespace);ok{
		for i,range_item:=range list.([]*v1beta1.Task){
			if range_item.Name==item.Name{
				list.([]*v1beta1.Task)[i]=item
			}
		}
		return nil
	}
	return fmt.Errorf("Role-%s not found",item.Name)
}
func(this *TaskMapStruct) Delete(svc *v1beta1.Task){
	if list,ok:=this.data.Load(svc.Namespace);ok{
		for i,range_item:=range list.([]*v1beta1.Task){
			if range_item.Name==svc.Name{
				newList:= append(list.([]*v1beta1.Task)[:i], list.([]*v1beta1.Task)[i+1:]...)
				this.data.Store(svc.Namespace,newList)
				break
			}
		}
	}
}
func(this *TaskMapStruct) ListAll(ns string)[]*v1beta1.Task{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*v1beta1.Task)
		sort.Sort(V1Task(newList))//  按时间倒排序
		return newList
	}
	return []*v1beta1.Task{} //返回空列表
}


//TaskRun相关
type V1TaskRun []*v1beta1.TaskRun
func(this V1TaskRun) Len() int{
	return len(this)
}
func(this V1TaskRun) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this V1TaskRun) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}
type TaskRunMapStruct struct {
	data sync.Map   // [ns string] []*v1beta1.Task
}
func(this *TaskRunMapStruct) Add(item *v1beta1.TaskRun){
	if list,ok:=this.data.Load(item.Namespace);ok{
		list=append(list.([]*v1beta1.TaskRun),item)
		this.data.Store(item.Namespace,list)
	}else{
		this.data.Store(item.Namespace,[]*v1beta1.TaskRun{item})
	}
}
func(this *TaskRunMapStruct) Update(item *v1beta1.TaskRun) error {
	if list,ok:=this.data.Load(item.Namespace);ok{
		for i,range_item:=range list.([]*v1beta1.TaskRun){
			if range_item.Name==item.Name{
				list.([]*v1beta1.TaskRun)[i]=item
			}
		}
		return nil
	}
	return fmt.Errorf("TaskRun-%s not found",item.Name)
}
func(this *TaskRunMapStruct) Delete(svc *v1beta1.TaskRun){
	if list,ok:=this.data.Load(svc.Namespace);ok{
		for i,range_item:=range list.([]*v1beta1.TaskRun){
			if range_item.Name==svc.Name{
				newList:= append(list.([]*v1beta1.TaskRun)[:i], list.([]*v1beta1.TaskRun)[i+1:]...)
				this.data.Store(svc.Namespace,newList)
				break
			}
		}
	}
}
func(this *TaskRunMapStruct) ListAll(ns string)[]*v1beta1.TaskRun{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*v1beta1.TaskRun)
		sort.Sort(V1TaskRun(newList))//  按时间倒排序
		return newList
	}
	return []*v1beta1.TaskRun{} //返回空列表
}


//pipeline
type V1Pipeline []*v1beta1.Pipeline
func(this V1Pipeline) Len() int{
	return len(this)
}
func(this V1Pipeline) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this V1Pipeline) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}
type PipelineMapStruct struct {
	data sync.Map   // [ns string] []*v1beta1.pipeline
}
func(this *PipelineMapStruct) Add(item *v1beta1.Pipeline){
	if list,ok:=this.data.Load(item.Namespace);ok{
		list=append(list.([]*v1beta1.Pipeline),item)
		this.data.Store(item.Namespace,list)
	}else{
		this.data.Store(item.Namespace,[]*v1beta1.Pipeline{item})
	}
}
func(this *PipelineMapStruct) Update(item *v1beta1.Pipeline) error {
	if list,ok:=this.data.Load(item.Namespace);ok{
		for i,range_item:=range list.([]*v1beta1.Pipeline){
			if range_item.Name==item.Name{
				list.([]*v1beta1.Pipeline)[i]=item
			}
		}
		return nil
	}
	return fmt.Errorf("Pipeline-%s not found",item.Name)
}
func(this *PipelineMapStruct) Delete(svc *v1beta1.Pipeline){
	if list,ok:=this.data.Load(svc.Namespace);ok{
		for i,range_item:=range list.([]*v1beta1.Pipeline){
			if range_item.Name==svc.Name{
				newList:= append(list.([]*v1beta1.Pipeline)[:i], list.([]*v1beta1.Pipeline)[i+1:]...)
				this.data.Store(svc.Namespace,newList)
				break
			}
		}
	}
}
func(this *PipelineMapStruct) ListAll(ns string)[]*v1beta1.Pipeline{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*v1beta1.Pipeline)
		sort.Sort(V1Pipeline(newList))//  按时间倒排序
		return newList
	}
	return []*v1beta1.Pipeline{} //返回空列表
}


// pipelinerun
//pipeline
type V1PipelineRun []*v1beta1.PipelineRun
func(this V1PipelineRun) Len() int{
	return len(this)
}
func(this V1PipelineRun) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this V1PipelineRun) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}
type PipelineRunMapStruct struct {
	data sync.Map   // [ns string] []*v1beta1.pipeline
}
func(this *PipelineRunMapStruct) Add(item *v1beta1.PipelineRun){
	if list,ok:=this.data.Load(item.Namespace);ok{
		list=append(list.([]*v1beta1.PipelineRun),item)
		this.data.Store(item.Namespace,list)
	}else{
		this.data.Store(item.Namespace,[]*v1beta1.PipelineRun{item})
	}
}
func(this *PipelineRunMapStruct) Update(item *v1beta1.PipelineRun) error {
	if list,ok:=this.data.Load(item.Namespace);ok{
		for i,range_item:=range list.([]*v1beta1.PipelineRun){
			if range_item.Name==item.Name{
				list.([]*v1beta1.PipelineRun)[i]=item
			}
		}
		return nil
	}
	return fmt.Errorf("PipelineRun-%s not found",item.Name)
}
func(this *PipelineRunMapStruct) Delete(svc *v1beta1.PipelineRun){
	if list,ok:=this.data.Load(svc.Namespace);ok{
		for i,range_item:=range list.([]*v1beta1.PipelineRun){
			if range_item.Name==svc.Name{
				newList:= append(list.([]*v1beta1.PipelineRun)[:i], list.([]*v1beta1.PipelineRun)[i+1:]...)
				this.data.Store(svc.Namespace,newList)
				break
			}
		}
	}
}
func(this *PipelineRunMapStruct) ListAll(ns string)[]*v1beta1.PipelineRun{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*v1beta1.PipelineRun)
		sort.Sort(V1PipelineRun(newList))//  按时间倒排序
		return newList
	}
	return []*v1beta1.PipelineRun{} //返回空列表
}
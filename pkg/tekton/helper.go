package tekton

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)


//把对象转化为 task
func ConvertToTask(obj interface{}) *v1beta1.Task  {
	var task *v1beta1.Task=&v1beta1.Task{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(*unstructured.Unstructured).Object, task); err != nil {
		return nil
	}
	return task
}
func ConvertToTaskRun(obj interface{}) *v1beta1.TaskRun  {
	var task *v1beta1.TaskRun=&v1beta1.TaskRun{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(*unstructured.Unstructured).Object, task); err != nil {
		return nil
	}
	return task
}
func ConvertToPipeline(obj interface{}) *v1beta1.Pipeline  {
	  pipeline:=&v1beta1.Pipeline{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(*unstructured.Unstructured).Object, pipeline); err != nil {
		return nil
	}
	return pipeline
}
func ConvertToPipelineRun(obj interface{}) *v1beta1.PipelineRun  {
	pipeline:=&v1beta1.PipelineRun{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(*unstructured.Unstructured).Object, pipeline); err != nil {
		return nil
	}
	return pipeline
}
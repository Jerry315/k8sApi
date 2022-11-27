package tekton

import "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

type TektonService struct {
		TaskMap *TaskMapStruct `inject:"-"`
		TaskRunMap *TaskRunMapStruct `inject:"-"`
		PipelineMap *PipelineMapStruct `inject:"-"`
		PipelineRunMap *PipelineRunMapStruct `inject:"-"`
}
func NewTektonService() *TektonService {
	return &TektonService{}
}
func(this *TektonService) LoadTask(ns string) []*v1beta1.Task{
	return this.TaskMap.ListAll(ns)
}
func(this *TektonService) LoadTaskRun(ns string) []*v1beta1.TaskRun{
	return this.TaskRunMap.ListAll(ns)
}
func(this *TektonService) LoadPipeline(ns string) []*v1beta1.Pipeline{
	return this.PipelineMap.ListAll(ns)
}
func(this *TektonService) LoadPipelineRun(ns string) []*v1beta1.PipelineRun{
	return this.PipelineRunMap.ListAll(ns)
}


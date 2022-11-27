package kubefed

import (
	"fmt"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"sync"
)


type FdeployMapStruct struct {
	data sync.Map   // [ns string] []*v1beta1.fdeploy
}
func(this *FdeployMapStruct) Add(item *unstructured.Unstructured){
	 fmt.Println(item )
}
func(this *FdeployMapStruct) Update(item *v1beta1.Task) error {
	 return nil
}
func(this *FdeployMapStruct) Delete(svc *v1beta1.Task){

}
func(this *FdeployMapStruct) ListAll(ns string)[]*v1beta1.Task{
	 return nil
}


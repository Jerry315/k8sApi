package kubefed

import "fmt"

//Task
type FdeployHandler struct {
	FdeployMap *FdeployMapStruct `inject:"-"`
}
func(this *FdeployHandler) OnAdd(obj interface{}){
	 fmt.Println(obj )
}
func(this *FdeployHandler) OnUpdate(oldObj, newObj interface{}){

}
func(this *FdeployHandler) OnDelete(obj interface{}){

}


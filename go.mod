module k8sapi

go 1.17

require (
	github.com/GehirnInc/crypt v0.0.0-20200316065508-bb7000b8a962
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/websocket v1.4.2
	github.com/shenyisyn/goft-gin v0.5.0
	github.com/tektoncd/pipeline v0.26.0
	github.com/tektoncd/triggers v0.15.0 // indirect
	golang.org/x/crypto v0.0.0-20210415154028-4f45737414dc
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
	k8s.io/kube-aggregator v0.21.2 // indirect
	k8s.io/metrics v0.21.0
	sigs.k8s.io/yaml v1.2.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/client-go/1.5/tools/cache"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/1.5/kubernetes"
	"k8s.io/client-go/1.5/rest"
	"k8s.io/client-go/1.5/pkg/fields"
	"k8s.io/client-go/1.5/pkg/api/v1"
)

func podCreated(obj interface{}) {
	pod := obj.(*v1.Service)
	fmt.Println("Service created: " + pod.ObjectMeta.Name)
}
func podDeleted(obj interface{}) {
	pod := obj.(*v1.Service)
	fmt.Println("Service deleted: " + pod.ObjectMeta.Name)
}
func watchServices(client *rest.RESTClient) cache.Store {
	//Define what we want to look for (Services)
	watchlist := cache.NewListWatchFromClient(client, "services", api.NamespaceAll, fields.Everything())
	resyncPeriod := 30 * time.Minute
	//Setup an informer to call functions when the watchlist changes
	eStore, eController := cache.NewInformer(
		watchlist,
		&v1.Service{},
		resyncPeriod,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    podCreated,
			DeleteFunc: podDeleted,
		},
	)
	//Run the controller as a goroutine
	go eController.Run(wait.NeverStop)
	return eStore
}
func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("ERROR - Could not get the K8s cluster config: [%v]", err)
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("ERROR - Could not get the client for K8s: [%v]", err)
	}

	//Watch for services
	watchServices(clientset.Core().GetRESTClient())
	//Keep alive
	log.Fatal(http.ListenAndServe(":8080", nil))
}
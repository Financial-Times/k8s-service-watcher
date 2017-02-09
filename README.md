Kubernetes services watcher 
=================================

[![Circle CI]()

Sample code for watching Kubernetes events when services are created and removed.

Implementation details
-----------------------------
It uses Kubernetes API for hooking into the K8s events.
I had to introduce vendoring for 2 reasons:
1. To have the version of the go client fixed to 1.5, as at this moment it keeps changing and breaks compatibility
1. I had to vendor the Kubernetes code, as at runtime we received the error: 
```
/k8s-service-watcher flag redefined: log_dir
   panic: /k8s-service-watcher flag redefined: log_dir
   goroutine 1 [running]:
   panic(0x1964040, 0xc82027ac00) 
```
The solution was to remove the glog package from Kubernetes and all the dependencies.

How to Build & Run the binary
-----------------------------

1. Build and test:

        go build
        go test

2. Run

        ./k8s-service-watcher
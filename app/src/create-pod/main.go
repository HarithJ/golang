package main

import (
	"create-pod/kubernetes/auth"
	"create-pod/kubernetes/pod"
)

var json = `
apiVersion: v1
kind: Pod
metadata:
  name: pod-created
spec:
  containers:
    - name: nginx
      image: nginx:1.7.9
      ports:
      - containerPort: 80
`

func main() {
	clientset := auth.GetClientSet()
	pod.CreatePod(clientset, json)
}

package main

import (
	"io/ioutil"
	"log"

	"create-pod/kubernetes/auth"
	"create-pod/kubernetes/pod"
)

func main() {
	clientset := auth.GetClientSet()

	podConfig, err := ioutil.ReadFile("pod.yml")

	if err != nil {
		log.Fatal(err)
	}
	
	pod.CreatePod(clientset, string(podConfig))
}

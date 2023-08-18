package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	fmt.Println(config)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Pods from the cluster")

		for _, pod := range pods.Items {
			fmt.Printf("%s \n", pod.Name)
		}
		deployments, err := clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("listing deployments %s\n", err.Error())
		}
		fmt.Println("\nDeployment from the cluster")

		for _, deployment := range deployments.Items {
			fmt.Printf("%s \n", deployment.Name)
		}

	}
}

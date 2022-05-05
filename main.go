package main

import (
	"context"
	"flag"
	"log"

	"fmt"

	"github.com/gofiber/fiber/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	app := fiber.New()
	//Configs
	kubeconfig := flag.String("kubeconfig", "./kubeconfig.yaml", "")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	//APIs
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/node/:name?", func(c *fiber.Ctx) error {

		no, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})

		if err != nil {
			log.Fatal(err)
		}
		var node []models.myNode
		for _, d := range no.Items {
			fmt.Printf(" %s\n", d.Name)
			node = append(node, models.myNode{d.Name})

		}
		return c.JSON(node)
	})

	app.Listen(":3000")
}

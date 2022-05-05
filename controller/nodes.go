package controller

import (
	"context"
	"log"

	"aziz/k8s/config"
	"aziz/k8s/models"

	"github.com/gofiber/fiber/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodes(c *fiber.Ctx) error {
	no, err := config.Myclientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{LabelSelector: c.Query("label")})
	if err != nil {
		log.Fatal(err)
	}
	var node []models.MyNode
	for _, d := range no.Items {

		node = append(node, models.MyNode{Name: d.Name, Labels: d.Labels})
	}
	return c.JSON(node)
}

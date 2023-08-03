package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func listContainers() error {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return err
	}

	for _, container := range containers {
		fmt.Println("Image:", container.Image, "Labels:", container.Labels,
			"Image ID:", container.ImageID, "Status:", container.Status)
	}
	return nil
}

func listImages() error {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil
	}
	defer cli.Close()

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{All: true})
	if err != nil {
		return err
	}

	for _, image := range images {
		fmt.Println("Images:", image.RepoTags, "Labels:", image.Labels, "Size:", image.Size)
	}
	return nil
}

func main() {
	fmt.Println("List of containers:")
	if err := listContainers(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	fmt.Println("List of images:")
	if err := listImages(); err != nil {
		log.Fatal(err)
	}
}

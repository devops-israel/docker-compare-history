package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
	"os"
	"strconv"
	s "strings"
)

func main() {

	app := cli.NewApp()
	app.Name = "docker-history-compare"
	app.Usage = "Docker image history compare"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "image, i",
			Value: "",
			Usage: "Image ids seperated by ,",
		},
		cli.StringFlag{
			Name:  "endpoint, e",
			Value: "unix:///var/run/docker.sock",
			Usage: "Docker endpoint. example using Boot2Docker tcp://[ip]:[port]",
		},
		cli.IntFlag{
			Name:  "match, m",
			Value: 0,
			Usage: "Number of matching layers",
		},
		cli.StringFlag{
			Name:  "boot2docker, b",
			Usage: "Using Boot2Docker?",
		},
	}
	app.Action = func(c *cli.Context) {
		endpoint := c.String("e")
		match, _ := strconv.ParseInt(c.String("m"), 10, 0)
		var client *docker.Client
		if c.Bool("b") == true {
			path := os.Getenv("DOCKER_CERT_PATH")
			ca := fmt.Sprintf("%s/ca.pem", path)
			cert := fmt.Sprintf("%s/cert.pem", path)
			key := fmt.Sprintf("%s/key.pem", path)
			client, _ = docker.NewTLSClient(endpoint, cert, key, ca)
		} else {
			client, _ = docker.NewClient(endpoint)
		}
		image_ids := s.Split(c.String("i"), ",")
		if !(len(image_ids) == 2) {
			panic("Error! you must have an image id")
		}

		img_history, _ := client.ImageHistory(image_ids[0])
		history_1 := make([]string, len(img_history))
		for _, history := range img_history {
			history_1 = append(history_1, history.ID)
		}

		img_history, _ = client.ImageHistory(image_ids[1])
		history_2 := make([]string, len(img_history))
		for _, history := range img_history {
			history_2 = append(history_2, history.ID)
		}

		fmt.Println(testEq(history_1, history_2, match))
	}

	app.Run(os.Args)
}

func testEq(a, b []string, match int64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			if i > int(match) {
				return true
			} else {
				return false
			}
		}
	}

	return true
}

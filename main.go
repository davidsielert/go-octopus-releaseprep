package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"gitlab.mirv.ninja/rcp/octopus"
	"gopkg.in/yaml.v2"
	_ "gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	dat, err := ioutil.ReadFile("settings.yml")
	if err != nil {
		fmt.Println("Error loading settings.yml")
	}
	var config octopus.Config
	err = yaml.Unmarshal([]byte(dat), &config)
	if err != nil {
		fmt.Println("Error Unmarshalling settings.yml, do you have a syntax error")
	}
	for _, o := range config.Servers {
		//start := time.Now()
		octo := octopus.NewOctopus(&o)
		candidates := process(octo)
		fmt.Println(fmt.Sprintf("==== Release candidates for %s ====", octo.Config.Name))
		for _, c := range candidates {
			fmt.Println(c)
		}
		//elapsed := time.Since(start)
		//log.Printf("Execution took %s", elapsed)
	}
}

func process(octo octopus.Octopus) []string {

	var candidates []string
	projects := octo.GetProjects()
	for _, project := range projects {
		if project.IsDisabled {
			continue
		}
		var prodIds []int
		var devIds []int
		var devEnvId string //t
		var prodEnvId string
		var prodChanId string
		var devChanid string
		var latestProd octopus.Releases
		var latestDev octopus.Releases
		//Output project ID & Name
		//fmt.Println(fmt.Sprintf("%s - %s", project.ID, project.Name))
		projects := octo.GetProjectProgress(project.ID)
		channels := octo.GetProjectChannels(project.ID)
		//project := octo.GetProject("Projects-562")
		for _, c := range channels.Items {
			switch c.Name {
			//case "Development":
			case octo.Config.Channels.Dev:
				devChanid = c.ID
			//case "Production":
			case octo.Config.Channels.Prod:
				prodChanId = c.ID
			}
		}

		for _, e := range projects.ChannelEnvironments[prodChanId] {
			if e.Name == octo.Config.Envs.Test {
				prodEnvId = e.ID
			}
		}
		for _, e := range projects.ChannelEnvironments[devChanid] {
			if e.Name == octo.Config.Envs.Test {
				devEnvId = e.ID
			}
		}

		for id, r := range projects.Releases {
			switch r.Channel.Name {
			//case "Development":
			case octo.Config.Channels.Dev:
				if _, ok := r.Deployments[devEnvId]; ok {
					devIds = append(devIds, id)
				}
			//case "Production":
			case octo.Config.Channels.Prod:
				if _, ok := r.Deployments[prodEnvId]; ok {
					prodIds = append(prodIds, id)
				}
			}
		}
		if len(prodIds) > 1 {
			var lastTime time.Time
			for _, id := range prodIds {
				rel := projects.Releases[id]
				if lastTime.Unix() < rel.Deployments[prodEnvId][0].CompletedTime.Unix() {
					lastTime = rel.Deployments[prodEnvId][0].CompletedTime
					latestProd = rel
				}
			}
		} else if len(prodIds) > 0 {
			latestProd = projects.Releases[prodIds[0]]
		}
		if len(devIds) > 1 {
			var lastTime time.Time
			for _, id := range devIds {
				rel := projects.Releases[id]
				if lastTime.Unix() < rel.Deployments[devEnvId][0].CompletedTime.Unix() {
					lastTime = rel.Deployments[devEnvId][0].CompletedTime
					latestDev = rel
				}
			}
		}

		if len(latestDev.Deployments[devEnvId]) > 0 && len(latestProd.Deployments[prodEnvId]) > 0 {
			if latestDev.Deployments[devEnvId][0].CompletedTime.Unix() > latestProd.Deployments[prodEnvId][0].CompletedTime.Unix() {
				candidates = append(candidates, fmt.Sprintf("%s %s", project.Name, latestDev.Deployments[devEnvId][0].ReleaseVersion))
			} else {
				//	fmt.Println(fmt.Sprintf("%s Not canidate for release", project.Name))
			}
		}
	}
	return candidates
}

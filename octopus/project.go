package octopus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func (o *Octopus) GetProjects() []Project2 {
	resp := o.doRequest("projects/all", "")
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal("Error closing Body")
		}
	}()

	body, _ := ioutil.ReadAll(resp.Body)
	//bodyString := string(body)
	project := []Project2{}
	err := json.Unmarshal(body, &project)
	if err != nil {
		fmt.Println("Error unmarshalling")
		fmt.Println(err)
		os.Exit(2)
	}
	return project
}

func (o *Octopus) GetProject(id string) *Project {
	resp := o.doRequest(fmt.Sprintf("projects/%s",id), "")
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal("Error closing Body")
		}
	}()

	body, _ := ioutil.ReadAll(resp.Body)
	//bodyString := string(body)
	project := Project{}
	err := json.Unmarshal(body, &project)
	if err != nil {
		fmt.Println("Error unmarshalling")
		fmt.Println(err)
		os.Exit(2)
	}
	return &project
}


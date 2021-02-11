package octopus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Projects-1
//Gets Projects Progress
func (o *Octopus) GetProjectProgress(project string) *ProgressResponse {
	resp := o.doRequest(fmt.Sprintf("progression/%s/", project), "")
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal("Error closing Body")
		}
	}()

	body, _ := ioutil.ReadAll(resp.Body)
	//bodyString := string(body)
	pr := ProgressResponse{}
	err := json.Unmarshal(body, &pr)
	if err != nil {
		fmt.Println("Error unmarshalling")
		fmt.Println(err)
		os.Exit(2)
	}
	return &pr
}

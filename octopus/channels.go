package octopus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func (o *Octopus) GetProjectChannels(project string) *ChannelsResponse {
	spaceId := "Spaces-1"
	resp := o.doRequest(fmt.Sprintf("%s/projects/%s/channels", spaceId, project), "")
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal("Error closing Body")
		}
	}()

	body, _ := ioutil.ReadAll(resp.Body)
	//bodyString := string(body)
	pr := ChannelsResponse{}
	err := json.Unmarshal(body, &pr)
	if err != nil {
		fmt.Println("Error unmarshalling")
		fmt.Println(err)
		os.Exit(2)
	}
	return &pr
}

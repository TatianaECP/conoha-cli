package servers

import (
	"encoding/json"
	"fmt"

	"github.com/miyabisun/conoha-cli/endpoints"
)

type VmList struct {
	Servers []Server
}
type Server struct {
	Id       string
	Status   string
	Metadata struct {
		Instance_name_tag string
	}
}

func Get(tenantId string, tokenId string, servers *[]Server) error {
	res := &endpoints.Response{}
	url := fmt.Sprintf("https://compute.tyo1.conoha.io/v2/%s/servers/detail", tenantId)
	err := endpoints.Get(url, tokenId, res)
	if err != nil {
		return err
	}

	vmlist := &VmList{}
	err = json.Unmarshal(res.Body, &vmlist)
	if err != nil {
		return err
	}
	*servers = vmlist.Servers

	return nil
}

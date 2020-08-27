package apisix

import (
	"encoding/json"
	"fmt"
	"testing"

	v1 "github.com/gxthrj/apisix-types/pkg/apis/apisix/v1"
)

func Test_convert2UpstreamRequestIgnoreNodes(t *testing.T) {
	id := "000001"
	uType := "hash"
	HashOn := "var"
	desc := "cloud_httpserver_8080"

	upstream := v1.Upstream{
		ID:     &id,
		Type:   &uType,
		HashOn: &HashOn,
		Name:   &desc,
	}
	ur := convert2UpstreamRequestIgnoreNodes(&upstream)
	b, _ := json.Marshal(ur)
	fmt.Println(string(b))
}

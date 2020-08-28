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

func Test_listUpstreams(t *testing.T) {
	var upstreamsResponse UpstreamsResponse
	ret := `{
		"node":{
			"nodes":[
				{
					"value":{
						"hash_on":"vars",
						"nodes":{
							"10.56.102.127:80":100
						},
						"key":"hash_key",
						"desc":"cloud_test_80",
						"type":"roundrobin"
					},
					"createdIndex":81288,
					"key":"/apisix/upstreams/00000000000000077429",
					"modifiedIndex":81288
				},
				{
					"value":{
						"hash_on":"vars",
						"nodes":{
							"10.56.104.222:9088":100
						},
						"desc":"cloud_httpserver_9088",
						"type":"roundrobin"
					},
					"createdIndex":79815,
					"key":"/apisix/upstreams/00000000000000078454",
					"modifiedIndex":79815
				}
			]
		}
	}`
	if err := json.Unmarshal([]byte(ret), &upstreamsResponse); err != nil {
		return
	} else {
		upstreams := make([]*v1.Upstream, 0)
		for _, u := range upstreamsResponse.Upstreams.Upstreams {
			if n, err := u.convert(""); err == nil {
				upstreams = append(upstreams, n)
			} else {

			}
		}
		b, _ := json.Marshal(upstreams)
		fmt.Println(string(b))
	}
}

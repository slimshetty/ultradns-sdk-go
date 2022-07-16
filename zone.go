package udnssdk

import (
	"fmt"
	"net/http"
	"strings"
	_ "time"
)

type ZoneService interface {
	SelectWithOffsetWithLimit(k *ZoneKey, page string, limit int) ([]Zone, CursorInfo, *http.Response, error)
}

// ZoneService provides access to Zone resources
type ZoneServiceHandler struct {
	client *Client
}

// Zone wraps an Zone resource
type Zone struct {
	Properties struct {
		Name                 string `json:"name"`
		AccountName          string `json:"accountName"`
		Type                 string `json:"type"`
		DnssecStatus         string `json:"dnssecStatus"`
		Status               string `json:"status"`
		Owner                string `json:"owner"`
		ResourceRecordCount  int    `json:"resourceRecordCount"`
		LastModifiedDateTime string `json:"lastModifiedDateTime"`
	} `json:"properties"`
	PrimaryNameServers NameServerLists `json:"primaryNameServers"`
}

type NameServerLists struct {
	NameServerList map[string]interface{} `json:"nameServerIpList"`
}

type CursorInfo struct {
	First    string `json:"first"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Last     string `json:"last"`
}

// ZoneListDTO wraps a list of Zone resources
type ZoneListDTO struct {
	Zones      []Zone     `json:"zones"`
	Queryinfo  QueryInfo  `json:"queryInfo"`
	Cursorinfo CursorInfo `json:"cursorInfo"`
}

// ZoneKey collects the identifiers of a Zone
type ZoneKey struct {
	Zone        string
	AccountName string
}

// URI generates the URI for an Zone
func (k ZoneKey) URI() string {
	//Escaping reverse domain
	zoneName := strings.Replace(k.Zone, "/", "%2F", -1)
	uri := fmt.Sprintf("v3/zones/?&q=name:%s", zoneName)
	if k.AccountName != "" {
		//Escaping space character
		accountName := strings.Replace(k.AccountName, " ", "%2520", -1)
		uri += fmt.Sprintf("+account_name:%s", accountName)
	}
	return uri
}

// QueryURI generates the query URI for an Zone and offset
func (k ZoneKey) QueryURI(page string, limit int) string {
	if page != "" {
		return fmt.Sprintf("%s&cursor=%s&limit=%d", k.URI(), page, limit)
	} else {
		return fmt.Sprintf("%s&limit=%d", k.URI(), limit)
	}
}

// SelectWithOffset requests zone rrsets by ZoneKey & optional offset
func (s *ZoneServiceHandler) SelectWithOffsetWithLimit(k *ZoneKey, page string, limit int) ([]Zone, CursorInfo, *http.Response, error) {
	var zoneld ZoneListDTO

	uri := k.QueryURI(page, limit)
	res, err := s.client.get(uri, &zoneld)
	return zoneld.Zones, zoneld.Cursorinfo, res, err
}

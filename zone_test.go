package udnssdk

import (
	_ "encoding/json"
	"log"
	_ "os"
	_ "reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Trying to run SelectWithOffsetWithLimit function
func Test_Zone_SelectWithOffsetWithLimit_WithZoneKey(t *testing.T) {
	if !enableIntegrationTests {
		t.SkipNow()
	}
	if !enableZoneTests {
		t.SkipNow()
	}

	testClient, err := NewClient(testUsername, testPassword, testBaseURL)
	if err != nil {
		t.Fatal(err)
	}

	r := &ZoneKey{
		Zone: testDomain,
	}
	t.Logf("SelectWithOffsetWithLimit(%v)", r)
	zones := []Zone{}
	maxerrs := 5
	waittime := 5 * time.Second
	errcnt := 0
	page := ""
	limit := 1000

	for {
		reqZones, ri, res, err := testClient.Zone.SelectWithOffsetWithLimit(r, page, limit)
		if err != nil {
			if res != nil && (res.StatusCode >= 500) {
				errcnt = errcnt + 1
				if errcnt < maxerrs {
					time.Sleep(waittime)
					continue
				}
			}
			t.Fatal(err)
		}

		log.Printf("ResultInfo: %+v\n", ri)
		for _, zone := range reqZones {
			zones = append(zones, zone)
		}

		if ri.Next == "" {
			t.Logf("zones: %v", zones)
			return
		} else {
			page = ri.Next
			continue
		}

	}
	assert.Equal(t, zones[0].Properties.Name, testDomain)

}

func Test_Zone_SelectWithOffsetWithLimit_WithOutAnyValue(t *testing.T) {

	if !enableIntegrationTests {
		t.SkipNow()
	}
	if !enableZoneTests {
		t.SkipNow()
	}

	testClient, err := NewClient(testUsername, testPassword, testBaseURL)
	if err != nil {
		t.Fatal(err)
	}

	zones := []Zone{}
	maxerrs := 5
	waittime := 5 * time.Second
	errcnt := 0
	page := ""
	limit := 1000

	for {
		reqZones, ri, res, err := testClient.Zone.SelectWithOffsetWithLimit(&ZoneKey{}, page, limit)
		if err != nil {
			if res != nil && (res.StatusCode >= 500) {
				errcnt = errcnt + 1
				if errcnt < maxerrs {
					time.Sleep(waittime)
					continue
				}
			}
			t.Fatal(err)
		}

		log.Printf("ResultInfo: %+v\n", ri)
		for _, zone := range reqZones {
			zones = append(zones, zone)
		}
		if ri.Next == "" {
			t.Logf("zones: %v", zones)
			return
		} else {
			page = ri.Next
			continue
		}
	}
	assert.NotNil(t, zones)
}

// Trying to run function with account or zone not found
func Test_Zone_InvalidZone(t *testing.T) {
	if !enableIntegrationTests {
		t.SkipNow()
	}
	if !enableZoneTests {
		t.SkipNow()
	}

	testClient, err := NewClient(testUsername, testPassword, testBaseURL)
	if err != nil {
		t.Fatal(err)
	}

	r := &ZoneKey{
		Zone: "abcdef-test23.com",
	}
	t.Logf("SelectWithOffsetWithLimit(%v)", r)
	page := ""
	limit := 1000
	_, _, _, err = testClient.Zone.SelectWithOffsetWithLimit(r, page, limit)
	assert.NotNil(t, err)
}

// Trying to run function with account not found
func Test_Zone_InvalidAccount(t *testing.T) {
	if !enableIntegrationTests {
		t.SkipNow()
	}
	if !enableZoneTests {
		t.SkipNow()
	}

	testClient, err := NewClient(testUsername, testPassword, testBaseURL)
	if err != nil {
		t.Fatal(err)
	}

	r := &ZoneKey{
		AccountName: "sddsfffrefref",
	}
	t.Logf("SelectWithOffsetWithLimit(%v)", r)
	page := ""
	limit := 1000

	_, _, _, err = testClient.Zone.SelectWithOffsetWithLimit(r, page, limit)
	assert.NotNil(t, err)

}

// Trying to run testcase for QueryURI Function
func Test_Zone_AccountNameWithSpace(t *testing.T) {
	assert := assert.New(t)
	accountName := "team%2520rest1"
	r := &ZoneKey{
		AccountName: "team rest1",
	}
	uri := r.QueryURI("", 100)
	t.Logf("URI: %s", uri)
	accountCheck := strings.Contains(uri, accountName)
	assert.Equal(true, accountCheck, "Both should be true")
}

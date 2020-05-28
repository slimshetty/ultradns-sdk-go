package udnssdk

import (
	_ "encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	_ "os"
	_ "reflect"
	"strings"
	"testing"
	"time"
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
	offset := 0
	limit := 1000

	for {
		reqZones, ri, res, err := testClient.Zone.SelectWithOffsetWithLimit(r, offset, limit)
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
		if ri.ReturnedCount+ri.Offset >= ri.TotalCount {
			t.Logf("zones: %v", zones)
			return
		}
		offset = ri.ReturnedCount + ri.Offset
		continue
		if err != nil {
			t.Fatal(err)
		}

	}

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
	offset := 0
	limit := 1000

	for {
		reqZones, ri, res, err := testClient.Zone.SelectWithOffsetWithLimit(&ZoneKey{}, offset, limit)
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
		if ri.ReturnedCount+ri.Offset >= ri.TotalCount {
			t.Logf("zones: %v", zones)
			return
		}
		offset = ri.ReturnedCount + ri.Offset
		continue
	}

	if err != nil {
		t.Fatal(err)
	}
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
	maxerrs := 5
	waittime := 5 * time.Second
	errcnt := 0
	offset := 0
	limit := 1000

	for {
		_, _, res, err := testClient.Zone.SelectWithOffsetWithLimit(r, offset, limit)
		if err != nil {
			if res != nil && (res.StatusCode >= 500) {
				errcnt = errcnt + 1
				if errcnt < maxerrs {
					time.Sleep(waittime)
					continue
				}
			}
			t.Logf("%v", err)
			return
		} else {
			t.Fatal(" Expected to fail")

		}
	}
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
	maxerrs := 5
	waittime := 5 * time.Second
	errcnt := 0
	offset := 0
	limit := 1000

	for {
		_, _, res, err := testClient.Zone.SelectWithOffsetWithLimit(r, offset, limit)
		if err != nil {
			if res != nil && (res.StatusCode >= 500) {
				errcnt = errcnt + 1
				if errcnt < maxerrs {
					time.Sleep(waittime)
					continue
				}
			}
			t.Logf("%v", err)
			return
		} else {
			t.Fatal(" Expected to fail")

		}
	}
}

// Trying to run testcase for QueryURI Function
func Test_Zone_AccountNameWithSpace(t *testing.T) {
	assert := assert.New(t)
	accountName := "team%2520rest1"
	r := &ZoneKey{
		AccountName: "team rest1",
	}
	uri := r.QueryURI(0, 100)
	t.Logf("URI: %s", uri)
	accountCheck := strings.Contains(uri, accountName)
	assert.Equal(true, accountCheck, "Both should be true")
}

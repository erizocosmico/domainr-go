package domainr

import (
	"fmt"
	"testing"
)

func Test_Register(t *testing.T) {
	result, err := Register("google.es", "")
	if err != nil {
		fmt.Println(err)
	} else {
		url := "http://www.shareasale.com/r.cfm?b=210737&u=303669&m=25581&afftrack=&urllink=iwantmyname.com/search/add/google.es%3Fr=domai.nr"
		if result != url {
			t.Errorf("Expecting: %s, %s received.", url, result)
		}
	}

	result, err = Register("myfancydomain.io", "gandi.net")
	if err != nil {
		fmt.Println(err)
	} else {
		url := "https://www.gandi.net/domain/buy/result/?domain_list=myfancydomain&tld=io"
		if result != url {
			t.Errorf("Expecting: %s, %s received.", url, result)
		}
	}
}

func Test_Json_Info(t *testing.T) {
	result, err := Json(METHOD_SEARCH, "myfancydomain", "")
	if err != nil {
		fmt.Println(err)
	} else {
		rLen := len(result.(map[string]interface{})["results"].([]interface{}))
		if rLen != 19 {
			t.Errorf("Expecting: %i, %i received.", 19, rLen)
		}
	}
}

func Test_Json_Info_Invalid(t *testing.T) {
	result, err := Json(METHOD_INFO, "myfancydomain", "")
	if err != nil {
		fmt.Println(err)
	} else {
		errMsg := result.(map[string]interface{})["error_message"].(string)
		if errMsg != "Invalid domain" {
			t.Errorf("Expecting: %s, %s received.", "Invalid domain", errMsg)
		}
	}
}

func Test_Json_Search(t *testing.T) {
	result, err := Json(METHOD_INFO, "github.io", "")
	if err != nil {
		fmt.Println(err)
	} else {
		errMsg := result.(map[string]interface{})["availability"].(string)
		if errMsg != "taken" {
			t.Errorf("Expecting: %s, %s received.", "taken", errMsg)
		}
	}
}

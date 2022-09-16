package halo

import (
	"encoding/json"
	"github.com/duke-git/lancet/v2/netutil"
	"io/ioutil"
	code2 "server/code"
)

func (*ServiceHaloGroup) GetTags(url string, token string) (tags interface{}, code int, err error) {
	url = url + "/api/admin/tags"
	header := map[string]string{
		"Content-Type":        "application/json",
		"Admin-Authorization": token,
	}
	resp, err := netutil.HttpGet(url, header, nil)
	if err != nil {
		return "", code2.HaloServerError, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//从body拿出数据
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", code2.HaloServerError, err
	}
	return result["data"], code2.SUCCESS, nil
}

func (*ServiceHaloGroup) GetCategory(url string, token string) (tags interface{}, code int, err error) {
	url = url + "/api/admin/categories"
	header := map[string]string{
		"Content-Type":        "application/json",
		"Admin-Authorization": token,
	}
	resp, err := netutil.HttpGet(url, header, nil)
	if err != nil {
		return "", code2.HaloServerError, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//从body拿出数据
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", code2.HaloServerError, err
	}
	return result["data"], code2.SUCCESS, nil
}

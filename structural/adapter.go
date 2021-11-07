/*
ActivityAdapter adapts map[string]interface{} data format to Activity structure format
main.go has a http REST request to API which returns map[string]interface{} containing Activity data
*/
package main

import (
	"fmt"
)

type Activity struct {
	title        string
	kind         string
	url          string
	participants int8
	price        float32
}

func (a Activity) String() string {
	_one := fmt.Sprintf("%v kind of activity \"%v\" is for %v participant(s)", a.kind, a.title, a.participants)
	var _two, _three string
	if a.price > 0 {
		_two = fmt.Sprintf(", it costs %v", a.price)
	} else if a.price == 0 {
		_two = ", it's free"
	} else {
		_two = "it's priceless"
	}
	if a.url != "" {
		_three = fmt.Sprintf(", link:\n%v", a.url)
	}
	return _one + _two + _three
}

func ActivityAdapter(activityMap map[string]interface{}) *Activity {
	return &Activity{
		title:        activityMap["activity"].(string),
		kind:         activityMap["type"].(string),
		url:          activityMap["link"].(string),
		participants: int8(activityMap["participants"].(float64)),
		price:        float32(activityMap["price"].(float64)),
	}
}

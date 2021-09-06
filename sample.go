package main

import (
	"errors"
	"fmt"

	"github.com/swaraj1802/GoJSON/gojson"
)

func main() {
	jsonParsed, err := gojson.ParseJSON([]byte(`{
		"ok": true,
		"channel": "C1H9RESGL",
		"ts": "1503435956.000247",
		"message": {
			"text": "Here's a message for you",
			"username": "ecto1",
			"bot_id": "B19LU7CSY",
			"attachments": [
				{
					"text": "This is an attachment",
					"id": 1,
					"fallback": "This is an attachment's fallback"
				},
				{
					"text": "This is an attachment2",
					"id": 2,
					"fallback": "This is an attachment's fallback2"
				}
			],
			"type": "message",
			"subtype": "bot_message",
			"ts": "1503435956.000247"
		}
	}`))
	if err != nil {
		panic(err)
	}

	value, ok := jsonParsed.Search("message", "attachments", "*", "id")
	if ok != nil {
		panic(errors.New("Element doesn't exist"))
	}
	output := value.JSONData()
	# See the output
	fmt.Println(output)
}

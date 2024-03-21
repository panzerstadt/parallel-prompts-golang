package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func label(description string) (string, error) {
	API_KEY := os.Getenv("OPENAI_API_KEY")
	posturl := "https://api.openai.com/v1/chat/completions"
	model := "gpt-3.5-turbo"

	messages := []Message{
		{Role: "system", Content: fmt.Sprintf("You are a helpful assistant designed to label descriptions with one of the following categories: %s. you only output JSON with two fields: description, category.", categories)},
		{Role: "user", Content: "Some context: I live and work in Tokyo, Japan. my hometown is in Kuala Lumpur, Malaysia. My hobbies are gaming, stationeries, and electronic gadgets."},
		{Role: "user", Content: `Here are some previous description labels: Akasaka Biz Tower,CAFE
		Ueshima Coffee Akasaka Hitotsugi,CAFE
		Ueshima Coffee Akasaka Hitotsugi,CAFE
		AIA IPOS ONLINE (Country: MY),TRAVEL
		AEON SMKT-METRO PRIM (Country: MY),DAILY
		SB367-SANDAKAN AIRPO (Country: MY),TRAVEL
		JIM-KSLMTN DAN PASSP (Country: MY),MISC
		Rakuten Mobile Upgrade,UTILITIES
		WREN CLIMATE PBC (Country: US),SUBSCRIPTIONS
		Rakuten Mobile Communication Fee,PHONE
		Rakuten Cafe 9F,CAFE
		TRAWICK TRAVEL INSUR (Country: US),TRAVEL
		660174 GOOGLE C,SUBSCRIPTIONS
		135065 WING TAK,MISC
		192110 DOUTOR C,CAFE
		231067 AMAZON C (AIRTAGS),MISC
		464149 APPLE.CO,SUBSCRIPTIONS
		093830 PATREON,SUBSCRIPTIONS
		809143 DOUTOR C,CAFE
		825463 DOUTOR C,CAFE
		162103 PAYPAL (WIKIMEDIA),SUBSCRIPTIONS
		361142 RAKUTENP,SOCIAL
		469885 AMAZON C (TIME BLOCK PLANNER),HOBBIES
		611165 KAI ANJI,TRAVEL
		640420 PAYPAL (PRINTFUL),HOBBIES
		698172 JR EAST,TRAVEL
		Mizuho Bank,SAVINGS
		807276 STEAMGAM,HOBBIES
		957286 UESHIMAK,CAFE
		967060 UESHIMAK,CAFE
		037358 UESHIMAK,CAFE
		096420 RAMEN KI,CAFE
		103450 BIG CAME (UBER EATS),DAILY
		285483 AMAZON C,MISC
		330439 MICROSOFT,SUBSCRIPTIONS
		335876 STARBUCK,CAFE
		502471 NEBULA S,SUBSCRIPTIONS
		968085 AMAZON C,MISC
		053760 AMAZON C,MISC
		853955 ADOBE CR,SUBSCRIPTIONS
		502471 NEBULA S,SUBSCRIPTIONS
		473003 APPLE.CO,SUBSCRIPTIONS
		853955 ADOBE CR,SUBSCRIPTIONS
		968085 AMAZON.C,MISC
		053760 AMAZON.C,MISC
		934042 STEAM PU,HOBBIES
		4th Year 12th Month,MISC
		502357 AIRASIA,TRAVEL
		RTK Merpay,CREDIT CARD
		Rakuten Card Service,CREDIT CARD
		502357 AIRASIA,TRAVEL
		502357 AIRASIA,TRAVEL
		Rakuten Card Service,CREDIT CARD
		RTK Merpay,CREDIT CARD
		502357 AIRASIA,TRAVEL
		4th Year 12th Month,MISC
		934042 STEAM PU,HOBBIES
		053760 AMAZON.C,MISC
		968085 AMAZON.C,MISC
		853955 ADOBE CR,SUBSCRIPTIONS
		473003 APPLE.CO,SUBSCRIPTIONS
		502471 NEBULA S,SUBSCRIPTIONS
		853955 ADOBE CR,SUBSCRIPTIONS
		053760 AMAZON C,MISC
		968085 AMAZON C,MISC
		502471 NEBULA S,SUBSCRIPTIONS
		335876 STARBUCK,CAFE
		330439 MICROSOFT,SUBSCRIPTIONS
		285483 AMAZON C,MISC
		103450 BIG CAME (UBER EATS),DAILY
		096420 RAMEN KI,CAFE
		037358 UESHIMAK,CAFE
		967060 UESHIMAK,CAFE
		957286 UESHIMAK,CAFE
		807276 STEAMGAM,HOBBIES
		Mizuho Bank,SAVINGS
		698172 JR EAST,TRAVEL
		640420 PAYPAL (PRINTFUL),HOBBIES
		611165 KAI ANJI,TRAVEL
		469885 AMAZON C (TIME BLOCK PLANNER),HOBBIES
		361142 RAKUTENP,SOCIAL
		162103 PAYPAL (WIKIMEDIA),SUBSCRIPTIONS
		825463 DOUTOR C,CAFE
		809143 DOUTOR C,CAFE
		093830 PATREON,SUBSCRIPTIONS
		464149 APPLE.CO,SUBSCRIPTIONS
		231067 AMAZON C (AIRTAGS),MISC
		192110 DOUTOR C,CAFE
		135065 WING TAK,MISC
		660174 GOOGLE C,SUBSCRIPTIONS
		TRAWICK TRAVEL INSUR (Country: US),TRAVEL
		Rakuten Cafe 9F,CAFE
		Rakuten Mobile Communication Fee,PHONE
		WREN CLIMATE PBC (Country: US),SUBSCRIPTIONS
		Rakuten Mobile Upgrade,UTILITIES
		JIM-KSLMTN DAN PASSP (Country: MY),MISC
		SB367-SANDAKAN AIRPO (Country: MY),TRAVEL
		AEON SMKT-METRO PRIM (Country: MY),DAILY
		AIA IPOS ONLINE (Country: MY),TRAVEL
		Ueshima Coffee Akasaka Hitotsugi,CAFE
		Ueshima Coffee Akasaka Hitotsugi,CAFE
		Akasaka Biz Tower,CAFE`},
		{Role: "user", Content: fmt.Sprintf("What is the label for this description: \"%s\"", description)},
	}

	payload := Payload{
		Model:       model,
		Seed:        seed,
		Temperature: temperature,
		Format: Format{
			Type: "json_object",
		},
		Messages: messages,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	body := strings.NewReader(string(jsonData))

	r, err := http.NewRequest("POST", posturl, body)
	if err != nil {
		return "", err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", API_KEY))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	response := &Response{}
	derr := json.NewDecoder(res.Body).Decode(response)
	if derr != nil {
		return "", derr
	}

	// jsonResponse, err := json.Marshal(response)
	// if err != nil {
	// 	return "", err
	// }

	// fmt.Println(string(jsonResponse))

	escapedContent, err := json.Marshal(response.Choices[0].Message.Content)
	if err != nil {
		return "", err
	}

	// fmt.Println("content: ", string(escapedContent))

	content, err := strconv.Unquote(string(escapedContent))
	if err != nil {
		return "", err
	}
	// fmt.Println("content: ", content)

	// "{\n    \"description\": \"Akasaka Biz Tower\",\n    \"category\": \"travel\"\n}"
	chatContent := &ChatContent{}
	err = json.Unmarshal([]byte(content), chatContent)
	if err != nil {
		return "", err
	}

	answer := chatContent.Category

	return string(answer), nil
}

type ChatContent struct {
	Description string `json:"description"`
	Category    string `json:"category"`
}

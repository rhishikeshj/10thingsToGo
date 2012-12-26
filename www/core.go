package www

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
)

type appLaunchedResponse struct {
	Id string
	Timestamp int64
}

func removeDuplicates(a []string) []string {
        result := []string{}
        seen := map[string]int{}
        for _, val := range a {
                if _, ok := seen[val]; !ok {
                        result = append(result, val)
                        seen[val] = 1
                }
        }
        return result
}

func api_handler (writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		post_api_handler(writer, request)
	} else if request.Method == "GET" {
		get_api_handler(writer, request)
	}
}

func get_api_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "GET request handler not implemented yet !")
}

func post_api_handler(writer http.ResponseWriter, request *http.Request) {

	apiString := request.URL.String()[11:]
	
	switch apiString {
	case "app-launched/", "app-launched":
		timeStamp := time.Now() 
		responseMessage := appLaunchedResponse{"id_10things_one", timeStamp.Unix()};
		responseJSON, err := json.Marshal(responseMessage)
		if err == nil {
			fmt.Println("Posting something")
			fmt.Fprintf(writer, string(responseJSON[:]))
		}
	case "login/", "login" :
	}
}

func init() {
	http.HandleFunc("/api/lib/1/", api_handler)

}

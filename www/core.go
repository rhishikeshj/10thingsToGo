package www

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"appengine"
	"appengine/datastore"
//	"strings"
//	"builtin"
)

type appLaunchedResponse struct {
	Id string
	Timestamp int64
	AppID string
}

type List struct {
	Id string
	UserId string
	Title string
	Category string
	ListItems []string
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
//		post_api_handler(writer, request)
		get_api_handler(writer, request)
	}
}

func get_api_handler(writer http.ResponseWriter, request *http.Request) {
	apiString := request.URL.Path[11:]
	switch apiString {
	case "list/", "list":
		listID := request.FormValue("list-id")
		context := appengine.NewContext(request)
			query := datastore.NewQuery("List").Filter("Id =", listID)
		list := make([]List, 0, 10)
		if _, err := query.GetAll(context, &list); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		} else {
			fmt.Fprintf(writer, "The list returned is : %s :", list[0].ListItems[0])
		}
		responseMessage := List{listID, "Rhishikesh", "First list", "General", []string{"first", "second"}}
		responseJSON, err := json.Marshal(responseMessage)
		if err == nil {
			fmt.Fprintf(writer, string(responseJSON[:]))
		}
	}
}

func post_api_handler(writer http.ResponseWriter, request *http.Request) {
	context := appengine.NewContext(request)
	apiString := request.URL.Path[11:]
	switch apiString {
	case "app-launched/", "app-launched":
		userId := request.FormValue("user-id")
		timeStamp := time.Now()
		responseMessage := appLaunchedResponse{"id_10things_one", timeStamp.Unix(), userId};
		responseJSON, err := json.Marshal(responseMessage)
		if err == nil {
			fmt.Fprintf(writer, string(responseJSON[:]))
		}
	case "list/", "list" :
		userId := request.FormValue("user-id")
		title := request.FormValue("title")
		category := request.FormValue("category")
		objJSON := make(map[string]interface{})
		unmarshal_err := json.Unmarshal([]byte(request.FormValue("list")), &objJSON)
		if unmarshal_err != nil {
			http.Error(writer, unmarshal_err.Error(), http.StatusInternalServerError)
		}
		listJSON := objJSON["list"]
		list := make([]string, 1)
		for _,v := range listJSON.([]interface{})  {
			list = append(list, fmt.Sprintf("%s",v))
		} 
		newList := List{"42", userId, title, category, list}
		_, err := datastore.Put(context, datastore.NewIncompleteKey(context, "List", nil), &newList)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(writer, "Success")
	}
}

func init() {
	http.HandleFunc("/api/lib/1/", api_handler)

}

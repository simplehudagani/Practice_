package list

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	dbs "sum/connection"
	"testing"

	//"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
)

// type Response struct {
// 	Message string `json:"message"`
// }

func TestSpoorthiList(t *testing.T) {
	var c SpoorthiBaselineQuestionnaire
	var _ = dbs.Connection()
	requestbody := []byte(`"Module":"Module1"`)
	req, err := http.NewRequest("GET", "/spoorthilist", bytes.NewBuffer(requestbody))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req)
	rr := httptest.NewRecorder()
	fmt.Println(rr)
	fmt.Println(rr.Code)
	handler := http.HandlerFunc(Spoorthilist)
	handler.ServeHTTP(rr, req)

	fmt.Println(rr.Code)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		// }

	}
	body := rr.Body.String()

	fmt.Println(body)

	decodedBody, _ := ioutil.ReadAll(rr.Body)
	//var c []Asset
	_ = json.Unmarshal(decodedBody, &c)
	fmt.Println(c)

	str := "SELECT firstName FROM training_participants train,SpoorthiBaselineQuestionnaire spoo where (spoo.partcipantId=train.id);"

	rows, err := db.Query(str)
	if err != nil {
		fmt.Println(err)
		//w.WriteHeader(http.StatusBadRequest)
		//json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
		return
	}

	var FirstName string

	type sporti struct {
		FirstName string
	}

	var result []sporti // creating slice

	for rows.Next() {

		err := rows.Scan(&FirstName)
		if err != nil {
			fmt.Println(err)
		}

		result = append(result, sporti{FirstName: FirstName})

	}
	fmt.Println(result)
	assert.Equal(t, body, result)
}

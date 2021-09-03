package jsonmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type testStruct struct {
	Tests []struct {
		Id int `json:"global_id"`
	}
}

func JsonEr() {
	datafile, _ := os.Open("jsonmanager/data.json")
	result, _ := ioutil.ReadAll(datafile)
	var resultEr testStruct
	json.Unmarshal(result, &resultEr)
	// fmt.Print(resultEr.Tests)
	var summEr int
	for _, v := range resultEr.Tests {
		summEr += v.Id
	}
	fmt.Println(summEr)
	fmt.Println("4 program")
}

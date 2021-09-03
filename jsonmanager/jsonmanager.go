package jsonmanager

import (
	"encoding/json"
	"fmt"
)

type (
	Student struct {
		Rating []int
	}

	Group struct {
		Students []Student
	}

	Rating struct {
		Average float32
	}
)

var jsonBlob = []byte(`{
    "ID":11,
    "Number":"ИЛМ-1274",
    "Year":2,
    "Students":[
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3]
        },
        {
            "LastName":"Марк",
            "FirstName":"Твен",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[5,2]
        }
    ]
}`)

func JsonManager() {
	var result Group
	json.Unmarshal(jsonBlob, &result)
	// fmt.Println(result.Students)
	var counterStudent int
	var counterRating int
	var ccc int
	for _, v := range result.Students {
		counterStudent++
		for _, k := range v.Rating {
			ccc += k
			counterRating++
		}
	}
	s := Rating{Average: float32(counterRating) / float32(counterStudent)}

	data, _ := json.MarshalIndent(s, "", "    ")
	fmt.Printf("%s\n", data)
	fmt.Println("3 program")
}

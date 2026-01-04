package main

import (
	"encoding/json"
	"net/http"
	"io"
)


func fetchLocationAreas(str *string) (*string,*string,[]string,error){

	client:= &http.Client{}
	resp,err := client.Get(*str)

	if err!=nil{
		return nil,nil,nil,err
	}

	defer resp.Body.Close()

	body,err := io.ReadAll(resp.Body)
	if err!=nil{
		return nil,nil,nil,err
	}

	var LocationRespnse locationAreaResponse
	err = json.Unmarshal(body,&LocationRespnse)
	if err!= nil{
		return nil,nil,nil,err
	}
	nxt := LocationRespnse.Next
	prev := LocationRespnse.Previous
	var Location []string
	for _,i := range(LocationRespnse.Results){
		Location = append(Location, i.Name)
	}
	return nxt,prev,Location,nil
}

package main

import(

)
type locationArea struct{
	Name string `json:"name"`
}

type locationAreaResponse struct{
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []locationArea `json:"results"`
}
package main

type Data struct {
	Results []Results `json:"results"`
	Info    Info      `json:"info"`
}

type Results struct {
	Name  Name   `json:"name"`
	Login Login  `json:"login"`
	Email string `json:"email"`
}

type Info struct {
	Name  Name   `json:"name"`
	Login Login  `json:"login"`
	Email string `json:"email"`
}

type Login struct {
	UUID string `json:"uuid"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

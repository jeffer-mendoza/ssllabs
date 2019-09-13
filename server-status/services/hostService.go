package services

import (
	"../models"
	"../utils"
	"encoding/json"
	"log"
	"net/http"
)

func GetInfoHost(hostname string) models.HostInfo {
	responseEntity := getInfoServer(hostname)
	hostInfo := models.HostInfo{}
	servers := make([]models.Server, len(responseEntity.Endpoints))
	// Iterate through list and print its contents.
	var minorGrade = " "
	for index, element := range responseEntity.Endpoints {
		if element.Grade[0] > minorGrade[0] {
			minorGrade = element.Grade
		}
		organization, country := utils.ExeWhois(element.IPAddress)
		servers[index] = models.Server{element.IPAddress, element.Grade, country, organization}
	}
	hostInfo.Servers = servers
	title, image := utils.ExtractFieldsFromHtml(hostname)
	hostInfo.Title = title
	hostInfo.Logo = image
	hostInfo.SslGrade = minorGrade
	return hostInfo
}

func getInfoServer(host string) models.Response {
	responseEntity := models.Response{}
	//TODO add the best form to handle string replace or concatenation
	resp, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + host)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseEntity)
	if err != nil {
		log.Fatal(err)
	}

	return responseEntity
}

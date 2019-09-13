package services

import (
	"../models"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func GetInfoDomain(hostname string) models.HostInfo {
	responseEntity := getInfoServer(hostname)
	hostInfo := models.HostInfo{}
	servers := make([]models.Server, len(responseEntity.Endpoints))
	// Iterate through list and print its contents.
	var minorGrade = " "
	for index, element := range responseEntity.Endpoints {
		if element.Grade[0] > minorGrade[0] {
			minorGrade = element.Grade
		}
		organization, country := getInfoOwnerServer(element.IPAddress)
		servers[index] = models.Server{element.IPAddress, element.Grade, country, organization}
	}
	hostInfo.Servers = servers
	title, image := extractHtml(hostname)
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

const command = "whois {ipAddress} | sed -n \"/Organization\\|Country/p\""

func getInfoOwnerServer(ipAddress string) (string, string) {
	cmd := strings.Replace(command, "{ipAddress}", ipAddress, -1)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(string(out), "\n")
	return extract(s[0]), extract(s[1])
}

func extract(line string) string {
	return strings.Trim(strings.Split(line, ":")[1], " ")
}

func extractHtml(hostname string) (string, string) {
	// Request the HTML page.
	res, err := http.Get("https://" + hostname)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var image = ""
	var title = ""
	// Find the review items
	doc.Find("link[rel*='icon']").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		image, _ = s.Attr("href")
		return
	})

	// Find the review items
	doc.Find("head title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title = s.Text()
		return
	})

	return title, image
}

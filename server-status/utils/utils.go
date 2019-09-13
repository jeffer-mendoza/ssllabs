package utils

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const command = "whois {ipAddress} | sed -n \"/Organization\\|Country/p\""

/**
Permite realizar la ejecución del comando whois del sistema operativo
para obtener los campos organización y país asocidados el dirección dada
*/
func ExeWhois(ipAddress string) (string, string) {
	cmd := strings.Replace(command, "{ipAddress}", ipAddress, -1)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(string(out), "\n")
	return extract(s[0]), extract(s[1])
}

/**
Realiza la extracción de los valores texto con el siguiente formato -> key: value
*/
func extract(line string) string {
	return strings.Trim(strings.Split(line, ":")[1], " ")
}

/**
Permite realizar la extracción del logo y title de una sitio web
*/
func ExtractFieldsFromHtml(hostname string) (string, string) {
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

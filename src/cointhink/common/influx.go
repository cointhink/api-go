package common

import "time"
import "net/http"
import "strings"
import "log"

import "cointhink/config"

func InfluxWrite(measurement string, tagName string, tagValue string, reading string) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}
	db := config.C.QueryString("influx.database")
	influx_url := config.C.QueryString("influx.url") + "/write?db=" + db
	data := measurement + "," + tagName + "=" + tagValue + " value=" + reading
	log.Printf("InfluxWrite db=%s %s\n", db, data)
	_, err := client.Post(influx_url, "", strings.NewReader(data))
	if err != nil {
		log.Printf("Influx post err %v\n", err)
	} else {
		//log.Printf("Influx response %s %s\n", response.Proto, response.Status)
	}
}

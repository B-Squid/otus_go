package configs

import "fmt"

type AppConfig struct {
	userName            string
	userSalt            string
	authtoken           string
	sdServer            string
	sdServerPort        string
	apiPath             string
	rmqServer           string
	rmqServerPort       string
	pgServer            string
	pgServerPort        string
	filterHeader        string
	filterBody          string
	pollingEnabled      bool
	pollingFrequencySec int
	rowCount            int
	RouteRule           string
}

type configCheck struct {
	regExpSdServer      string
	regExpSdServerPort  string
	regExpRmqServer     string
	regExpRmqServerPort string
	regExpPgServer      string
	regExpPgServerPort  string
}

func (ac *AppConfig) Setter() {
	//read config file
	ac.userName = "Some User"
	ac.userSalt = "k2m3n4n5"
	ac.authtoken = "2354ljkh25b45ik2t4r234goj4g5n924ngkl4gn69g"
	ac.sdServer = "https://servicedesk.domain"
}

func (ac AppConfig) Getter() {
	fmt.Println("User name: ", ac.userName)
}

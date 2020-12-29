package main

import (
	"github.com/maxmind/mmdbwriter/mmdbtype"
	"github.com/oschwald/geoip2-golang"
	log "github.com/sirupsen/logrus"
	"net"
)

func parseCIDRs(txtList []string) []*net.IPNet  {
	var networkList = make([]*net.IPNet,0, 50)
	for _, cidrTxt := range txtList {
		_,network, err := net.ParseCIDR(cidrTxt)
		if err != nil || network == nil {
			log.Printf("%s fail to parse to CIDR\n", cidrTxt)
			continue
		}
		networkList = append(networkList, network)
	}

	return networkList
}

func parseCountry(record *geoip2.Country) mmdbtype.DataType {
	return mmdbtype.Map{
		"country": mmdbtype.Map{
			"geoname_id":           mmdbtype.Uint32(record.Country.GeoNameID),
			"is_in_european_union": mmdbtype.Bool(record.Country.IsInEuropeanUnion),
			"iso_code":             mmdbtype.String(record.Country.IsoCode),
			"names": mmdbtype.Map{
				"de":    mmdbtype.String(record.Country.Names["de"]),
				"en":    mmdbtype.String(record.Country.Names["en"]),
				"es":    mmdbtype.String(record.Country.Names["es"]),
				"fr":    mmdbtype.String(record.Country.Names["fr"]),
				"ja":    mmdbtype.String(record.Country.Names["ja"]),
				"pt-BR": mmdbtype.String(record.Country.Names["pt-BR"]),
				"ru":    mmdbtype.String(record.Country.Names["ru"]),
				"zh-CN": mmdbtype.String(record.Country.Names["zh-CN"]),
			},
		},
	}
}
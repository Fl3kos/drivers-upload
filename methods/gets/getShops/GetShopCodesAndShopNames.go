package getShops

import "strings"

func GetShopCodesAndShopNames(shops []string) ([]string, []string) {
	var shopCodes []string
	var shopNames []string

	for i, _ := range shops {
		shop := strings.Split(shops[i], "-")

		shopCode := shop[0]
		shopName := shop[1]

		shopCode = strings.TrimSpace(shopCode)
		shopName = strings.TrimSpace(shopName)

		//change white space to -
		shopName = strings.ReplaceAll(shopName, " ", "_")
		shopCodes = append(shopCodes, shopCode)
		shopNames = append(shopNames, shopName)
	}

	return shopCodes, shopNames
}

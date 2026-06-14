package thirdparty

import (
	"eattheitch/backend/utils"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

type Response struct {
	Data struct {
		Objects []struct {
			Names struct {
				Common string `json:"common"`
			} `json:"names"`

			Codes struct {
				Alpha2 string `json:"alpha_2"`
			} `json:"codes"`

			Flag struct {
				Emoji string `json:"emoji"`
			} `json:"flag"`
		} `json:"objects"`
	} `json:"data"`
}

func GetCountries(context *gin.Context) {
	apiKey := utils.MustGetEnv("REST_COUNTRIES_API_KEY")

	req, _ := http.NewRequest("GET", "https://api.restcountries.com/countries/v5?region=Europe&response_fields=names.common,codes.alpha_2,flag.emoji&pretty=1", nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var data Response
	json.Unmarshal(body, &data)

	var countries []Country

	for _, obj := range data.Data.Objects {
		countries = append(countries, Country{
			Code: obj.Codes.Alpha2,
			Name: obj.Names.Common,
			Flag: obj.Flag.Emoji,
		})
	}

	context.JSON(http.StatusOK, countries)
}

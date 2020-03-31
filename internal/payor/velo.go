package payor

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/antihax/optional"
	velo "github.com/velopaymentsapi/velo-go"
)

// VeloOAuthRefresh will check if OAuth2 token is about to expire ... if so .. issue event to refresh it
func VeloOAuthRefresh() {
	now := time.Now()
	secs := now.Unix()

	// check if VELO_API_ACCESSTOKENEXPIRATION is expired
	eex, err := strconv.Atoi(os.Getenv("VELO_API_ACCESSTOKENEXPIRATION"))
	if err == nil {
		if eex <= int(secs) {
			print("CALL VELO API TO GET TOKEN")

			cfg := velo.NewConfiguration()
			client := velo.NewAPIClient(cfg)
			args := velo.VeloAuthOpts{}
			args.GrantType = optional.NewString("client_credentials")
			authctx := context.WithValue(context.Background(), velo.ContextBasicAuth, velo.BasicAuth{
				UserName: os.Getenv("VELO_API_APIKEY"),
				Password: os.Getenv("VELO_API_APISECRET"),
			})
			r, h, err := client.LoginApi.VeloAuth(authctx, &args)
			if err != nil {
				fmt.Println("FAILURE: could not authenticate with Velo api", err)
			}

			if h.StatusCode == 200 {
				os.Setenv("VELO_API_ACCESSTOKEN", r.AccessToken)
				ne := (int(secs) + int(r.ExpiresIn) - 300)
				os.Setenv("VELO_API_ACCESSTOKENEXPIRATION", strconv.Itoa(ne))
			}
		} else {
			fmt.Println("velo oauth token has not expired")
		}
	} else {
		fmt.Println("FAILURE: VELO_API_ACCESSTOKENEXPIRATION default value needs to 0")
	}
}

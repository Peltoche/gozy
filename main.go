package main

import (
	"context"
	"log"

	"github.com/Peltoche/cozy-oauth2/sdk"
)

const oauthFilePath = "cozy-cli/oauth.json"

func main() {
	ctx := context.Background()

	client := sdk.NewHTTPClient("https://jeanbon.mycozy.cloud")
	cfg := sdk.NewXDGConfig("test-sdk")

	oauthClient, err := client.RegisterClient(ctx, &sdk.RegisterClientCmd{
		RedirectURIs: []string{"http://localhost:9090"},
		ClientName:   "test-3",
		SoftwareID:   "github.com/Peltoche/cozy-client",
	})
	if err != nil {
		log.Fatal(err)
	}

	cfg.SaveClient(oauthClient)

	// 	app, err := getAppInfo(ctx)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	cfg := oauth2.Config{
	// 		ClientID:     app.ClientID,
	// 		ClientSecret: app.ClientSecret,
	// 		Endpoint: oauth2.Endpoint{
	// 			AuthURL:   "https://jeanbon.mycozy.cloud/auth/authorize",
	// 			TokenURL:  "https://jeanbon.mycozy.cloud/auth/access_token",
	// 			AuthStyle: oauth2.AuthStyleAutoDetect,
	// 		},
	// 		RedirectURL: "http://localhost:9090",
	// 		Scopes:      []string{"io.cozy.files"},
	// 	}

	// 	tok, err := reuseToken()
	// 	if tok == nil {
	// 		tok, err = authenticate(ctx, &cfg)
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	client := cfg.Client(ctx, tok)

	// 	res, err := client.Get("https://jeanbon.mycozy.cloud/files/io.cozy.files.root-dir")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer res.Body.Close()

	// 	raw, _ := io.ReadAll(res.Body)

	// 	fmt.Printf("everything is ok: %s\n", string(raw))
	// }

	// func getAppInfo(ctx context.Context) (*Client, error) {
	// 	req, _ := http.NewRequest(http.MethodGet, "https://jeanbon.mycozy.cloud/auth/register/f823c0dfd772df110803b3253160a803", nil)

	// 	req.Header.Add("authorization", "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJyZWdpc3RyYXRpb24iLCJpYXQiOjE2ODE4ODU3NDQsImlzcyI6ImplYW5ib24ubXljb3p5LmNsb3VkIiwic3ViIjoiZjgyM2MwZGZkNzcyZGYxMTA4MDNiMzI1MzE2MGE4MDMifQ.bx3JLvfweh3kbn1c5DSdPU3Q7FnMEvJTqxv0m19M213-1cpjbZNQNitxLi1TphfVMulclU5B5Tf-o4x_ip0JiA")
	// 	req.Header.Add("accept", "application/json")
	// 	req.Header.Add("content-type", "application/json")

	// 	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	// 	if err != nil {
	// 		return nil, fmt.Errorf("fail to get the app infos: %w", err)
	// 	}
	// 	defer res.Body.Close()

	// 	var resBody Client
	// 	err = json.NewDecoder(res.Body).Decode(&resBody)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to read the response body: %w", err)
	// 	}

	// 	return &resBody, nil
	// }

	// func reuseToken() (*oauth2.Token, error) {
	// 	fileName, err := xdg.SearchStateFile(oauthFilePath)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	raw, err := os.ReadFile(fileName)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to read the content of %q: %w", fileName, err)
	// 	}

	// 	var tok oauth2.Token
	// 	err = json.Unmarshal(raw, &tok)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("invalid content for %q: %w", fileName, err)
	// 	}

	// 	return &tok, nil
	// }

	// func authenticate(ctx context.Context, cfg *oauth2.Config) (*oauth2.Token, error) {
	// 	// Create a state
	// 	b := make([]byte, 32)
	// 	if _, err := io.ReadFull(rand.Reader, b); err != nil {
	// 		return nil, fmt.Errorf("failed to create the state: %w", err)
	// 	}

	// 	state := base64.StdEncoding.EncodeToString(b)

	// 	// Redirect user to consent page to ask for permission
	// 	// for the scopes specified above.
	// 	url := cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
	// 	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)

	// 	// Use the authorization code that is pushed to the redirect
	// 	// URL. Exchange will do the handshake to retrieve the
	// 	// initial access token. The HTTP Client returned by
	// 	// conf.Client will refresh the token as necessary.
	// 	var code string
	// 	fmt.Printf("Please enter the code: ")
	// 	if _, err := fmt.Scan(&code); err != nil {
	// 		return nil, fmt.Errorf("failed to retrieve the code")
	// 	}

	// 	tok, err := cfg.Exchange(ctx, code)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fileName, err := xdg.StateFile(oauthFilePath)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to save the oauth token into a file: %w", err)
	// 	}

	// 	rawToken, err := json.Marshal(tok)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to marshal the token: %w", err)
	// 	}

	// 	err = os.WriteFile(fileName, rawToken, 0o600)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to save the token into %q: %w", fileName, err)
	// 	}

	// return tok, nil
}

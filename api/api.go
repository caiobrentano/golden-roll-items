package api

import (
	"encoding/json"
	// "ioutil"
	"fmt"
	"net/http"
	"github.com/dimfeld/httptreemux"
	"log"
)

var bungie_url = "https://www.bungie.net/Platform/Destiny/2/Stats/GetMembershipIdByDisplayName/%s/"

type UserInfoClient struct {
	Client *http.Client
}

type UserInfo struct {
	Response string
}

func (u *UserInfoClient) GetBungieMember(psnId string) (*UserInfo, error) {
	url := fmt.Sprintf(bungie_url, psnId)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("x-api-key", "ab1fdc2acddc4746a747fd29f0c1790c")

	// res, err := u.Client.Do(req)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	userInfo := &UserInfo{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(userInfo)

	defer res.Body.Close()
	return userInfo, err
}

type CreateDestinyUser struct{}

func (h *CreateDestinyUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())

	uiClient := UserInfoClient{Client: http.DefaultClient}
	u, err := uiClient.GetBungieMember(params["psnId"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User Bungie id: %s!", u.Response)
}

// type GetDestinyUser struct{}

// func (h *GetDestinyUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	params := httptreemux.ContextParams(r.Context())
// }

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/JamesStewy/sse"
)

// Claim is my awesome sruct that is public
type Claim struct {
	ID          string  `json:"id"`
	CompanyName string  `json:"companyName"`
	BatchDate   string  `json:"batchDate"`
	BilledAmt   float64 `json:"billedAmt"`
	Active      bool    `json:"active"`
	Changed     bool    `json:"changed"`
}

var claims []Claim
var clients map[*sse.Client]bool

func streamHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// Initialise (REQUIRED)
	client, err := sse.ClientInit(w)
	// Return error if unable to initialise Server-Sent Events
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Add client to external variable for later use
	clients[client] = true
	// Remove client from external variable on exit
	defer delete(clients, client)
	// Run Client (REQUIRED)
	client.Run(r.Context())
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getAllClaims(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(claims)
}

func updateClaim(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body) // []byte
	if err != nil {
		log.Println("Error:", err)
	}
	var claim Claim
	json.Unmarshal(body, &claim)

	for i := range claims {
		if claims[i].ID == claim.ID {
			claims[i].Active = claim.Active
		}
	}

	updatedClaim := fmt.Sprintf("{ \"id\": \"%s\", \"active\": %t }", claim.ID, claim.Active)
	msg := sse.Msg{
		Event: "claimUpdate",
		Data:  updatedClaim,
	}
	for client := range clients {
		client.Send(msg)
	}
}

func main() {
	claims = append(claims, Claim{ID: "162244ae-c0e6-4468-b9e6-b42abdec31f3", CompanyName: "Prabhupada", BatchDate: "2018-10-31T01:50:45Z", BilledAmt: 4020.03, Active: false})
	claims = append(claims, Claim{ID: "bcd64286-8209-4078-b5ad-3e87abdd88c4", CompanyName: "Yakidoo", BatchDate: "2017-12-12T06:29:54Z", BilledAmt: 2576.37, Active: false})
	claims = append(claims, Claim{ID: "fbe62e94-6c3d-4702-ac5b-4397b51452a7", CompanyName: "Pixope", BatchDate: "2018-05-07T00:39:54Z", BilledAmt: 1403.46, Active: false})
	claims = append(claims, Claim{ID: "adb2eb2e-c3e3-4611-8f82-3957b7b78570", CompanyName: "Fanoodle", BatchDate: "2018-04-18T03:06:49Z", BilledAmt: 358.05, Active: false})
	claims = append(claims, Claim{ID: "2903b93f-9e49-4b21-8ddb-3edaf6742866", CompanyName: "Thoughtmix", BatchDate: "2018-09-08T13:43:51Z", BilledAmt: 1529.08, Active: false})
	claims = append(claims, Claim{ID: "dd5f174e-5bd9-4b8b-b9fb-c86674f4e26e", CompanyName: "Zooxo", BatchDate: "2018-04-18T15:09:49Z", BilledAmt: 4081.89, Active: false})
	claims = append(claims, Claim{ID: "9add925e-8d0d-4239-8bd5-82e5dcf193e5", CompanyName: "Skyndu", BatchDate: "2018-05-18T23:28:16Z", BilledAmt: 2638.19, Active: false})
	claims = append(claims, Claim{ID: "8fe3ede2-e9d1-41e3-ad80-6f04a1856599", CompanyName: "Bluejam", BatchDate: "2018-01-24T16:21:56Z", BilledAmt: 1171.71, Active: false})
	claims = append(claims, Claim{ID: "070b3737-4bef-475b-926a-6d4747aee0aa", CompanyName: "Quinu", BatchDate: "2017-12-01T17:13:02Z", BilledAmt: 4178.14, Active: false})
	claims = append(claims, Claim{ID: "527bfe85-783b-4fb4-91b5-4c572df8aaf3", CompanyName: "Edgewire", BatchDate: "2018-09-13T22:24:12Z", BilledAmt: 4099.26, Active: false})
	claims = append(claims, Claim{ID: "ea29f25b-234f-4a6a-84d3-e59b69472ca9", CompanyName: "Eire", BatchDate: "2017-12-09T15:20:42Z", BilledAmt: 2046.13, Active: false})
	claims = append(claims, Claim{ID: "35f8ed6e-2e5a-4098-bad6-f057c2266572", CompanyName: "Wikizz", BatchDate: "2017-12-26T10:53:22Z", BilledAmt: 204.89, Active: false})
	claims = append(claims, Claim{ID: "f3eb9568-175e-4590-98e2-8c2bca3112bc", CompanyName: "Viva", BatchDate: "2017-12-26T15:49:06Z", BilledAmt: 1356.67, Active: false})
	claims = append(claims, Claim{ID: "da5cdb27-a600-4b2d-8bc6-43fa33842a80", CompanyName: "Kamba", BatchDate: "2017-12-26T10:59:36Z", BilledAmt: 3035.09, Active: false})
	claims = append(claims, Claim{ID: "8044dd0e-2274-4b06-955a-b2b992a47dbd", CompanyName: "Jamia", BatchDate: "2018-07-05T09:14:10Z", BilledAmt: 2811.15, Active: false})
	claims = append(claims, Claim{ID: "5c9f308d-84f4-4bc0-9ed1-624d289b04f6", CompanyName: "Kamba", BatchDate: "2018-02-08T15:36:57Z", BilledAmt: 4171.05, Active: false})
	claims = append(claims, Claim{ID: "dbfbb686-ae59-4e35-a3e5-6bcd35954278", CompanyName: "Photobean", BatchDate: "2018-10-12T01:05:26Z", BilledAmt: 1210.49, Active: false})
	claims = append(claims, Claim{ID: "8376b5c7-5c26-4165-880c-67ef48a754b8", CompanyName: "Realfire", BatchDate: "2018-07-06T12:19:15Z", BilledAmt: 2579.26, Active: false})
	claims = append(claims, Claim{ID: "75d75870-ec57-4da0-a918-45e4e24f5e37", CompanyName: "Realmix", BatchDate: "2018-02-05T14:55:27Z", BilledAmt: 1476.42, Active: false})
	claims = append(claims, Claim{ID: "995b37e8-1800-49ec-94a6-97c49ca22262", CompanyName: "Flipbug", BatchDate: "2018-05-23T03:33:54Z", BilledAmt: 4573.01, Active: false})

	clients = make(map[*sse.Client]bool)

	r := false
	ticker := time.NewTicker(time.Second * 29)
	go func() {
		for range ticker.C {
			msg := sse.Msg{
				Event: "ping",
			}
			if r == true {
				i := rand.Intn(len(claims))
				tc := claims[i] // TargetClaim
				claims[i].Active = !tc.Active
				updatedClaim := fmt.Sprintf("{ \"id\": \"%s\", \"active\": %t }", tc.ID, tc.Active)
				msg = sse.Msg{
					Event: "claimUpdate",
					Data:  updatedClaim,
				}
			}
			for client := range clients {
				client.Send(msg)
			}
		}
	}()

	http.HandleFunc("/api/claims", getAllClaims)
	http.HandleFunc("/api/update-claim", updateClaim)
	http.HandleFunc("/api/stream/", streamHandler)

	log.Printf("Listening on 8000...")
	http.ListenAndServe(":8000", nil)
}

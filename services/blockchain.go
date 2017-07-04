package Services

import (
	"blockchain/blockchain"
	"encoding/json"
	"log"
	"net/http"
)

//RestBlock structure d'entrée du service REST
type RestBlock struct {
	Data string `json:"data"`
}

//CreateBlock Crée et rajoute un block dans la blockChain
func CreateBlock(w http.ResponseWriter, r *http.Request) {
	Option(w, r)

	blockLu := new(RestBlock)
	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&blockLu)

	if error != nil {
		log.Println(error.Error())
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	// On ajout un block a la block chain
	blockchain.AddDataBlock(blockLu.Data)

}

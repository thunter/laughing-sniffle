package main

import (
	"fmt"
	"encoding/json"
)

type KillResponse struct {
	Package KillPackage `json:"package"`
}

type KillPackage struct {
	KillId int `json:"killID"`
	Killmail KillMail `json:"killmail"`
	Zkb KillZkillBoard `json:"zkb"`
}

type KillMail struct {
	Attackers []KillAttackers `json:"attackers"`
	KillmailId int `json:"killmail_id"`
	KillmailTime string `json:"killmail_time"`
	SolarSystemId int `json:"solar_system_id"`
	Victim KillVictim `json:"victim"`
}

type KillAttackers struct {
	AllianceId int `json:"alliance_id"`
	CharacterId int `json:"character_id"`
	CorporationId int `json:"corporation_id"`
	DamageDone int `json:"damage_done"`
	FinalBlow bool `json:"final_blow"`
	SecurityStatus float32 `json:"security_status"`
	ShipTypeId int `json:"ship_type_id"`
	WeaponTypeId int `json:"weapon_type_id"`
}

type KillVictim struct {
	AllianceId int `json:"alliance_id"`
	CharacterId int `json:"character_id"`
	CorporationId int `json:"corporation_id"`
	DamageTaken int `json:"damage_taken"`
	ShipTypeId int `json:"ship_type_id"`
	Items []KillItems `json:"items"`
	Position KillPosition `json:"position"`
}

type KillItems struct {
	Flag int `json:"flag"`
	ItemTypeId int `json:"item_type_id"`
	QuantityDropped int `json:"quantity_dropped"`
	Singleton int `json:"singleton"`
}

type KillPosition struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type KillZkillBoard struct {
	LocationID int `json:"locationID"`
	Hash string `json:"hash"`
	FittedValue float32 `json:"fittedValue"`
	TotalValue float32 `json:"totalValue"`
	Points int `json:"points"`
	Npc bool `json:"npc"`
	Solo bool `json:"solo"`
	Awox bool `json:"awox"`
	Href string `json:"href"`
}

func parseKillMail(body []byte) (*KillResponse, error)  {
	var kill = new(KillResponse)
	err := json.Unmarshal(body, &kill)
	if err != nil {
		fmt.Println("Error parshing JSON:", err)
	}

	return kill, err
}
package main

import (
	"log"
	"encoding/json"
)

type KillResponse struct {
	Package KillPackage
}

type KillPackage struct {
	KillID int
	Killmail KillMail
	Zkb KillZkillBoard
}

type KillMail struct {
	Attackers []KillAttackers
	Killmail_id int
	Killmail_time string
	Solar_system_id int
	Victim KillVictim
}

type KillAttackers struct {
	Alliance_id int
	Character_id int
	Corporation_id int
	Damage_done int
	Final_blow bool
	Security_status float32
	Ship_type_id int
	Weapon_type_id int
}

type KillVictim struct {
	Alliance_id int
	Character_id int
	Corporation_id int
	Damage_taken int
	Ship_type_id int
	Items []KillItems
}

type KillItems struct {
	Flag int
	Item_type_id int
	Quantity_dropped int
	singleton int
}

type KillPosition struct {
	x float32
	y float32
	z float32
}

type KillZkillBoard struct {
	LocationID int
	Hash string
	FittedValue float32
	totalValue float32
	points int
	npc bool
	solo bool
	awox bool
	href string
}

func parseKillMail(killMail string) KillResponse  {
	var kill KillResponse
	if err := json.Unmarshal([]byte(killMail), &kill); err != nil {
		log.Fatal(err)
	}
	return kill
}
package main

import (
	"log"
	"testing"
)

const sample1  = `{"package":{"killID":72770296,"killmail":{"attackers":[{"damage_done":58188,"faction_id":500011,"final_blow":true,"security_status":0,"ship_type_id":24130}],"killmail_id":72770296,"killmail_time":"2018-10-06T07:55:56Z","solar_system_id":30000581,"victim":{"alliance_id":99003581,"character_id":2113673312,"corporation_id":98538918,"damage_taken":58188,"items":[{"flag":13,"item_type_id":4405,"quantity_dropped":1,"singleton":0},{"flag":16,"item_type_id":33824,"quantity_destroyed":1,"singleton":0},{"flag":27,"item_type_id":23527,"quantity_dropped":1,"singleton":0},{"flag":22,"item_type_id":3841,"quantity_destroyed":1,"singleton":0},{"flag":12,"item_type_id":4405,"quantity_destroyed":1,"singleton":0},{"flag":19,"item_type_id":5955,"quantity_dropped":1,"singleton":0},{"flag":20,"item_type_id":2281,"quantity_dropped":1,"singleton":0},{"flag":21,"item_type_id":3841,"quantity_destroyed":1,"singleton":0},{"flag":92,"item_type_id":31790,"quantity_destroyed":1,"singleton":0},{"flag":15,"item_type_id":4405,"quantity_destroyed":1,"singleton":0},{"flag":93,"item_type_id":31790,"quantity_destroyed":1,"singleton":0},{"flag":11,"item_type_id":1447,"quantity_dropped":1,"singleton":0},{"flag":14,"item_type_id":4405,"quantity_destroyed":1,"singleton":0},{"flag":87,"item_type_id":2488,"quantity_destroyed":7,"singleton":0},{"flag":87,"item_type_id":2488,"quantity_dropped":3,"singleton":0},{"flag":94,"item_type_id":31790,"quantity_destroyed":1,"singleton":0}],"position":{"x":-218152289631.39032,"y":-12295863249.963987,"z":307341589050.7495},"ship_type_id":17843}},"zkb":{"locationID":40036389,"hash":"339eca8dd66ef644f4d5ee50e2fbe9ae99aa3c28","fittedValue":71457063.77,"totalValue":83219205.4,"points":1,"npc":true,"solo":false,"awox":false,"href":"https://esi.evetech.net/v1/killmails/72770296/339eca8dd66ef644f4d5ee50e2fbe9ae99aa3c28/"}}}`

func TestJSONParsing(t *testing.T) {
	kill, err := parseKillMail([]byte(sample1))

	if err != nil {
		log.Fatal(err)
	}
	
	if kill.Package.KillId != 72770296 {
		t.Errorf("KillID was incorrect, got: %d", kill.Package.KillId)
	}

	if len(kill.Package.Killmail.Attackers) != 1 {
		t.Errorf("Attackers not parsed correctly, got: %d", len(kill.Package.Killmail.Attackers))
	}

	if len(kill.Package.Killmail.Victim.Items) != 16 {
		t.Errorf("Items not parsed correctly, got %d", len(kill.Package.Killmail.Victim.Items))
	}



}

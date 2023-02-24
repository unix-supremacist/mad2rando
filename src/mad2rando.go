package main

import (
	//"fmt"
    "log"
    "os"
	//"github.com/bigkevmcd/go-configparser"
)

func main(){
	file, err := os.Create("spoiler.log")
	if err != nil {
        log.Fatal(err)
    }
	defer file.Close()

	SpoilerLogger := log.New(file, "SPOILER: ", 0)

	type Level struct {
		accessibleLevels []string
		hasCoins bool
		saneToRandomize bool
		monkeys int
		climbMonkeys int
		isMinigame bool
		breaksInMinigame bool
		boring bool
	}

	f := false;
	t := true;

	levels := map[string]Level{
		//accessible levels, hasCoins, saneToRandomize, monkeys, climbmonkeys, boring
		"animal_chess": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"Card_Match_Game": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"Credits": Level{[]string{"title", "map"}, f, t, 0, 0, f, f, f},
		"Minigame_Diving_Location_Menu": Level{[]string{"title", "DivingLocation_IslandFever", "DivingLocation_PrepareToLaunch", "DivingLocation_RitesOfPassage", "DivingLocation_Waterhole"}, f, t, 0, 0, t, f, f},
		"DivingLocation_IslandFever": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"DivingLocation_PrepareToLaunch": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"DivingLocation_RitesOfPassage": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"DivingLocation_Waterhole": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"DutyFree": Level{[]string{"title", "map"}, f, t, 0, 0, f, f, f},
		"golf_cake": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"golf_lovelylumps": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"golf_targetTree": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"golf_3holes": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"golf_baggagecheck": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"golf_cratercrossing": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"golf_crossedpaths": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"golf_foosaBall": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"golf_junkyard": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"golf_maze": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"Golf_minigame": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"golf_ravenousRhinos": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"HungryHippo": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"Minigame_HotDurian": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"RoP_MusicalChairs": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"Soccer": Level{[]string{"title"}, f, t, 0, 0, t, f, f},
		"map": Level{[]string{"title", "DutyFree", "IslandFever"}, f, f, 0, 0, f, f, f},
		"title": Level{[]string{"DutyFree", "IslandFever"}, f, f, 0, 0, f, f, f},
		"Dam_Busters": Level{[]string{"Credits""title", "map"}, f, f, 0, 0, f, f, f},
		"FixThePlane": Level{[]string{"Waterhole", "Dam_Busters", "title", "map"}, f, f, 0, 0, f, t, f},
		"Waterhole": Level{[]string{"BraveNewWild", "FixThePlane", "DrMelman", "MartyRace", "RitesOfPassage", "VolcanoRave", "Watercaves", "ConvoyChase", "Morts_Adventure"}, f, f, 0, 0, f, f, f},
		"penguins": Level{[]string{"penguins2", "title", "map"}, f, t, 0, 0, f, f, f},
		"penguins2": Level{[]string{"ConvoyChase", "title", "map"}, f, t, 0, 0, f, f, f},
		"Prepare2Launch_Plane": Level{[]string{"BraveNewWild", "title", "map"}, f, t, 0, 0, f, f, f},
		"Prepare2Launch": Level{[]string{"Prepare2Launch_Plane", "title", "map"}, f, t, 0, 0, f, f, f},
		"BraveNewWild": Level{[]string{"Waterhole", "title", "map"}, f, t, 0, 0, f, f, f},
		"ConvoyChase": Level{[]string{"Waterhole", "title", "map"}, f, t, 0, 0, f, f, f},
		"DrMelman": Level{[]string{"Waterhole", "title", "map"}, f, t, 0, 0, f, f, f},
		"MartyRace": Level{[]string{"Waterhole", "title", "map"}, f, t, 0, 0, f, f, f},
		"RitesOfPassage": Level{[]string{"Waterhole", "title", "map"}, f, t, 0, 0, f, f, f},
		"VolcanoRave": Level{[]string{"title", "map"}, f, t, 0, 0, f, f, f},
		"Watercaves": Level{[]string{"Waterhole", "Wooing_Gloria", "title", "map"}, f, t, 0, 0, f, f, f},
		"Wooing_Gloria": Level{[]string{"Waterhole", "title", "map"}, f, t, 0, 0, f, f, f},
		"IslandFever": Level{[]string{"Prepare2Launch", "title", "map"}, f, t, 0, 0, f, f, f},
		"Morts_Adventure": Level{[]string{"Waterhole", "title", "map"}, f, t, 0, 0, f, f, f},
	}

	type LevelSlot struct {
		reqCoins bool
		reqMonkeys bool
		reqClimb bool
		isMinigame bool
		reqWooingGloria bool
	}

	levelslots := map[string]LevelSlot{
		"animal_chess": LevelSlot{f, f, f, t, f},
		"Card_Match_Game": LevelSlot{f, f, f, t, f},
		"Credits": LevelSlot{f, f, f, f, f},
		"Minigame_Diving_Location_Menu": LevelSlot{f, f, f, t, f},
		"DivingLocation_IslandFever": LevelSlot{f, f, f, t, f},
		"DivingLocation_PrepareToLaunch": LevelSlot{t, f, f, t, f},
		"DivingLocation_RitesOfPassage": LevelSlot{t, f, f, t, f},
		"DivingLocation_Waterhole": LevelSlot{f, f, f, t, f},
		"DutyFree": LevelSlot{f, f, f, t, f},
		"golf_cake": LevelSlot{t, f, f, t, f},
		"golf_lovelylumps": LevelSlot{t, f, f, t, f},
		"golf_3holes": LevelSlot{f, f, f, t, f},
		"golf_baggagecheck": LevelSlot{f, f, f, t, f},
		"golf_cratercrossing": LevelSlot{f, f, f, t, f},
		"golf_crossedpaths": LevelSlot{f, f, f, t, f},
		"golf_foosaBall": LevelSlot{f, f, f, t, f},
		"golf_junkyard": LevelSlot{f, f, f, t, f},
		"golf_maze": LevelSlot{f, f, f, t, f},
		"Golf_minigame": LevelSlot{f, f, f, t, f},
		"golf_ravenousRhinos": LevelSlot{f, f, f, t, f},
		"HungryHippo": LevelSlot{f, f, f, t, f},
		"Minigame_HotDurian": LevelSlot{f, f, f, t, f},
		"RoP_MusicalChairs": LevelSlot{f, f, f, t, f},
		"Soccer": LevelSlot{f, f, f, t, f},
		"map": LevelSlot{f, f, f, f, f},
		"title": LevelSlot{f, f, f, f, f},
		"Dam_Busters": LevelSlot{f, t, f, f, f},
		"FixThePlane": LevelSlot{f, f, f, f, f},
		"Waterhole": LevelSlot{f, f, f, f, f},
		"penguins": LevelSlot{f, f, f, f, f},
		"penguins2": LevelSlot{f, f, f, f, f},
		"Prepare2Launch_Plane": LevelSlot{f, f, f, f, f},
		"Prepare2Launch": LevelSlot{f, f, f, f, f},
		"BraveNewWild": LevelSlot{f, f, f, f, f},
		"ConvoyChase": LevelSlot{f, f, f, f, f},
		"DrMelman": LevelSlot{f, f, f, f, f},
		"MartyRace": LevelSlot{f, f, f, f, f},
		"RitesOfPassage": LevelSlot{f, f, f, f, f},
		"VolcanoRave": LevelSlot{f, f, f, f, t},
		"Watercaves": LevelSlot{f, f, t, f, f},
		"Wooing_Gloria": LevelSlot{f, f, f, f, f},
		"IslandFever": LevelSlot{f, f, f, f, f},
		"Morts_Adventure": LevelSlot{f, f, f, f, f},
	}

	//deadend_pool := []string{
	//	"animal_chess", "Card_Match_Game", "Credits",
	//	"DivingLocation_IslandFever", "DivingLocation_PrepareToLaunch",
	//	"DivingLocation_RitesOfPassage", "DivingLocation_Waterhole",
	//	"DutyFree", "golf_cake", "golf_lovelylumps", "golf_targetTree",
	//	"HungryHippo", "Minigame_HotDurian", "RoP_MusicalChairs", "Soccer",
	//}

	//minigame_levels := []string{
	//	"animal_chess", "Card_Match_Game",
	//	"DivingLocation_IslandFever", "DivingLocation_PrepareToLaunch",
	//	"DivingLocation_RitesOfPassage", "DivingLocation_Waterhole",
	//	"Minigame_Diving_Location_Menu", "HungryHippo", "Minigame_HotDurian",
	//	"RoP_MusicalChairs", "Soccer",
	//	"golf_3holes", "golf_baggagecheck",
	//	"golf_cratercrossing", "golf_crossedpaths", "golf_foosaBall",
	//	"golf_junkyard",  "golf_maze",
	//	"Golf_minigame", "golf_ravenousRhinos", "golf_cake",
	//	"golf_lovelylumps", "golf_targetTree",
	//}

	//SpoilerLogger.Println(levels["animal_chess"])

	for k, v := range levels {
		SpoilerLogger.Println(k, v)
	}
	
}
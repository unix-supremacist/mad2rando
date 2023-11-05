package main

import (
	"fmt"
    "log"
    "os"
	"io/ioutil"
	"math/rand"
	"time"
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	cp "github.com/otiai10/copy"
	//"sort"
	//"github.com/bigkevmcd/go-configparser"
)

var SpoilerLogger *log.Logger
var seed int
var level_groups [][]string
var dam_minigame bool

func main(){
	statusLabel := widget.NewLabel("Waiting to randomize.")

    _, err := os.Stat("ContentOG")
    if os.IsNotExist(err) {
        statusLabel.SetText("ContentOG Folder does not exist, creating it.")
		_, err := os.Stat("Content")
		if !os.IsNotExist(err) {
			eror(os.Mkdir("ContentOG", os.ModePerm))
			cp.Copy("Content", "ContentOG")
			statusLabel.SetText("Done")
		} else {
			statusLabel.SetText("Content folder not found, is the rando installed in the game directory?")
		}
    }
	
	file, err := os.Create("spoiler.log")
	if err != nil {
        log.Fatal(err)
    }
	defer file.Close()
	SpoilerLogger = log.New(file, "SPOILER: ", 0)

	initializeLevels()

	rApp := app.New()
	rWindow := rApp.NewWindow("E2A: Randomizer")
	rWindow.Show()

	randomizeSeed()
	seedLabel := widget.NewLabel(strconv.Itoa(seed))
	randomizeSeedButton := widget.NewButton("Randomize Seed", func() {
		statusLabel.SetText("Randomizing Seed")
		randomizeSeed()
		seedLabel.SetText(strconv.Itoa(seed))
		statusLabel.SetText("Done")
	})

	randomizeLevelButton := widget.NewButton("Randomize Levels", func() {
		statusLabel.SetText("Randomizing Levels")
		randomizeLevels()
		statusLabel.SetText("Done")
	})

	damMinigameButton := widget.NewButton("Dam Busters Minigame: false", func() {
		if dam_minigame {
			dam_minigame = false
			level_groups[4] = pop(level_groups[4], "Dam_Busters")
		} else {
			dam_minigame = true
			level_groups[4] = append(level_groups[4], "Dam_Busters")
		}
	})

	go func() {
		for range time.Tick(time.Millisecond * 250) {
			if dam_minigame {
				damMinigameButton.SetText("Dam Busters Minigame: true")
			} else {
				damMinigameButton.SetText("Dam Busters Minigame: false")
			}
		}
	}()

	rContent := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), statusLabel, layout.NewSpacer(), seedLabel, layout.NewSpacer(), randomizeSeedButton, layout.NewSpacer(), randomizeLevelButton, layout.NewSpacer(), damMinigameButton)

	rWindow.SetContent(rContent)
	

	rWindow.Resize(fyne.NewSize(100, 100))
	rApp.Run()
}

func initializeLevels(){
	main_levels := []string{"BraveNewWild", "ConvoyChase", "Credits", "DrMelman", "DutyFree", "FixThePlane", "MartyRace", "Morts_Adventure", "RitesOfPassage", "VolcanoRave", "Wooing_Gloria"}
	string_levels := []string{"IslandFever", "penguins", "penguins2", "Prepare2Launch", "Prepare2Launch_Plane", "Watercaves"}
	hub_levels := []string{"map", "title", "Waterhole"}
	minigame_levels := []string{"animal_chess", "Card_Match_Game", "DivingLocation_IslandFever", "DivingLocation_PrepareToLaunch", "DivingLocation_RitesOfPassage", "DivingLocation_Waterhole", "golf_targetTree", "HungryHippo", "Minigame_HotDurian", "RoP_MusicalChairs", "Soccer"}
	minigame_string_levels := []string{"golf_3holes", "golf_baggagecheck", "golf_cake", "golf_cratercrossing", "golf_crossedpaths", "golf_foosaBall", "golf_junkyard", "golf_lovelylumps", "golf_maze", "Golf_minigame", "golf_ravenousRhinos", "Minigame_Diving_Location_Menu"}

	level_groups = [][]string{main_levels, string_levels, hub_levels, minigame_levels, minigame_string_levels}
}

func randomizeSeed(){
	rand.Seed(time.Now().UnixNano())
	seed = rand.Intn(1000000000)
	SpoilerLogger.Println(seed)
	rand.Seed(int64(seed))
}

func shuffleLevels(levels_slots, levels []string) []string {
	copy(levels_slots, levels)
	rand.Shuffle(len(levels_slots), func(i, j int) { levels_slots[i], levels_slots[j] = levels_slots[j], levels_slots[i] })
	if contains(levels_slots, "VolcanoRave") && contains(levels_slots, "Wooing_Gloria"){
		i := 0
		for levels_slots[i] != "VolcanoRave" {
			i++
		}
		if levels[i] == "Wooing_Gloria" {
			s := "this is a rare seed where wooing gloria manages to get sorted to wooing gloria, please send me the seed and any issues you have since i can't easily test this edge case without it"
			SpoilerLogger.Println(s)
			fmt.Println(s)
			return shuffleLevels(levels_slots, levels)
		}
	}

	return levels_slots
}

func randomizeLevels(){
	for group := range level_groups {
		levels_slots := make([]string, len(level_groups[group]))
		levels_slots = shuffleLevels(levels_slots, level_groups[group])
		

		for level := range level_groups[group] {
			copyFile("ContentOG/Streams/win/"+level_groups[group][level]+".arc", "Content/Streams/win/"+levels_slots[level]+".arc")
			copyFile("ContentOG/Streams/win/"+level_groups[group][level]+".bld", "Content/Streams/win/"+levels_slots[level]+".bld")
			SpoilerLogger.Println(level_groups[group][level]+" to "+levels_slots[level])
		}
	}
}

func eror(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func copyFile(path string, out string) {
	in, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	eror(ioutil.WriteFile(out, in, 0644))
}

func pop(slice []string, toPop string) []string {
	i := 0
	for slice[i] != toPop {
		i++
	}
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}

func contains(slice []string, str string) bool {
    for _, str2 := range slice {
        if str == str2 {
            return true
        }
    }
    return false
}
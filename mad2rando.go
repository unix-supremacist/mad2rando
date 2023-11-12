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

type Level struct {
	Name string
	HasCoins, HasMonkeys, IsStringLevel, IsFirstStringLevel, IsHubLevel, IsMinigame, BreaksMinigame bool
}

var (
	SpoilerLogger *log.Logger
	dam_minigame, minigames, amac, sanity bool
	levels []Level
	level_groups [][]string
	seed int
	breaksminigames []string
	isminigame []string
)

func initializeLevels(){
	main_levels := []string{}
	main_levels_without_coins := []string{}
	first_string_levels := []string{}
	string_levels := []string{}
	hub_levels := []string{}
	minigame_levels := []string{}
	minigame_string_levels := []string{}
	no_rando := []string{}
	breaksminigames = []string{}
	isminigame = []string{}

	for i := range levels {
		if sanity {
			main_levels = append(main_levels, levels[i].Name)
		} else {
			if levels[i].Name == "Dam_Busters" {
				if dam_minigame {
					minigame_levels = append(minigame_levels, levels[i].Name)
				} else {
					no_rando = append(no_rando, levels[i].Name)
				}
			} else if levels[i].IsMinigame && !minigames {
				if levels[i].IsStringLevel {
					minigame_string_levels = append(minigame_string_levels, levels[i].Name)
				} else {
					minigame_levels = append(minigame_levels, levels[i].Name)
				}
			} else {
				if levels[i].IsFirstStringLevel {
					first_string_levels = append(first_string_levels, levels[i].Name)
				} else if levels[i].IsStringLevel {
					string_levels = append(string_levels, levels[i].Name)
				} else if levels[i].IsHubLevel {
					if !amac {
						hub_levels = append(hub_levels, levels[i].Name)
					} else {
						no_rando = append(no_rando, levels[i].Name)
					}
				} else {
					if amac && !levels[i].HasCoins {
						main_levels_without_coins = append(main_levels_without_coins, levels[i].Name)
					} else {
						main_levels = append(main_levels, levels[i].Name)
					}
				}
			}
		}

		if levels[i].BreaksMinigame {
			breaksminigames = append(breaksminigames, levels[i].Name)
		}
		if levels[i].IsMinigame {
			isminigame = append(isminigame, levels[i].Name)
		}
	}

	SpoilerLogger.Println(main_levels)
	SpoilerLogger.Println(main_levels_without_coins)
	SpoilerLogger.Println(first_string_levels)
	SpoilerLogger.Println(string_levels)
	SpoilerLogger.Println(hub_levels)
	SpoilerLogger.Println(minigame_levels)
	SpoilerLogger.Println(minigame_string_levels)
	SpoilerLogger.Println(no_rando)

	level_groups = [][]string{main_levels, main_levels_without_coins, first_string_levels, string_levels, hub_levels, minigame_levels, minigame_string_levels, no_rando}
}

func main(){
	minigames = true
	t := true
	f := false

	levels = []Level{
		Level{"animal_chess", f, f, f, f, f, t, f},
		Level{"BraveNewWild", t, t, f, f, f, f, f},
		Level{"Card_Match_Game", f, f, f, f, f, t, f},
		Level{"ConvoyChase", t, t, f, f, f, f, f},
		Level{"Credits", f, f, f, f, f, f, f},
		Level{"Dam_Busters", t, f, f, f, f, f, f},
		Level{"DivingLocation_IslandFever", f, f, f, f, f, t, f},
		Level{"DivingLocation_PrepareToLaunch", f, f, f, f, f, t, f},
		Level{"DivingLocation_RitesOfPassage", f, f, f, f, f, t, f},
		Level{"DivingLocation_Waterhole", f, f, f, f, f, t, f},
		Level{"DrMelman", t, t, f, f, f, f, t},
		Level{"DutyFree", f, f, f, f, f, f, f},
		Level{"FixThePlane", t, f, f, f, f, f, f},
		Level{"golf_3holes", f, f, t, f, f, t, f},
		Level{"golf_baggagecheck", f, f, t, f, f, t, f},
		Level{"golf_cake", f, f, t, f, f, t, f},
		Level{"golf_cratercrossing", f, f, t, f, f, t, f},
		Level{"golf_crossedpaths", f, f, t, f, f, t, f},
		Level{"golf_foosaBall", f, f, t, f, f, t, f},
		Level{"golf_junkyard", f, f, t, f, f, t, f},
		Level{"golf_lovelylumps", f, f, t, f, f, t, f},
		Level{"golf_maze", f, f, t, f, f, t, f},
		Level{"Golf_minigame", f, f, t, f, f, t, f},
		Level{"golf_ravenousRhinos", f, f, t, f, f, t, f},
		Level{"golf_targetTree", f, f, f, f, f, t, f},
		Level{"HungryHippo", f, f, f, f, f, t, f},
		Level{"IslandFever", t, f, t, t, f, f, t},
		Level{"map", f, f, f, f, t, f, f},
		Level{"MartyRace", t, t, f, f, f, f, f},
		Level{"Minigame_Diving_Location_Menu", f, f, t, f, f, t, f},
		Level{"Minigame_HotDurian", f, f, f, f, f, t, f},
		Level{"Morts_Adventure", t, f, f, f, f, f, f},
		Level{"penguins", t, f, t, t, f, f, f},
		Level{"penguins2", t, f, t, f, f, f, f},
		Level{"Prepare2Launch_Plane", t, f, t, f, f, f, f},
		Level{"Prepare2Launch", t, f, t, f, f, f, f},
		Level{"RitesOfPassage", t, t, f, f, f, f, t},
		Level{"RoP_MusicalChairs", f, f, f, f, f, t, f},
		Level{"Soccer", f, f, f, f, f, t, f},
		Level{"title", f, f, f, f, t, f, f},
		Level{"VolcanoRave", t, t, f, f, f, f, t},
		Level{"Watercaves", t, t, t, t, f, f, f},
		Level{"Waterhole", t, t, f, f, t, f, f},
		Level{"Wooing_Gloria", t, t, f, f, f, f, f},
	}

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
	eror(err)
	defer file.Close()
	SpoilerLogger = log.New(file, "SPOILER: ", 0)

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

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	damMinigameButton := widget.NewButton("Dam Busters Minigame: false", newToggle(&dam_minigame))
	minigameButton := widget.NewButton("Minigames", newToggle(&minigames))
	amacButton := widget.NewButton("All Monkeys All Coins", newToggle(&amac))
	sanityButton := widget.NewButton("Are You Insane", newToggle(&sanity))

	go func() {
		for range time.Tick(time.Millisecond * 200) {
			if input.Text != "" {
				i, err := strconv.Atoi(input.Text)
				if err != nil {
					eror(err)
				} else {
					if seed != i {
						seed = i
						seedLabel.SetText(input.Text)
						SpoilerLogger.Println(seed)
						rand.Seed(int64(seed))
					}
				}
			}

			newUiToggle("Dam Busters Minigame", &dam_minigame, damMinigameButton)
			newUiToggle("Minigames", &minigames, minigameButton)
			newUiToggle("All Monkeys All Coins", &amac, amacButton)
			newUiToggle("Are You Insane", &sanity, sanityButton)
		}
	}()

	rContent := container.New(
		layout.NewVBoxLayout(), layout.NewSpacer(),
		statusLabel, layout.NewSpacer(),
		seedLabel, layout.NewSpacer(),
		input, layout.NewSpacer(),
		randomizeSeedButton, layout.NewSpacer(),
		randomizeLevelButton, layout.NewSpacer(),
		minigameButton, layout.NewSpacer(),
		sanityButton, layout.NewSpacer(),
		amacButton, layout.NewSpacer(),
		damMinigameButton)

	rWindow.SetContent(rContent)
	rWindow.Resize(fyne.NewSize(100, 100))
	rApp.Run()
}

func newUiToggle(name string, b *bool, widget *widget.Button){
	if *b {
		widget.SetText(name+": true")
	} else {
		widget.SetText(name+": false")
	}
}

func newToggle(b *bool) func() {
	return func() {
		if *b {
			*b = false
		} else {
			*b = true
		}
	}
}

func randomizeSeed(){
	rand.Seed(time.Now().UnixNano())
	seed = rand.Intn(1000000000)
	SpoilerLogger.Println(seed)
	rand.Seed(int64(seed))
}

func shuffleLevels(levels_slots, levelsToShuffle []string) []string {
	copy(levels_slots, levelsToShuffle)
	rand.Shuffle(len(levels_slots), func(i, j int) { levels_slots[i], levels_slots[j] = levels_slots[j], levels_slots[i] })

	if !sanity {
		for l := range isminigame {
			if contains(levels_slots, isminigame[l]) {
				i := 0
				for levels_slots[i] != isminigame[l] {
					i++
				}
				if contains(breaksminigames, levelsToShuffle[i]) {
					s := "minigame breaking level in minigame, reshuffling"
					SpoilerLogger.Println(s)
					fmt.Println(s)
					return shuffleLevels(levels_slots, levelsToShuffle)
				}
			}
		}
	
		if contains(levels_slots, "VolcanoRave") && contains(levels_slots, "Wooing_Gloria"){
			i := 0
			for levels_slots[i] != "VolcanoRave" {
				i++
			}
			if levelsToShuffle[i] == "Wooing_Gloria" {
				s := "this is a rare seed where wooing gloria manages to get sorted to volcano rave, please send me the seed and any issues you have since i can't easily test this edge case without it"
				SpoilerLogger.Println(s)
				fmt.Println(s)
				return shuffleLevels(levels_slots, levelsToShuffle)
			}
		}
	
		if contains(levels_slots, "Watercaves") && contains(levels_slots, "penguins"){
			i := 0
			for levels_slots[i] != "penguins" {
				i++
			}
			if levelsToShuffle[i] == "Watercaves" {
				s := "watercaves in penguins, reshuffling"
				SpoilerLogger.Println(s)
				fmt.Println(s)
				return shuffleLevels(levels_slots, levelsToShuffle)
			}
		}
	}

	return levels_slots
}

func randomizeLevels(){
	initializeLevels()
	rand.Seed(int64(seed))

	for group := range level_groups {
		if group < len(level_groups)-1 {
			levels_slots := make([]string, len(level_groups[group]))
			levels_slots = shuffleLevels(levels_slots, level_groups[group])
			
			for level := range level_groups[group] {
				copyFile("ContentOG/Streams/win/"+level_groups[group][level]+".arc", "Content/Streams/win/"+levels_slots[level]+".arc")
				copyFile("ContentOG/Streams/win/"+level_groups[group][level]+".bld", "Content/Streams/win/"+levels_slots[level]+".bld")
				SpoilerLogger.Println(level_groups[group][level]+" to "+levels_slots[level])
			}
		} else {
			for level := range level_groups[group] {
				copyFile("ContentOG/Streams/win/"+level_groups[group][level]+".arc", "Content/Streams/win/"+level_groups[group][level]+".arc")
				copyFile("ContentOG/Streams/win/"+level_groups[group][level]+".bld", "Content/Streams/win/"+level_groups[group][level]+".bld")
			}
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
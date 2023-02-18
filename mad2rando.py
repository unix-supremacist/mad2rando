import shutil
import os
import random
import logging

print("Hey the randomizer is randomizing *cough* xeelium *cough*")

seed = 0
deadends = False
#Any minigames including golf can trigger minigame mode
#this makes certain levels impossible to beat
#if you cant get to them in game mode use with caution
randomizeGolf = False
randomizeDiving = False
randomizeMinigames = False
content = "Content"
contentog = "ContentOG"
streams = "/Streams/win/"
mptourney = "MultiplayerTourneyFinish"
bld = ".bld"
arc = ".arc"
logger = logging.getLogger()
logger.setLevel(logging.INFO)
file_handler = logging.FileHandler('spoiler.log', mode='w')
file_handler.setLevel(logging.DEBUG)
logger.addHandler(file_handler)

deadends_golf_levels = ["golf_cake", "golf_lovelylumps", "golf_targetTree"]
golf_levels = ["golf_3holes", "golf_baggagecheck", \
"golf_cratercrossing", "golf_crossedpaths", "golf_foosaBall",\
"golf_junkyard",  "golf_maze",\
"Golf_minigame", "golf_ravenousRhinos", ]

diving_levels = ["DivingLocation_IslandFever",\
"DivingLocation_PrepareToLaunch", "DivingLocation_RitesOfPassage",\
"DivingLocation_Waterhole", "Minigame_Diving_Location_Menu"]

minigame_levels = ["animal_chess", "Card_Match_Game", "DrMelman",\
"HungryHippo", "Minigame_HotDurian", "RoP_MusicalChairs",\
"Soccer"]

deadends_levels = ["Credits", "Dam_Busters", "FixThePlane", "Waterhole"]

levels = ["BraveNewWild", "ConvoyChase", "DutyFree", "IslandFever",\
"MartyRace", "Morts_Adventure", "penguins", "penguins2",\
"Prepare2Launch_Plane", "Prepare2Launch", "RitesOfPassage",\
"VolcanoRave", "Watercaves", "Wooing_Gloria"]

minilevels = []

manualoverrides = ["map", "title", "global"]

if deadends == True:
	levels.extend(deadends_levels)
	golf_levels.extend(deadends_golf_levels)
else:
	manualoverrides.extend(deadends_levels)
	manualoverrides.extend(deadends_golf_levels)

if randomizeGolf == True:
	levels.extend(golf_levels)
else:
	minilevels.extend(golf_levels)

if randomizeDiving == True:
	levels.extend(diving_levels)
else:
	minilevels.extend(diving_levels)

if randomizeMinigames == True:
	levels.extend(minigame_levels)
else:
	minilevels.extend(minigame_levels)

levels.sort()
levelscopied = []
levelscopied.extend(levels)
minilevels.sort()
minilevelscopied = []
minilevelscopied.extend(minilevels)

if os.path.isdir(content) == False:
	quit()

if os.path.isdir(contentog) == False:
	print("copy Content to ContentOG")
	quit()

random.seed(seed)

shutil.rmtree(content)
os.makedirs(content+streams)
logger.info('seed:'+str(seed))
for level in levels:
	x = random.randrange(0, len(levelscopied))
	shutil.copy(contentog+streams+level+arc, content+streams+levelscopied[x]+arc)
	shutil.copy(contentog+streams+level+bld, content+streams+levelscopied[x]+bld)
	logger.info('Main Pool Level:'+level+':'+levelscopied[x])
	levelscopied.pop(x)

for level in minilevels:
	x = random.randrange(0, len(minilevelscopied))
	shutil.copy(contentog+streams+level+arc, content+streams+minilevelscopied[x]+arc)
	shutil.copy(contentog+streams+level+bld, content+streams+minilevelscopied[x]+bld)
	logger.info('Minigame Pool Level:'+level+':'+minilevelscopied[x])
	minilevelscopied.pop(x)

for level in manualoverrides:
	shutil.copy(contentog+streams+level+arc, content+streams+level+arc)
	shutil.copy(contentog+streams+level+bld, content+streams+level+bld)

shutil.copy(contentog+streams+mptourney+bld, content+streams+mptourney+bld)

print("Game randomized")
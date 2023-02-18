import shutil
import os
import random
import logging

seed = 0
randomizeGolf = True
randomizeDiving = True
randomizeMinigames = True
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

golf_levels = ["golf_3holes", "golf_baggagecheck", "golf_cake",\
"golf_cratercrossing", "golf_crossedpaths", "golf_foosaBall",\
"golf_junkyard", "golf_lovelylumps", "golf_maze",\
"Golf_minigame", "golf_ravenousRhinos", "golf_targetTree"]

diving_levels = ["DivingLocation_IslandFever",\
"DivingLocation_PrepareToLaunch", "DivingLocation_RitesOfPassage",\
"DivingLocation_Waterhole", "Minigame_Diving_Location_Menu"]

minigame_levels = ["animal_chess", "Card_Match_Game", "DrMelman",\
"DutyFree", "HungryHippo", "Minigame_HotDurian", "RoP_MusicalChairs",\
"Soccer"]

levels = ["BraveNewWild", "ConvoyChase", "Credits", "Dam_Busters", \
"FixThePlane", "IslandFever", "MartyRace", "Morts_Adventure",\
"penguins", "penguins2", "Prepare2Launch_Plane", "Prepare2Launch",\
"RitesOfPassage", "VolcanoRave", "Watercaves", "Waterhole",\
"Wooing_Gloria"]

manualoverrides = ["map", "title", "global"]

if randomizeGolf == True:
	levels.extend(golf_levels)
else:
	manualoverrides.extend(golf_levels)

if randomizeDiving == True:
	levels.extend(diving_levels)
else:
	manualoverrides.extend(diving_levels)

if randomizeMinigames == True:
	levels.extend(minigame_levels)
else:
	manualoverrides.extend(minigame_levels)

levels.sort()
levelscopied = []
levelscopied.extend(levels)

if os.path.isdir(content) == False:
	quit()

if os.path.isdir(contentog) == False:
	shutil.copytree(r"Content", r"ContentOG")

random.seed(seed)

shutil.rmtree(content)
os.makedirs(content+streams)
logger.info('seed:'+str(seed))
for level in levels:
	x = random.randrange(0, len(levelscopied))
	shutil.copy(contentog+streams+level+arc, content+streams+levelscopied[x]+arc)
	shutil.copy(contentog+streams+level+bld, content+streams+levelscopied[x]+bld)
	logger.info(level+':'+levelscopied[x])
	levelscopied.pop(x)

for level in manualoverrides:
	shutil.copy(contentog+streams+level+arc, content+streams+level+arc)
	shutil.copy(contentog+streams+level+bld, content+streams+level+bld)

shutil.copy(contentog+streams+mptourney+bld, content+streams+mptourney+bld)
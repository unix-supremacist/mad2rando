import shutil
import os
import random
import logging
import configparser
import json

#print that the rando is working to a certain someone
print("Hey the randomizer is randomizing *cough* xeelium *cough*")

#Load/Create Config
configname = "rando.ini"
config = configparser.ConfigParser()
config['DEFAULT'] = {	'seed': '0',
						'sanity': 'True',
						'randomizeMinigames': 'False',
						'monkeylessLevels': 'True',
						'hasDeadends': 'False'}

if os.path.isfile(configname) == False:
	with open(configname, 'w') as configfile:
		config.write(configfile)
else:
	config.read(configname)

seed = int(config['DEFAULT']['seed'])
sanity = json.loads(config['DEFAULT']['sanity'].lower())
randomizeMinigames = json.loads(config['DEFAULT']['randomizeMinigames'].lower())
monkeylesslevels = json.loads(config['DEFAULT']['monkeylesslevels'].lower())
hasDeadends = json.loads(config['DEFAULT']['hasDeadends'].lower())

#Set basic vars
content = "Content"
contentog = "ContentOG"
streams = "/Streams/win/"
mptourney = "MultiplayerTourneyFinish"
bld = ".bld"
arc = ".arc"

#Setup the spoiler log
logger = logging.getLogger()
logger.setLevel(logging.INFO)
file_handler = logging.FileHandler('spoiler.log', mode='w')
file_handler.setLevel(logging.DEBUG)
logger.addHandler(file_handler)

#Level pools
deadend_pool = ["animal_chess", "Card_Match_Game", "Credits",\
"DivingLocation_IslandFever", "DivingLocation_PrepareToLaunch",\
"DivingLocation_RitesOfPassage", "DivingLocation_Waterhole",\
"DutyFree", "golf_cake", "golf_lovelylumps", "golf_targetTree",\
"HungryHippo", "Minigame_HotDurian", "RoP_MusicalChairs", "Soccer"]

minigame_levels = ["animal_chess", "Card_Match_Game",\
"DivingLocation_IslandFever", "DivingLocation_PrepareToLaunch",\
"DivingLocation_RitesOfPassage", "DivingLocation_Waterhole",\
"Minigame_Diving_Location_Menu", "HungryHippo", "Minigame_HotDurian",\
"RoP_MusicalChairs", "Soccer",\
"golf_3holes", "golf_baggagecheck",\
"golf_cratercrossing", "golf_crossedpaths", "golf_foosaBall",\
"golf_junkyard",  "golf_maze",\
"Golf_minigame", "golf_ravenousRhinos", "golf_cake",\
"golf_lovelylumps", "golf_targetTree"]

system_pool = ["map", "title", "global"]
required_pool = ["Dam_Busters", "FixThePlane", "Waterhole"]

#do nothing with these for now they are impossible to reach
impossible_boring_pool = ["penguins", "penguins2", "Prepare2Launch_Plane", "Prepare2Launch"]

campaign_pool = ["BraveNewWild", "ConvoyChase", "Credits",\
"DrMelman", "MartyRace", "RitesOfPassage",\
"VolcanoRave", "Watercaves", "Wooing_Gloria"]

boring_pool = ["DutyFree", "IslandFever", "Morts_Adventure"]
campaign_pool.extend(boring_pool)

minigame_pool = []

#Level blacklist doesn't allow these levels to be put into the final rando
blacklist = []

#Config handling
if hasDeadends == False:
	blacklist.extend(deadend_pool)

random.seed(seed)

if sanity == False:
	campaign_pool.extend(required_pool)	
else:
	system_pool.extend(required_pool)

if randomizeMinigames == True:
	campaign_pool.extend(minigame_levels)
else:
	minigame_pool.extend(minigame_levels)

#Lists of how many files we've yet to replace
levelscopied = []
levelscopied.extend(campaign_pool)
minigame_poolcopied = []
minigame_poolcopied.extend(minigame_pool)

#Finsh Config handling
if monkeylesslevels == True:
	campaign_pool.extend(impossible_boring_pool)
	levelscopied.extend(impossible_boring_pool)
else:
	blacklist.extend(boring_pool)

levelscopied.extend(boring_pool)

#Sort the pools so seeds are more stable
campaign_pool.sort()
levelscopied.sort()
minigame_pool.sort()
minigame_poolcopied.sort()

#Check for ContentOG
if os.path.isdir(contentog) == False:
	print("copy Content to ContentOG")
	quit()

#Replace content folder
shutil.rmtree(content)
os.makedirs(content+streams)

#Log the seed and the levels rando'd
logger.info('seed:'+str(seed))
logger.info('lvls:'+str(len(campaign_pool)))

#Blacklist lists setup
to_remove = []
to_add = []

#campaign pool blacklist filtering
for level in campaign_pool:
	if level in blacklist:
		to_remove.append(level)
		y = [item for item in campaign_pool if item not in blacklist]
		z = random.randrange(0, len(y))
		to_add.append(y[z])

campaign_pool = [item for item in campaign_pool if item not in to_remove]
campaign_pool.extend(to_add)
campaign_pool.sort()

#Blacklist lists cleanup
to_remove.clear()
to_add.clear()

#minigame pool blacklist filtering
if len(minigame_pool) != 0:
	for level in minigame_pool:
		if level in blacklist:
			to_remove.append(level)
			y = [item for item in minigame_pool if item not in blacklist]
			z = random.randrange(0, len(y))
			to_add.append(y[z])

minigame_pool = [item for item in minigame_pool if item not in to_remove]
minigame_pool.extend(to_add)
minigame_pool.sort()

#Copy campaign levels
for level in campaign_pool:
		x = random.randrange(0, len(levelscopied))
		shutil.copy(contentog+streams+level+arc, content+streams+levelscopied[x]+arc)
		shutil.copy(contentog+streams+level+bld, content+streams+levelscopied[x]+bld)
		logger.info('Main Pool Level:'+level+':'+levelscopied[x])
		levelscopied.pop(x)

#Copy minigame levels
if len(minigame_pool) != 0:
	for level in minigame_pool:
		x = random.randrange(0, len(minigame_poolcopied))
		shutil.copy(contentog+streams+level+arc, content+streams+minigame_poolcopied[x]+arc)
		shutil.copy(contentog+streams+level+bld, content+streams+minigame_poolcopied[x]+bld)
		logger.info('Minigame Pool Level:'+level+':'+minigame_poolcopied[x])
		minigame_poolcopied.pop(x)

#Copy system(unrandomized) levels
for level in system_pool:
	shutil.copy(contentog+streams+level+arc, content+streams+level+arc)
	shutil.copy(contentog+streams+level+bld, content+streams+level+bld)

shutil.copy(contentog+streams+mptourney+bld, content+streams+mptourney+bld)

#Say its done
print("Game randomized")
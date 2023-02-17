import shutil
import os
import random
import logging

seed = 0
logger = logging.getLogger()
logger.setLevel(logging.INFO)
file_handler = logging.FileHandler('spoiler.log')
file_handler.setLevel(logging.DEBUG)
logger.addHandler(file_handler)

levels = ["animal_chess", "BraveNewWild", "Card_Match_Game", "ConvoyChase", "Credits", "Dam_Busters", "DivingLocation_IslandFever", \
"DivingLocation_PrepareToLaunch", "DivingLocation_RitesOfPassage", "DivingLocation_Waterhole", "DrMelman", "DutyFree", \
"FixThePlane", "golf_3holes", "golf_baggagecheck", "golf_cake", "golf_cratercrossing", "golf_crossedpaths", \
"golf_foosaBall", "golf_junkyard", "golf_lovelylumps", "golf_maze", "Golf_minigame", "golf_ravenousRhinos", \
"golf_targetTree", "HungryHippo", "IslandFever", "MartyRace", "Minigame_Diving_Location_Menu", "Minigame_HotDurian", \
"Morts_Adventure", "penguins", "penguins2", "Prepare2Launch_Plane", "Prepare2Launch", "RitesOfPassage", "RoP_MusicalChairs", \
"Soccer", "VolcanoRave", "Watercaves", "Waterhole", "Wooing_Gloria"]

levelscopied = ["animal_chess", "BraveNewWild", "Card_Match_Game", "ConvoyChase", "Credits", "Dam_Busters", "DivingLocation_IslandFever", \
"DivingLocation_PrepareToLaunch", "DivingLocation_RitesOfPassage", "DivingLocation_Waterhole", "DrMelman", "DutyFree", \
"FixThePlane", "golf_3holes", "golf_baggagecheck", "golf_cake", "golf_cratercrossing", "golf_crossedpaths", \
"golf_foosaBall", "golf_junkyard", "golf_lovelylumps", "golf_maze", "Golf_minigame", "golf_ravenousRhinos", \
"golf_targetTree", "HungryHippo", "IslandFever", "MartyRace", "Minigame_Diving_Location_Menu", "Minigame_HotDurian", \
"Morts_Adventure", "penguins", "penguins2", "Prepare2Launch_Plane", "Prepare2Launch", "RitesOfPassage", "RoP_MusicalChairs", \
"Soccer", "VolcanoRave", "Watercaves", "Waterhole", "Wooing_Gloria"]

if os.path.isdir('Content') == False:
	quit()

if os.path.isdir('ContentOG') == False:
	shutil.copytree(r"Content", r"ContentOG")

random.seed(seed)

shutil.rmtree("Content")
os.makedirs("Content/Streams/win/")
for level in levels:
	x = random.randrange(0, len(levelscopied))
	shutil.copy('ContentOG/Streams/win/'+level+'.arc', 'Content/Streams/win/'+levelscopied[x]+'.arc')
	shutil.copy('ContentOG/Streams/win/'+level+'.bld', 'Content/Streams/win/'+levelscopied[x]+'.bld')
	logger.info('ContentOG/Streams/win/'+level+': Moved to: Content/Streams/win/'+levelscopied[x])
	levelscopied.pop(x)

shutil.copy('ContentOG/Streams/win/map.arc', 'Content/Streams/win/map.arc')
shutil.copy('ContentOG/Streams/win/map.bld', 'Content/Streams/win/map.bld')
shutil.copy('ContentOG/Streams/win/title.arc', 'Content/Streams/win/title.arc')
shutil.copy('ContentOG/Streams/win/title.bld', 'Content/Streams/win/title.bld')
shutil.copy('ContentOG/Streams/win/MultiplayerTourneyFinish.bld', 'Content/Streams/win/MultiplayerTourneyFinish.bld')
shutil.copy('ContentOG/Streams/win/global.arc', 'Content/Streams/win/global.arc')
shutil.copy('ContentOG/Streams/win/global.bld', 'Content/Streams/win/global.bld')
package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

/*
	NPC Name: 		Wisp
	Map(s): 		Ludibrium: Eos Tower Entrance
	Description: 	
*/
class NPC2040030 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (status >= 0 && mode == 0) {
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }

         if (status == 0) {
            cm.sendSimple("2040030_HELLO")
         } else if (status == 1) {
            if (selection == 0) {
               cm.sendNext("2040030_I_AM_WISP")
               cm.dispose()
            } else if (selection == 1) {
               cm.sendNext("2040030_LOT_OF_QUESTIONS")
            } else if (selection == 2) {
               cm.sendNext("2040030_COMMAND")
            } else if (selection == 3) {
               cm.sendNext("2040030_DYING")
            } else if (selection == 4) {
               cm.sendNext("2040030_KITTY_COMMAND")
               cm.dispose()
            } else if (selection == 5) {
               cm.sendNext("2040030_BROWN_PUPPY_COMMAND")
               cm.dispose()
            } else if (selection == 6) {
               cm.sendNext("2040030_BUNNY_COMMAND")
               cm.dispose()
            } else if (selection == 7) {
               cm.sendNext("2040030_CARGO_COMMAND")
               cm.dispose()
            } else if (selection == 8) {
               cm.sendNext("2040030_HUSKY_COMMAND")
               cm.dispose()
            } else if (selection == 9) {
               cm.sendNext("2040030_BLACK_PIG_COMMAND")
               cm.dispose()
            } else if (selection == 10) {
               cm.sendNext("2040030_PANDA_COMMAND")
               cm.dispose()
            } else if (selection == 11) {
               cm.sendNext("2040030_DINO_COMMAND")
               cm.dispose()
            } else if (selection == 12) {
               cm.sendNext("2040030_RUDOLPH")
               cm.dispose()
            } else if (selection == 13) {
               cm.sendNext("2040030_MONKEY")
               cm.dispose()
            } else if (selection == 14) {
               cm.sendNext("2040030_ROBOT_COMMAND")
               cm.dispose()
            } else if (selection == 15) {
               cm.sendNext("2040030_ELEPHANT_COMMAND")
               cm.dispose()
            } else if (selection == 16) {
               cm.sendNext("2040030_GOLDEN_PING")
               cm.dispose()
            } else if (selection == 17) {
               cm.sendNext("2040030_PENGUIN_COMMAND")
               cm.dispose()
            } else if (selection == 18) {
               cm.sendNext("2040030_MINI_YETI_COMMAND")
               cm.dispose()
            } else if (selection == 19) {
               cm.sendNext("2040030_JR_BALROG_COMMAND")
               cm.dispose()
            } else if (selection == 20) {
               cm.sendNext("2040030_BABY_DRAGON_COMMAND")
               cm.dispose()
            } else if (selection == 21) {
               cm.sendNext("2040030_DRAGON_COMMAND")
               cm.dispose()
            } else if (selection == 22) {
               cm.sendNext("2040030_BLACK_DRAGON_COMMAND")
               cm.dispose()
            } else if (selection == 23) {
               cm.sendNext("2040030_SNOWMAN_COMMAND")
               cm.dispose()
            } else if (selection == 24) {
               cm.sendNext("2040030_SUN_WU_KONG_COMMAND")
               cm.dispose()
            } else if (selection == 25) {
               cm.sendNext("2040030_JR_REAPER_COMMAND")
               cm.dispose()
            } else if (selection == 26) {
               cm.sendNext("2040030_CRYSTAL_RUDOLPH_COMMAND")
               cm.dispose()
            } else if (selection == 27) {
               cm.sendNext("2040030_KINO_COMMAND")
               cm.dispose()
            } else if (selection == 28) {
               cm.sendNext("2040030_WHITE_DUCK_COMMAND")
               cm.dispose()
            } else if (selection == 29) {
               cm.sendNext("2040030_PINK_BEAN_COMMAND")
               cm.dispose()
            } else if (selection == 30) {
               cm.sendNext("2040030_PORCUPINE_COMMAND")
               cm.dispose()
            }
         } else if (status == 2) {
            cm.sendNextPrev("2040030_WATER_OF_LIFE")
         } else if (status == 3) {
            cm.sendNextPrev("2040030_SPECIAL_COMMANDS")
            cm.dispose()
         } else if (status == 4) {
            cm.sendNextPrev("2040030_TALK_TO_THE_PET")
         } else if (status == 5) {
            cm.sendNextPrev("2040030_GETS_HUNGRY")
         } else if (status == 6) {
            cm.sendNextPrev("2040030_CANNOT_EAT_NORMAL_FOOD")
         } else if (status == 7) {
            cm.sendNextPrev("2040030_GOES_BACK_HOME")
            cm.dispose()
         } else if (status == 8) {
            cm.sendNextPrev("2040030_AFTER_SOME_TIME")
         } else if (status == 9) {
            cm.sendNextPrev("2040030_SAD_TO_SEE_THEM_STOP")
            cm.dispose()
         }
      }
   }
}

NPC2040030 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2040030(cm: cm))
   }
   return (NPC2040030) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
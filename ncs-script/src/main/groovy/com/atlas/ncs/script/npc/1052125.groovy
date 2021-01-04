package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

/*
	NPC Name: 		June
	Map(s): 		Kerning Square : 7th Floor
	Description: 	Entrance to Spirit of Rock
*/
class NPC1052125 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      cm.sendSimple("1052125_ACCESS_LIMITED")
   }

   def action(Byte mode, Byte type, Integer selection) {
      status++
      if (mode != 1) {
         if (mode == 0 && type != 4) {
            status -= 2
         } else {
            cm.dispose()
            return
         }
      }
      if (status == 0) {
         if (selection == 0) {
            if (cm.isQuestStarted(2286) || cm.isQuestStarted(2287) || cm.isQuestStarted(2288)) {
               EventManager em = cm.getEventManager("RockSpirit")
               if (!em.startInstance(cm.getPlayer())) {
                  cm.sendOk("1052125_ROOMS_CROWDED")
               }
               cm.dispose()
               return
            } else {
               cm.sendOk("1052125_DID_YOU_HEAR_ANYTHING")
            }
         } else {
            if (cm.isQuestCompleted(2290)) {
               if (cm.getLevel() > 50) {
                  cm.sendOk("1052125_VIP_LEVEL_REQUIREMENT")
               } else {
                  cm.sendOk("1052125_VIP_NEEDS_TICKET")
               }
            } else {
               cm.sendOk("1052125_GET_LOST")
            }
         }
         cm.dispose()
      }
   }
}

NPC1052125 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1052125(cm: cm))
   }
   return (NPC1052125) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
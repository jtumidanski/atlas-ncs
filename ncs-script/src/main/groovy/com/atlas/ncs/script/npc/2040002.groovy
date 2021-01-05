package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

/*
	NPC Name: 		Olson the Toy Soldier
	Map(s): 		
	Description: 	
*/
class NPC2040002 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   EventManager em

   def start() {
      if (cm.isQuestStarted(3230)) {
         em = cm.getEventManager("DollHouse").orElseThrow()

         if (em.getProperty("noEntry") == "false") {
            cm.sendNext("2040002_PENDULUM_INSIDE")
         } else {
            cm.sendOk("2040002_SOMEONE_ALREADY")
            cm.dispose()
         }
      } else {
         cm.sendOk("2040002_NOT_ALLOWED")
         cm.dispose()
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode < 1) {
         cm.dispose()
      } else {
         status++
         if (status == 1) {
            cm.sendYesNo("2040002_ARE_YOU_READY")
         } else if (status == 2) {
            em = cm.getEventManager("DollHouse").orElseThrow()
            if (!em.startInstance(cm.getCharacterId())) {
               cm.sendOk("2040002_ALREADY_BEING_CHALLENGED")
            }

            cm.dispose()
         }
      }
   }
}

NPC2040002 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2040002(cm: cm))
   }
   return (NPC2040002) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
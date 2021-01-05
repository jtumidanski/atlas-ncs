package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC1032008 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.haveItem(4031045)) {
         Optional<EventManager> eventManager = cm.getEventManager("Boats")
         if (eventManager.isPresent() && eventManager.get().getProperty("entry") == "true") {
            cm.sendYesNo("1032008_GO_TO_ORBIS")
         } else {
            cm.sendOk("1032008_ALREADY_TRAVELLING")
            cm.dispose()
         }
      } else {
         cm.sendOk("1032008_CHECK_YOUR_INVENTORY")
         cm.dispose()
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode <= 0) {
         cm.sendOk("1032008_TALK_TO_ME_AGAIN")
         cm.dispose()
         return
      }
      Optional<EventManager> eventManager = cm.getEventManager("Boats")
      if (eventManager.isPresent() && eventManager.get().getProperty("entry") == "true") {
         cm.warp(101000301)
         cm.gainItem(4031045, (short) -1)
         cm.dispose()
      } else {
         cm.sendOk("1032008_BE_PATIENT_FOR_NEXT_ONE")
         cm.dispose()
      }
   }
}

NPC1032008 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1032008(cm: cm))
   }
   return (NPC1032008) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
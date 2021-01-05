package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC2060005 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.isQuestCompleted(6002)) {
         cm.sendOk("2060005_THANKS_FOR_SAVING_THE_PORK")

      } else if (cm.isQuestStarted(6002)) {
         if (cm.haveItem(4031507, 5) && cm.haveItem(4031508, 5)) {
            cm.sendOk("2060005_THANKS_FOR_SAVING_THE_PORK")

         } else {
            EventManager em = cm.getEventManager("3rdJob_mount").orElseThrow()
            if (em == null) {
               cm.sendOk("2060005_IS_CLOSED")

            } else {
               if (em.startInstance(cm.getCharacterId())) {
                  cm.removeAll(4031507)
                  cm.removeAll(4031508)
               } else {
                  cm.sendOk("2060005_SOMEONE_IN_THE_MAP")

               }
            }
         }
      } else {
         cm.sendOk("2060005_NOT_ELIGIBLE")

      }

      cm.dispose()
   }

   def action(Byte mode, Byte type, Integer selection) {

   }
}

NPC2060005 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2060005(cm: cm))
   }
   return (NPC2060005) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
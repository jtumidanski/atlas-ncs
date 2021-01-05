package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9010021 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      status = -1
      if (!cm.getConfiguration().useRebirthSystem()) {
         cm.sendOk("9010021_ASSIST_THE_FIGHT")
         cm.dispose()
         return
      }
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == 1) {
         status++
      } else {
         cm.dispose()
         return
      }
      if (status == 0) {
         cm.sendNext("9010021_REBORN", cm.countRebirths())
      } else if (status == 1) {
         cm.sendSimple("9010021_WHAT_DO_YOU_WANT_ME_TO_DO")
      } else if (status == 2) {
         if (selection == 0) {
            if (cm.getLevel() == 200) {
               cm.sendYesNo("9010021_ARE_YOU_SURE")
            } else {
               cm.sendOk("9010021_NOT_LEVEL_200")
               cm.dispose()
            }
         } else if (selection == 1) {
            cm.sendOk("9010021_BYE")
            cm.dispose()
         }
      } else if (status == 3 && type == 1) {
         cm.executeRebirth()
         cm.sendOk("9010021_REBORN_SUCCESS", cm.countRebirths())
         cm.dispose()
      }
   }
}

NPC9010021 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9010021(cm: cm))
   }
   return (NPC9010021) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
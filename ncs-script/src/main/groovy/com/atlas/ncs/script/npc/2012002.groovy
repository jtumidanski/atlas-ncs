package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2012002 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   def start() {
      cm.sendYesNo("2012002_WISH_TO_LEAVE")
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == 0 && status == 1) {
         cm.sendOk("2012002_GOOD_CHOICE")
         cm.dispose()
      }
      if (mode > 0) {
         status++
      } else {
         cm.dispose()
      }

      if (status == 1) {
         cm.sendNext("2012002_ALRIGHT")
      } else if (status == 2) {
         cm.warp(200000111, 0)// back to Orbis jetty
         cm.dispose()
      }
   }
}

NPC2012002 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2012002(cm: cm))
   }
   return (NPC2012002) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
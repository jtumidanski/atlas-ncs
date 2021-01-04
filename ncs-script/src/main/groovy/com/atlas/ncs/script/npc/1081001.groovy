package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1081001 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int returnMap

   def start() {
      returnMap = cm.peekSavedLocation("FLORINA")
      if (returnMap == -1) {
         returnMap = 104000000
      }
      cm.sendNext("1081001_SO_YOU_WANT_TO_LEAVE", returnMap)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else if (mode == 0) {
         cm.sendNext("1081001_MUST_HAVE_SOME_BUSINESS", returnMap)
         cm.dispose()
      } else if (mode == 1) {
         status++
         if (status == 1) {
            cm.sendYesNo("1081001_ARE_YOU_SURE", returnMap, returnMap)
         } else {
            cm.getSavedLocation("FLORINA")
            cm.warp(returnMap)
            cm.dispose()
         }
      }
   }
}

NPC1081001 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1081001(cm: cm))
   }
   return (NPC1081001) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
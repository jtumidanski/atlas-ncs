package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9120015 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      cm.sendSimple("9120015_WHAT_DO_YOU_WANT")
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode < 1) {
         cm.dispose()
      } else {
         status++
         if (status == 1) {
            if (selection == 0) {
               cm.sendNext("9120015_TO_THE_HIDEOUT")
               cm.dispose()
            } else if (selection == 1) {
               cm.sendNext("9120015_THE_BRAVE_ONE")
            } else {
               cm.sendOk("9120015_I_AM_BUSY")
               cm.dispose()
            }
         } else {
            cm.warp(801040000, "in00")
            cm.dispose()
         }
      }
   }
}

NPC9120015 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9120015(cm: cm))
   }
   return (NPC9120015) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
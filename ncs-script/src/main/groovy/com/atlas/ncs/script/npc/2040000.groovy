package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2040000 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int cost = 6000

   def start() {
      cm.sendYesNo("2040000_HELLO")
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 1) {
            status++
         }
         if (mode == 0) {
            cm.sendNext("2040000_MUST_HAVE_SOME_BUSINESS")
            cm.dispose()
            return
         }
         if (status == 1) {
            if (cm.getMeso() >= cost && cm.canHold(4031045)) {
               cm.gainItem(4031045, (short) 1)
               cm.gainMeso(-cost)
            } else {
               cm.sendOk("2040000_ARE_YOU_SURE", cost)
            }
            cm.dispose()
         }
      }
   }
}

NPC2040000 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2040000(cm: cm))
   }
   return (NPC2040000) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
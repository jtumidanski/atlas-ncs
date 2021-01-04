package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2102002 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int cost = 6000

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 1) {
            status++
         }
         if (mode == 0) {
            cm.sendNext("2102002_MUST_HAVE_SOME_OTHER_BUSINESS")

            cm.dispose()
            return
         }
         if (status == 0) {
            cm.sendYesNo("2102002_HELLO", cost)

         } else if (status == 1) {
            if (cm.getMeso() >= cost && cm.canHold(4031045)) {
               cm.gainItem(4031045, (short) 1)
               cm.gainMeso(-cost)
            } else {
               cm.sendOk("2102002_NOT_ENOUGH_MESOS", cost)

            }
            cm.dispose()
         }
      }
   }
}

NPC2102002 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2102002(cm: cm))
   }
   return (NPC2102002) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
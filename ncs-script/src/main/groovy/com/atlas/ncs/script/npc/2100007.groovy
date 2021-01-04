package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2100007 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int[] skin = [0, 1, 2, 3, 4]

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode < 1) {
         cm.dispose()
      } else {
         if (mode == 1) {
            status++
         } else {
            status--
         }
         if (status == 0) {
            cm.sendNext("2100007_HELLO")

         } else if (status == 1) {
            cm.sendStyle("With our specialized machine, you can see yourself after the treatment in advance. What kind of skin-treatment would you like to do? Choose the style of your liking...", skin)
         } else if (status == 2) {
            cm.dispose()
            if (cm.haveItem(5153007)) {
               cm.gainItem(5153007, (short) -1)
               cm.setSkin(skin[selection])
               cm.sendOk("2100007_ENJOY_NEW_SKIN")

            } else {
               cm.sendNext("2100007_MISSING_SKIN_COUPON")

            }
         }
      }
   }
}

NPC2100007 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2100007(cm: cm))
   }
   return (NPC2100007) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
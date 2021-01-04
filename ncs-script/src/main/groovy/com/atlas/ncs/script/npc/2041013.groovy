package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2041013 {
   NPCConversationManager cm
   int status = 0
   int sel = -1
   int price = 1000000
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
            cm.sendSimple("2041013_HELLO")

         } else if (status == 1) {
            if (selection == 2) {
               cm.sendStyle("With our specialized machine, you can see the way you'll look after the treatment PRIOR to the procedure. What kind of a look are you looking for? Go ahead and choose the style of your liking~!", skin)
            }
         } else if (status == 2) {
            cm.dispose()
            if (cm.haveItem(5153002)) {
               cm.gainItem(5153002, (short) -1)
               cm.setSkin(skin[selection])
               cm.sendOk("2041013_ENJOY_NEW_SKIN")

            } else {
               cm.sendOk("2041013_MISSING_SKIN_COUPON")

            }
         }
      }
   }
}

NPC2041013 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2041013(cm: cm))
   }
   return (NPC2041013) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager
import com.atlas.ncs.util.ScriptUtils

class NPC9201061 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int price = 1000000
   int[] colors = []

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
            cm.sendSimple("9201061_HI_THERE")

         } else if (status == 1) {
            if (selection == 2) {
               int current = 0
               if (cm.getGender() == 0) {
                  current = cm.getFace() % 100 + 20000
               } else if (cm.getGender() == 1) {
                  current = cm.getFace() % 100 + 21000
               }
               int[] temp = [current + 100, current + 200, current + 300, current + 400, current + 500, current + 600, current + 700]
               colors = ScriptUtils.pushItemsIfTrue(colors, temp, { itemId -> cm.cosmeticExistsAndIsNotEquipped(itemId) })
               cm.sendYesNo("9201061_REG_CONFIRM")

            }
         } else if (status == 2) {
            cm.dispose()
            if (cm.haveItem(5152035)) {
               cm.gainItem(5152035, (short) -1)
               cm.setFace(colors[Math.floor(Math.random() * colors.length).intValue()])
               cm.sendOk("9201061_ENJOY_NEW_LENS")

            } else {
               cm.sendOk("9201061_MISSING_LENS_COUPON")

            }
         }
      }
   }
}

NPC9201061 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201061(cm: cm))
   }
   return (NPC9201061) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
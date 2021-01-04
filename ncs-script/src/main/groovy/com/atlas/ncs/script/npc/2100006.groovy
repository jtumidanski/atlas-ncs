package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2100006 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int beauty = 0
   int[] maleHair = [30150, 30170, 30180, 30320, 30330, 30410, 30460, 30820, 30900]
   int[] femaleHair = [31040, 31090, 31190, 31330, 31340, 31400, 31420, 31620, 31660]
   int[] hairNew = []
   int[] hairColor = []

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
            cm.sendSimple("2100006_HELLO")

         } else if (status == 1) {
            if (selection == 0) {
               beauty = 1
               hairNew = []
               if (cm.getGender() == 0) {
                  for (int i = 0; i < maleHair.length; i++) {
                     hairNew = ScriptUtils.pushItemIfTrue(hairNew, maleHair[i] + (cm.getHair() % 10).intValue(), { itemId -> cm.cosmeticExistsAndIsntEquipped(itemId) })
                  }
               }
               if (cm.getGender() == 1) {
                  for (int i = 0; i < femaleHair.length; i++) {
                     hairNew = ScriptUtils.pushItemIfTrue(hairNew, femaleHair[i] + (cm.getHair() % 10).intValue(), { itemId -> cm.cosmeticExistsAndIsntEquipped(itemId) })
                  }
               }
               cm.sendStyle("Hahaha~all you need is #bAriant hair style coupon(VIP)#k to change up your hairstyle. Choose the new style, and let me do the rest.", hairNew)
            } else if (selection == 1) {
               beauty = 2
               hairColor = []
               int current = (cm.getHair() / 10).intValue() * 10
               for (int i = 0; i < 8; i++) {
                  hairColor = ScriptUtils.pushItemIfTrue(hairColor, current + i, { itemId -> cm.cosmeticExistsAndIsntEquipped(itemId) })
               }
               cm.sendStyle("Every once in a while, it doesn't hurt to change up your hair color... it's fun. Allow me, the great Mazra, to dye your hair, so you just bring me #bAriant hair color coupon(VIP)#k, and choose your new hair color.", hairColor)
            }
         } else if (status == 2) {
            cm.dispose()
            if (beauty == 1) {
               if (cm.haveItem(5150027)) {
                  cm.gainItem(5150027, (short) -1)
                  cm.setHair(hairNew[selection])
                  cm.sendOk("2100006_ENJOY_NEW_STYLE")

               } else {
                  cm.sendNext("2100006_MISSING_STYLE_COUPON")

               }
            }
            if (beauty == 2) {
               if (cm.haveItem(5151022)) {
                  cm.gainItem(5151022, (short) -1)
                  cm.setHair(hairColor[selection])
                  cm.sendOk("2100006_ENJOY_NEW_COLOR")

               } else {
                  cm.sendNext("2100006_MISSING_COLOR_COUPON")

               }
            }
         }
      }
   }
}

NPC2100006 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2100006(cm: cm))
   }
   return (NPC2100006) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager
import com.atlas.ncs.util.ScriptUtils

class NPC9120100 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int beauty = 0
   int hairPrice = 1000000
   int hairColorPrice = 1000000
   int[] maleHair = [30260, 30280, 30340, 30710, 30780, 30800, 30810, 30820, 30920]
   int[] femaleHair = [31000, 31030, 31100, 31350, 31460, 31550, 31770, 31790, 31850]
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
            cm.sendSimple("9120100_WELCOME")

         } else if (status == 1) {
            if (selection == 1) {
               beauty = 1
               hairNew = []
               if (cm.getGender() == 0) {
                  for (int i = 0; i < maleHair.length; i++) {
                     hairNew = ScriptUtils.pushItemIfTrue(hairNew, maleHair[i] + (cm.getHair() % 10).intValue(), { itemId -> cm.cosmeticExistsAndIsNotEquipped(itemId) })
                  }
               }
               if (cm.getGender() == 1) {
                  for (int i = 0; i < femaleHair.length; i++) {
                     hairNew = ScriptUtils.pushItemIfTrue(hairNew, femaleHair[i] + (cm.getHair() % 10).intValue(), { itemId -> cm.cosmeticExistsAndIsNotEquipped(itemId) })
                  }
               }
               cm.sendStyle("I can totally change up your hairstyle and make it look so good. Why don't you change it up a bit? With #b#t5150009##k, I'll take care of the rest for you. Choose the style of your liking!", hairNew)
            } else if (selection == 2) {
               beauty = 2
               hairColor = []
               int current = (cm.getHair() / 10).intValue() * 10
               for (int i = 0; i < 8; i++) {
                  hairColor = ScriptUtils.pushItemIfTrue(hairColor, current + i, { itemId -> cm.cosmeticExistsAndIsNotEquipped(itemId) })
               }
               cm.sendStyle("I can totally change your hair color and make it look so good. Why don't you change it up a bit? With #b#t5151009##k, I'll take care of the rest. Choose the color of your liking!", hairColor)
            }
         } else if (status == 2) {
            cm.dispose()
            if (beauty == 1) {
               if (cm.haveItem(5150009)) {
                  cm.gainItem(5150009, (short) -1)
                  cm.setHair(hairNew[selection])
                  cm.sendOk("9120100_ENJOY_NEW_STYLE")

               } else {
                  cm.sendOk("9120100_MISSING_STYLE_COUPON")

               }
            }
            if (beauty == 2) {
               if (cm.haveItem(5151009)) {
                  cm.gainItem(5151009, (short) -1)
                  cm.setHair(hairColor[selection])
                  cm.sendOk("9120100_ENJOY_NEW_COLOR")

               } else {
                  cm.sendOk("9120100_MISSING_COLOR_COUPON")

               }
            }
            if (beauty == 0) {
               if (selection == 0 && cm.getMeso() >= hairPrice) {
                  cm.gainMeso(-hairPrice)
                  cm.gainItem(5150009, (short) 1)
                  cm.sendOk("9120100_ENJOY")

               } else if (selection == 1 && cm.getMeso() >= hairColorPrice) {
                  cm.gainMeso(-hairColorPrice)
                  cm.gainItem(5151009, (short) 1)
                  cm.sendOk("9120100_ENJOY")

               } else {
                  cm.sendOk("9120100_NOT_ENOUGH_MESOS")

               }
            }
         }
      }
   }
}

NPC9120100 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9120100(cm: cm))
   }
   return (NPC9120100) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
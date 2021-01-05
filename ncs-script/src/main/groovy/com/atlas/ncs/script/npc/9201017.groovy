package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager
import com.atlas.ncs.util.ScriptUtils

class NPC9201017 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int beauty = 0
   int regularPrice = 1000000
   int vipPrice = 1000000
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
            cm.sendSimple("9201017_HELLO")

         } else if (status == 1) {
            if (selection == 1) {
               beauty = 1
               int current = 0
               if (cm.getGender() == 0) {
                  current = cm.getFace() % 100 + 20000
               }
               if (cm.getGender() == 1) {
                  current = cm.getFace() % 100 + 21000
               }
               int[] temp = [current, current + 100, current + 300, current + 400, current + 500, current + 700]
               colors = ScriptUtils.pushItemsIfTrue(colors, temp, { itemId -> cm.cosmeticExistsAndIsNotEquipped(itemId) })
               cm.sendYesNo("9201017_REG_CONFIRM")

            } else if (selection == 2) {
               beauty = 2
               int current = 0
               if (cm.getGender() == 0) {
                  current = cm.getFace() % 100 + 20000
               }
               if (cm.getGender() == 1) {
                  current = cm.getFace() % 100 + 21000
               }
               int[] temp = [current, current + 100, current + 300, current + 400, current + 500, current + 700]
               colors = ScriptUtils.pushItemsIfTrue(colors, temp, { itemId -> cm.cosmeticExistsAndIsNotEquipped(itemId) })
               cm.sendStyle("With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking.", colors)
            } else if (selection == 3) {
               beauty = 3
               int current = 0
               if (cm.getGender() == 0) {
                  current = cm.getFace() % 100 + 20000
               }
               if (cm.getGender() == 1) {
                  current = cm.getFace() % 100 + 21000
               }

               colors = []
               for (int i = 0; i < 8; i++) {
                  if (cm.haveItem(5152100 + i)) {
                     colors = ScriptUtils.pushItemIfTrue(colors, current + 100 * i, { itemId -> cm.cosmeticExistsAndIsNotEquipped(itemId) })
                  }
               }

               if (colors.length == 0) {
                  cm.sendOk("9201017_NO_ONE_TIME_COSMETIC_LENS")

                  cm.dispose()
                  return
               }

               cm.sendStyle("What kind of lens would you like to wear? Please choose the style of your liking.", colors)
            }
         } else if (status == 2) {
            if (beauty == 1) {
               if (cm.haveItem(5152025)) {
                  cm.gainItem(5152025, (short) -1)
                  cm.setFace(colors[Math.floor(Math.random() * colors.length).intValue()])
                  cm.sendOk("9201017_ENJOY_NEW_LENS")

               } else {
                  cm.sendOk("9201017_MISSING_LENS_COUPON")

                  cm.dispose()
               }
            } else if (beauty == 2) {
               if (cm.haveItem(5152026)) {
                  cm.gainItem(5152026, (short) -1)
                  cm.setFace(colors[selection])
                  cm.sendOk("9201017_ENJOY_NEW_LENS")

               } else {
                  cm.sendOk("9201017_MISSING_LENS_COUPON")

                  cm.dispose()
               }
            } else if (beauty == 3) {
               int color = (colors[selection] / 100) % 100 | 0

               if (cm.haveItem(5152100 + color)) {
                  cm.gainItem(5152100 + color, (short) -1)
                  cm.setFace(colors[selection])
                  cm.sendOk("9201017_ENJOY_NEW_LENS")

               } else {
                  cm.sendOk("9201017_MISSING_LENS_COUPON")

               }
            } else if (beauty == 0) {
               if (selection == 0 && cm.getMeso() >= regularPrice) {
                  cm.gainMeso(-regularPrice)
                  cm.gainItem(5152025, (short) 1)
                  cm.sendOk("9201017_ENJOY")

                  cm.dispose()
               } else if (selection == 1 && cm.getMeso() >= vipPrice) {
                  cm.gainMeso(-vipPrice)
                  cm.gainItem(5152026, (short) 1)
                  cm.sendOk("9201017_ENJOY")

                  cm.dispose()
               } else {
                  cm.sendOk("9201017_NOT_ENOUGH_MESOS")

               }
            }
         }
      }
   }
}

NPC9201017 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201017(cm: cm))
   }
   return (NPC9201017) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
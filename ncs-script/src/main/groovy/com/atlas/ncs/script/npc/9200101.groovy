package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9200101 {
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
            cm.sendSimple("Hello, I'm Dr. Rhomes, head of the cosmetic lens department here at the Orbis Plastic Surgery Shop.\r\nMy goal here is to add personality to everyone's eyes through the wonders of cosmetic lenses, and with #b#t5152011##k or #b#t5152014##k, I can do the same for you, too! Now, what would you like to use?\r\n#L1#Cosmetic Lenses: #i5152011##t5152011##l\r\n#L2#Cosmetic Lenses: #i5152014##t5152014##l\r\n#L3#One-time Cosmetic Lenses: #i5152104# (any color)#l")
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
               int[] temp = [current + 100, current + 300, current + 400, current + 700]
               colors = ScriptUtils.pushItemsIfTrue(colors, temp, { itemId -> cm.cosmeticExistsAndIsntEquipped(itemId) })
               cm.sendYesNo("If you use the regular coupon, you'll be awarded a random pair of cosmetic lenses. Are you going to use a #b#t5152011##k and really make the change to your eyes?")
            } else if (selection == 2) {
               beauty = 2
               int current = 0
               if (cm.getGender() == 0) {
                  current = cm.getFace() % 100 + 20000
               }
               if (cm.getGender() == 1) {
                  current = cm.getFace() % 100 + 21000
               }
               int[] temp = [current + 100, current + 300, current + 400, current + 700]
               colors = ScriptUtils.pushItemsIfTrue(colors, temp, { itemId -> cm.cosmeticExistsAndIsntEquipped(itemId) })
               cm.sendStyle("With our new computer program, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Please choose the style of your liking.", colors)
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
                     colors = ScriptUtils.pushItemIfTrue(colors, current + 100 * 1, { itemId -> cm.cosmeticExistsAndIsntEquipped(itemId) })
                  }
               }

               if (colors.length == 0) {
                  cm.sendOk("You don't have any One-Time Cosmetic Lens to use.")
                  cm.dispose()
                  return
               }

               cm.sendStyle("What kind of lens would you like to wear? Please choose the style of your liking.", colors)
            }
         } else if (status == 2) {
            cm.dispose()
            if (beauty == 1) {
               if (cm.haveItem(5152011)) {
                  cm.gainItem(5152011, (short) -1)
                  cm.setFace(colors[Math.floor(Math.random() * colors.length).intValue()])
                  cm.sendOk("Enjoy your new and improved cosmetic lenses!")
               } else {
                  cm.sendOk("I'm sorry, but I don't think you have our cosmetic lens coupon with you right now. Without the coupon, I'm afraid I can't do it for you..")
               }
            } else if (beauty == 2) {
               if (cm.haveItem(5152014)) {
                  cm.gainItem(5152014, (short) -1)
                  cm.setFace(colors[selection])
                  cm.sendOk("Enjoy your new and improved cosmetic lenses!")
               } else {
                  cm.sendOk("I'm sorry, but I don't think you have our cosmetic lens coupon with you right now. Without the coupon, I'm afraid I can't do it for you..")
               }
            } else if (beauty == 3) {
               int color = (colors[selection] / 100) % 100 | 0

               if (cm.haveItem(5152100 + color)) {
                  cm.gainItem(5152100 + color, (short) -1)
                  cm.setFace(colors[selection])
                  cm.sendOk("Enjoy your new and improved cosmetic lenses!")
               } else {
                  cm.sendOk("I'm sorry, but I don't think you have our cosmetic lens coupon with you right now. Without the coupon, I'm afraid I can't do it for you..")
               }
            } else if (beauty == 0) {
               if (selection == 0 && cm.getMeso() >= regularPrice) {
                  cm.gainMeso(-regularPrice)
                  cm.gainItem(5152011, (short) 1)
                  cm.sendOk("Enjoy!")
               } else if (selection == 1 && cm.getMeso() >= vipPrice) {
                  cm.gainMeso(-vipPrice)
                  cm.gainItem(5152014, (short) 1)
                  cm.sendOk("Enjoy!")
               } else {
                  cm.sendOk("You don't have enough mesos to buy a coupon!")
               }
            }
         }
      }
   }
}

NPC9200101 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9200101(cm: cm))
   }
   return (NPC9200101) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
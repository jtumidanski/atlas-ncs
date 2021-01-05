package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager
import com.atlas.ncs.util.ScriptUtils

class NPC2012009 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int beauty = 0
   int price = 1000000
   int[] maleFace = [20003, 20011, 20021, 20022, 20023, 20027, 20031]
   int[] femaleFace = [21004, 21007, 21010, 21012, 21020, 21021, 21030]
   int[] faceNew = []

   def start() {
      cm.sendSimple("2012009_HELLO")
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode < 1) {
         cm.dispose()
      } else {
         status++
         if (status == 1) {
            if (selection == 2) {
               faceNew = []
               if (cm.getGender() == 0) {
                  for (int i = 0; i < maleFace.length; i++) {
                     faceNew = ScriptUtils.pushItemIfTrue(faceNew, maleFace[i] + cm.getFace() % 1000 - (cm.getFace() % 100), { itemId -> cm.cosmeticExistsAndIsNotEquipped(itemId) })
                  }
               } else {
                  for (int i = 0; i < femaleFace.length; i++) {
                     faceNew = ScriptUtils.pushItemIfTrue(faceNew, femaleFace[i] + cm.getFace() % 1000 - (cm.getFace() % 100), { itemId -> cm.cosmeticExistsAndIsNotEquipped(itemId) })
                  }
               }
               cm.sendYesNo("2012009_REGULAR_COUPON_EXPLANATION")
            }
         } else if (status == 2) {
            if (cm.haveItem(5152004)) {
               cm.gainItem(5152004, (short) -1)
               cm.setFace(faceNew[Math.floor(Math.random() * faceNew.length).intValue()])
               cm.sendOk("2012009_ENJOY_NEW_FACE")
            } else {
               cm.sendOk("2012009_MISSING_FACE_COUPON")
            }

            cm.dispose()
         }
      }
   }
}

NPC2012009 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2012009(cm: cm))
   }
   return (NPC2012009) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager
import com.atlas.ncs.util.ScriptUtils

class NPC9201018 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int beauty = 0
   int price = 1000000
   int[] maleFace = [20000, 20001, 20003, 20004, 20005, 20006, 20007, 20008, 20018, 20019]
   int[] femaleFace = [21001, 21002, 21003, 21004, 21005, 21006, 21007, 21012, 21018, 21019]
   int[] faceNew = []

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
            cm.sendSimple("9201018_HELLO")

         } else if (status == 1) {
            if (selection == 2) {
               faceNew = []
               if (cm.getGender() == 0) {
                  for (int i = 0; i < maleFace.length; i++) {
                     faceNew = ScriptUtils.pushItemIfTrue(faceNew, maleFace[i] + cm.getFace() % 1000 - (cm.getFace() % 100), { itemId -> cm.cosmeticExistsAndIsNotEquipped(itemId) })
                  }
               }
               if (cm.getGender() == 1) {
                  for (int i = 0; i < femaleFace.length; i++) {
                     faceNew = ScriptUtils.pushItemIfTrue(faceNew, femaleFace[i] + cm.getFace() % 1000 - (cm.getFace() % 100), { itemId -> cm.cosmeticExistsAndIsNotEquipped(itemId) })
                  }
               }
               cm.sendStyle("Let's see... I can totally transform your face into something new. Don't you want to try it? For #b#t5152022##k, you can get the face of your liking. Take your time in choosing the face of your preference.", faceNew)
            }
         } else if (status == 2) {
            if (cm.haveItem(5152022)) {
               cm.gainItem(5152022, (short) -1)
               cm.setFace(faceNew[selection])
               cm.sendOk("9201018_ENJOY_NEW_FACE")

            } else {
               cm.sendOk("9201018_MISSING_FACE_COUPON")

               cm.dispose()
            }
         }
      }
   }
}

NPC9201018 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201018(cm: cm))
   }
   return (NPC9201018) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
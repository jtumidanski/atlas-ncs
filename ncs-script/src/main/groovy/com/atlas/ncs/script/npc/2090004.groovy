package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2090004 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int selectedType = -1
   int selectedItem = -1
   Object item
   Object mats
   Object matQty
   int matMeso
   List rewardSet
   int makeQty = 1

   Object[] itemSet
   List matSet
   List matQtySet
   int[] matQtyMeso

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == 1) {
         status++
      } else {
         cm.sendOk("2090004_I_AM_VERY_BUSY")

         cm.dispose()
         return
      }

      if (status == 0) {
         if (cm.isQuestActive(3821) && !cm.haveItem(4031554) && !cm.haveItem(4161030) && cm.isQuestCompleted(3830)) {
            //player lost his book, help him complete quest anyways

            if (cm.canHold(4031554)) {
               cm.sendOk("2090004_OH_BOY")

               cm.gainItem(4031554, (short) 1)
               cm.dispose()
               return
            } else {
               cm.sendOk("2090004_OH_BOY_MAKE_ETC_ROOM")

               cm.dispose()
               return
            }
         }
         String selStr = "I am a man of many talents. Let me know what you'd like to do. #b"
         String[] options = ["Make a medicine", "Make a scroll", "Donate medicine ingredients"]
         for (int i = 0; i < options.length; i++) {
            selStr += "\r\n#L" + i + "# " + options[i] + "#l"
         }

         cm.sendSimple(selStr)
      } else if (status == 1) {
         selectedType = selection
         String selStr
         if (selectedType == 0) { //Make a medicine
            itemSet = [2022145, 2022146, 2022147, 2022148, 2022149, 2022150, 2050004, 4031554]
            matSet = [2022116, 2022116, [4000281, 4000293], [4000276, 2002005], [4000288, 4000292], 4000295, [2022131, 2022132], [4000286, 4000287, 4000293]]
            matQtySet = [3, 3, [10, 10], [20, 1], [20, 20], 10, [1, 1], [20, 20, 20]]
            matQtyMeso = [0, 0, 910, 950, 1940, 600, 700, 1000]

            if (!cm.haveItem(4161030)) {
               cm.sendNext("2090004_STUDY_THE_BOOK")

               cm.dispose()
               return
            }

            selStr = "What kind of medicine are you interested in making?#b"

            for (int i = 0; i < itemSet.length; i++) {
               selStr += "\r\n#L" + i + "# #v" + itemSet[i] + "# #t" + itemSet[i] + "##l"
            }
            selStr += "#k"
         } else if (selectedType == 1) { //Make a scroll
            status++

            selStr = "What kind of scrolls are you interested in making?#b"
            itemSet = ["Scroll for One-Handed Sword for ATT", "Scroll for One-Handed Axe for ATT", "Scroll for One-Handed BW for ATT",
                       "Scroll for Dagger for ATT", "Scroll for Wand for Magic Att.", "Scroll for Staff for Magic Att.",
                       "Scroll for Two-handed Sword for ATT.", "Scroll for Two-handed Axe for ATT", "Scroll for Two-handed BW for ATT",
                       "Scroll for Spear for ATT", "Scroll for Pole Arm for ATT", "Scroll for Bow for ATT", "Scroll for Crossbow for ATT ",
                       "Scroll for Claw for ATT", "Scroll for Knuckle for ATT", "Scroll for Gun for ATT#k"]

            for (int i = 0; i < itemSet.length; i++) {
               selStr += "\r\n#L" + i + "# " + itemSet[i] + "#l"
            }
         } else {//Donate medicine ingredients
            status++

            selStr = "So you wish to donate some medicine ingredients? This is great news! Donations will be accepted in the unit of #b100#k. The donator will receive a marble that enables one to make a scroll. Which of these would you like to donate? #b"
            itemSet = [4000276, 4000277, 4000278, 4000279, 4000280, 4000291, 4000292, 4000286, 4000287, 4000293, 4000294, 4000298, 4000284, 4000288, 4000285, 4000282, 4000295, 4000289, 4000296, 4000297]

            for (int i = 0; i < itemSet.length; i++) {
               selStr += "\r\n#L" + i + "# #v" + itemSet[i] + "# #t" + itemSet[i] + "##l"
            }
         }

         cm.sendSimple(selStr)
      } else if (status == 2) {
         selectedItem = selection
         cm.sendGetText("2090004_HOW_MANY_TO_MAKE", itemSet[selectedItem])

      } else if (status == 3) {
         if (selectedType == 0) { //Medicines
            String text = cm.getText()
            makeQty = (text).toInteger()
            if (makeQty == null) {
               makeQty = 1
            }

            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            matMeso = matQtyMeso[selectedItem]

            String prompt = "You want to make #b" + makeQty + " #t" + item + "##k? In order to make " + makeQty + " #t" + item + "#, you'll need the following items:\r\n"
            if (mats instanceof ArrayList && matQty instanceof ArrayList) {
               for (int i = 0; i < mats.size(); i++) {
                  prompt += "\r\n#i" + mats[i] + "# " + ((matQty[i] as Integer) * makeQty) + " #t" + mats[i] + "#"
               }
            } else {
               prompt += "\r\n#i" + mats + "# " + ((matQty as Integer) * makeQty) + " #t" + mats + "#"
            }

            if (matMeso > 0) {
               prompt += "\r\n#i4031138# " + matMeso * makeQty + " meso"
            }

            cm.sendYesNo(prompt)
         } else if (selectedType == 1) { //Scrolls
            selectedItem = selection

            itemSet = [2043000, 2043100, 2043200, 2043300, 2043700, 2043800, 2044000, 2044100, 2044200, 2044300, 2044400, 2044500, 2044600, 2044700, 2044800, 2044900]
            matSet = [[4001124, 4010001], [4001124, 4010001], [4001124, 4010001], [4001124, 4010001], [4001124, 4010001],
                      [4001124, 4010001], [4001124, 4010001], [4001124, 4010001], [4001124, 4010001], [4001124, 4010001], [4001124, 4010001],
                      [4001124, 4010001], [4001124, 4010001], [4001124, 4010001], [4001124, 4010001], [4001124, 4010001]]
            matQtySet = [[100, 10], [100, 10], [100, 10], [100, 10], [100, 10], [100, 10], [100, 10],
                         [100, 10], [100, 10], [100, 10], [100, 10], [100, 10], [100, 10], [100, 10], [100, 10],
                         [100, 10]]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            String prompt = "You want to make #b#t" + item + "##k? In order to make #t" + item + "# you'll need the following items:"
            if (mats instanceof ArrayList && matQty instanceof ArrayList) {
               for (int i = 0; i < mats.size(); i++) {
                  prompt += "\r\n#i" + mats[i] + "# " + matQty[i] + " #t" + mats[i] + "#"
               }
            } else {
               prompt += "\r\n#i" + mats + "# " + matQty + " #t" + mats + "#"
            }

            cm.sendYesNo(prompt)
         } else if (selectedType == 2) {
            selectedItem = selection

            itemSet = [4000276, 4000277, 4000278, 4000279, 4000280, 4000291, 4000292, 4000286, 4000287, 4000293, 4000294, 4000298, 4000284, 4000288, 4000285, 4000282, 4000295, 4000289, 4000296, 4000297]
            rewardSet = [7, 7, [7, 8], 10, 11, 8, [7, 8], [7, 9], [7, 8], 9, 10, [10, 11], 11, [11, 12], 13, 13, 14, 15, [15, 16], 17]

            item = itemSet[selectedItem]
            String prompt = "Are you sure you want to donate #b100 #t " + item + "##k?"
            cm.sendYesNo(prompt)
         }
      } else if (status == 4) {
         if (selectedType == 0) {
            boolean complete = true
            if (mats instanceof ArrayList && matQty instanceof ArrayList) {
               for (int i = 0; i < mats.size(); i++) {
                  if (!cm.haveItem(mats[i] as Integer, (matQty[i] as Integer) * makeQty)) {
                     complete = false
                  }
               }
            } else {
               if (!cm.haveItem(mats as Integer, (matQty as Integer) * makeQty)) {
                  complete = false
               }
            }

            if (cm.getMeso() < matMeso * makeQty) {
               complete = false
            }

            if (!complete || !cm.canHold(item as Integer, makeQty)) {
               cm.sendOk("2090004_NEED_INGREDIENTS_OR_INVENTORY_SPACE")

            } else {
               if (mats instanceof ArrayList && matQty instanceof ArrayList) {
                  for (int i = 0; i < mats.size(); i++) {
                     cm.gainItem(mats[i] as Integer, (short) ((-matQty[i] as Integer) * makeQty))
                  }
               } else {
                  cm.gainItem(mats as Integer, (short) ((-matQty as Integer) * makeQty))
               }

               if (matMeso > 0) {
                  cm.gainMeso(-matMeso * makeQty)
               }
               cm.gainItem(item as Integer, (short) makeQty)
            }

            cm.dispose()
         } else if (selectedType == 1) {
            boolean complete = true
            if (mats instanceof ArrayList && matQty instanceof ArrayList) {
               for (int i = 0; i < mats.size(); i++) {
                  if (!cm.haveItem(mats[i] as Integer, matQty[i] as Integer)) {
                     complete = false
                  }
               }
            } else {
               if (!cm.haveItem(mats as Integer, matQty as Integer)) {
                  complete = false
               }
            }

            if (Math.random() >= 0.9) //A lucky find! Scroll 60%
            {
               item += 1
            }

            if (!complete || !cm.canHold(item as Integer, 1)) {
               cm.sendOk("2090004_NEED_INGREDIENTS_OR_INVENTORY_SPACE")

            } else {
               if (mats instanceof ArrayList && matQty instanceof ArrayList) {
                  for (int i = 0; i < mats.size(); i++) {
                     cm.gainItem(mats[i] as Integer, (short) (-matQty[i] as Integer))
                  }
               } else {
                  cm.gainItem(mats as Integer, (short) (-matQty as Integer))
               }

               cm.gainItem(item as Integer, (short) 1)
            }

            cm.dispose()
         } else if (selectedType == 2) {
            boolean complete = true

            if (!cm.haveItem(item as Integer, 100)) {
               complete = false
            }

            if (!complete) {
               cm.sendOk("2090004_NEED_INGREDIENTS_OR_INVENTORY_SPACE")

               cm.dispose()
               return
            }

            int reward
            if (rewardSet[selectedItem] instanceof ArrayList) {
               List selectedReward = (List) rewardSet[selectedItem]
               int length = (selectedReward[1] as Integer) - (selectedReward[0] as Integer)
               reward = (selectedReward[0] as Integer) + Math.round(Math.random() * length).intValue()
            } else {
               reward = rewardSet[selectedItem] as Integer
            }

            if (!cm.canHold(4001124, reward)) {
               cm.sendOk("2090004_NEED_INGREDIENTS_OR_INVENTORY_SPACE")

            } else {
               cm.gainItem(item as Integer, (short) -100)
               cm.gainItem(4001124, (short) reward)
            }

            cm.dispose()
         }
      }
   }
}

NPC2090004 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2090004(cm: cm))
   }
   return (NPC2090004) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
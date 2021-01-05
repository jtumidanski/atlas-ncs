package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2040021 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int selectedType = -1
   int selectedItem = -1
   int item
   Object mats
   Object matQty
   int cost
   int qty
   boolean equip

   boolean stimulator = false
   int stimulatorId = 4130001

   def start() {
      String selStr = "Hello, and welcome to the Ludibrium Shoe Store. How can I help you today??#b"
      String[] options = ["What's a stimulator?", "Create Warrior shoes", "Create Bowman shoes", "Create Magician shoes", "Create Thief shoes", "Create Warrior shoes with a Stimulator", "Create Bowman shoes with a Stimulator", "Create Magician shoes with a Stimulator", "Create Thief shoes with a Stimulator"]
      for (int i = 0; i < options.length; i++) {
         selStr += "\r\n#L" + i + "# " + options[i] + "#l"
      }
      cm.sendSimple(selStr)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode < 1) {
         cm.dispose()
         return
      }
      status++
      if (status == 1) {
         selectedType = selection
         String selStr = ""
         String[] shoes = []
         if (selectedType > 4) {
            stimulator = true
            selectedType -= 4
         } else {
            stimulator = false
         }
         if (selectedType == 0) { // what is stimulator
            cm.sendNext("2040021_STIMULATOR_INFO")
            cm.dispose()
         } else if (selectedType == 1) { //warrior shoe
            selStr = "Warrior shoes? Sure thing, which kind?#b"
            shoes = ["Emerald Battle Grieve#k - Warrior Lv. 30#b", "Mithril Battle Grieve#k - Warrior Lv. 30#b", "Silver Battle Grieve#k - Warrior Lv. 30#b", "Blood Battle Grieve#k - Warrior Lv. 30#b", "Steel Trigger#k - Warrior Lv. 35#b", "Mithril Trigger#k - Warrior Lv. 35#b", "Dark Trigger#k - Warrior Lv. 35#b", "Brown Jangoon Boots#k - Warrior Lv. 40#b", "Maroon Jangoon Boots#k - Warrior Lv. 40#b", "Blue Jangoon Boots#k - Warrior Lv. 40#b", "Emerald Hildon Boots#k - Warrior Lv. 50#b", "Mithril Hildon Boots#k - Warrior Lv. 50#b", "Orihalcon Hildon Boots#k - Warrior Lv. 50#b", "Gold Hildon Boots#k - Warrior Lv. 50#b"]
         } else if (selectedType == 2) { //bowman shoe
            selStr = "Bowman shoes? Sure thing, which kind?#b"
            shoes = ["Red Hunter Boots#k - Bowman Lv. 30#b", "Blue Hunter Boots#k - Bowman Lv. 30#b", "Green Hunter Boots#k - Bowman Lv. 30#b", "Black Hunter Boots#k - Bowman Lv. 30#b", "Brown Hunter Boots#k - Bowman Lv. 30#b", "Blue Silky Boots#k - Bowman Lv. 35#b", "Green Silky Boots#k - Bowman Lv. 35#b", "Red Silky Boots#k - Bowman Lv. 35#b", "Red Pierre Shoes#k - Bowman Lv. 40#b", "Yellow Pierre Shoes#k - Bowman Lv. 40#b", "Brown Pierre Shoes#k - Bowman Lv. 40#b", "Blue Pierre Shoes#k - Bowman Lv. 40#b", "Brown Steel-Tipped Boots#k - Bowman Lv. 50#b", "Green Steel-Tipped Boots#k - Bowman Lv. 50#b", "Blue Steel-Tipped Boots#k - Bowman Lv. 50#b", "Purple Steel-Tipped Boots#k - Bowman Lv. 50#b"]
         } else if (selectedType == 3) { //magician shoe
            selStr = "Magician shoes? Sure thing, which kind?#b"
            shoes = ["Red Magicshoes#k - Magician Lv. 30#b", "Blue Magicshoes#k - Magician Lv. 30#b", "White Magicshoes#k - Magician Lv. 30#b", "Black Magicshoes#k - Magician Lv. 30#b", "Purple Salt Shoes#k - Magician Lv. 35#b", "Red Salt Shoes#k - Magician Lv. 35#b", "Black Salt Shoes#k - Magician Lv. 35#b", "Red Moon Shoes#k - Magician Lv. 40#b", "Blue Moon Shoes#k - Magician Lv. 40#b", "Gold Moon Shoes#k - Magician Lv. 40#b", "Dark Moon Shoes#k - Magician Lv. 40#b", "Pink Goldwind Shoes#k - Magician Lv. 50#b", "Blue Goldwind Shoes#k - Magician Lv. 50#b", "Purple Goldwind Shoes#k - Magician Lv. 50#b", "Green Goldwind Shoes#k - Magician Lv. 50#b"]
         } else if (selectedType == 4) { //thief shoe
            selStr = "Thief shoes? Sure thing, which kind?#b"
            shoes = ["Bronze Chain Boots#k - Thief Lv. 30#b", "Iron Chain Boots#k - Thief Lv. 30#b", "Silver Chain Boots#k - Thief Lv. 30#b", "Gold Chain Boots#k - Thief Lv. 30#b", "Red White-Line Boots#k - Thief Lv. 35#b", "Green White-Line Boots#k - Thief Lv. 35#b", "Blue White-Line Boots#k - Thief Lv. 35#b", "Black Red-Lined Shoes#k - Thief Lv. 40#b", "Black Green-Lined Shoes#k - Thief Lv. 40#b", "Black Yellow-Lined Shoes#k - Thief Lv. 40#b", "Black Blue-Lined Shoes#k - Thief Lv. 40#b", "Blue Goni Shoes#k - Thief Lv. 50#b", "Red Goni Shoes#k - Thief Lv. 50#b", "Green Goni Shoes#k - Thief Lv. 50#b", "Purple Goni Shoes#k - Thief Lv. 50#b"]
         }
         if (selectedType != 0) {
            for (int i = 0; i < shoes.length; i++) {
               selStr += "\r\n#L" + i + "# " + shoes[i] + "#l"
            }
            cm.sendSimple(selStr)
         }
      } else if (status == 2) {
         selectedItem = selection
         if (selectedType == 1) { //warrior shoe
            int[] itemSet = [1072003, 1072039, 1072040, 1072041, 1072002, 1072112, 1072113, 1072000, 1072126, 1072127, 1072132, 1072133, 1072134, 1072135]
            List matSet = [[4021003, 4011001, 4000021, 4003000], [4011002, 4011001, 4000021, 4003000],
                           [4011004, 4011001, 4000021, 4003000], [4021000, 4011001, 4000021, 4003000], [4011001, 4021004, 4000021, 4000030, 4003000], [4011002, 4021004, 4000021, 4000030, 4003000], [4021008, 4021004, 4000021, 4000030, 4003000],
                           [4011003, 4000021, 4000030, 4003000, 4000103], [4011005, 4021007, 4000030, 4003000, 4000104], [4011002, 4021007, 4000030, 4003000, 4000105], [4021008, 4011001, 4021003, 4000030, 4003000],
                           [4021008, 4011001, 4011002, 4000030, 4003000], [4021008, 4011001, 4011005, 4000030, 4003000], [4021008, 4011001, 4011006, 4000030, 4003000]]
            List matQtySet = [[4, 2, 45, 15], [4, 2, 45, 15], [4, 2, 45, 15], [4, 2, 45, 15], [3, 1, 30, 20, 25], [3, 1, 30, 20, 25], [2, 1, 30, 20, 25],
                              [4, 100, 40, 30, 100], [4, 1, 40, 30, 100], [4, 1, 40, 30, 100], [1, 3, 6, 65, 45], [1, 3, 6, 65, 45], [1, 3, 6, 65, 45], [1, 3, 6, 65, 45]]
            int[] costSet = [20000, 20000, 20000, 20000, 22000, 22000, 25000, 38000, 38000, 38000, 50000, 50000, 50000, 50000]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
         } else if (selectedType == 2) { //bowman shoe
            int[] itemSet = [1072079, 1072080, 1072081, 1072082, 1072083, 1072101, 1072102, 1072103, 1072118, 1072119, 1072120, 1072121, 1072122, 1072123, 1072124, 1072125]
            List matSet = [[4000021, 4021000, 4003000], [4000021, 4021005, 4003000], [4000021, 4021003, 4003000],
                           [4000021, 4021004, 4003000], [4000021, 4021006, 4003000], [4021002, 4021006, 4000030, 4000021, 4003000], [4021003, 4021006, 4000030, 4000021, 4003000], [4021000, 4021006, 4000030, 4000021, 4003000],
                           [4021000, 4003000, 4000030, 4000106], [4021006, 4003000, 4000030, 4000107], [4011003, 4003000, 4000030, 4000108], [4021002, 4003000, 4000030, 4000099], [4011001, 4021006, 4021008, 4000030, 4003000, 4000033],
                           [4011001, 4021006, 4021008, 4000030, 4003000, 4000032], [4011001, 4021006, 4021008, 4000030, 4003000, 4000041], [4011001, 4021006, 4021008, 4000030, 4003000, 4000042]]
            List matQtySet = [[50, 2, 15], [50, 2, 15], [50, 2, 15], [50, 2, 15], [50, 2, 15],
                              [3, 1, 15, 30, 20], [3, 1, 15, 30, 20], [3, 1, 15, 30, 20], [4, 30, 45, 100], [4, 30, 45, 100], [5, 30, 45, 100], [5, 30, 45, 100],
                              [3, 3, 1, 60, 35, 80], [3, 3, 1, 60, 35, 150], [3, 3, 1, 60, 35, 100], [3, 3, 1, 60, 35, 250]]
            int[] costSet = [19000, 19000, 19000, 19000, 19000, 19000, 20000, 20000, 20000, 32000, 32000, 40000, 40000, 50000, 50000, 50000, 50000]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
         } else if (selectedType == 3) { //magician shoe
            int[] itemSet = [1072075, 1072076, 1072077, 1072078, 1072089, 1072090, 1072091, 1072114, 1072115, 1072116, 1072117, 1072140, 1072141, 1072142, 1072143, 1072136, 1072137, 1072138, 1072139]
            List matSet = [[4021000, 4000021, 4003000], [4021002, 4000021, 4003000], [4011004, 4000021, 4003000], [4021008, 4000021, 4003000], [4021001, 4021006, 4000021, 4000030, 4003000], [4021000, 4021006, 4000021, 4000030, 4003000],
                           [4021008, 4021006, 4000021, 4000030, 4003000], [4021000, 4000030, 4000110, 4003000], [4021005, 4000030, 4000111, 4003000], [4011006, 4021007, 4000030, 4000100, 4003000], [4021008, 4021007, 4000030, 4000112, 4003000],
                           [4021009, 4011006, 4021000, 4000030, 4003000], [4021009, 4011006, 4021005, 4000030, 4003000], [4021009, 4011006, 4021001, 4000030, 4003000], [4021009, 4011006, 4021003, 4000030, 4003000]]
            List matQtySet = [[2, 50, 15], [2, 50, 15], [2, 50, 15], [1, 50, 15], [3, 1, 30, 15, 20], [3, 1, 30, 15, 20], [2, 1, 40, 25, 20], [4, 40, 100, 25], [4, 40, 100, 25], [2, 1, 40, 100, 25], [2, 1, 40, 100, 30],
                              [1, 3, 3, 60, 40], [1, 3, 3, 60, 40], [1, 3, 3, 60, 40], [1, 3, 3, 60, 40]]
            int[] costSet = [18000, 18000, 18000, 18000, 20000, 20000, 22000, 30000, 30000, 35000, 40000, 50000, 50000, 50000, 50000]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
         } else if (selectedType == 4) { //thief shoe
            int[] itemSet = [1072032, 1072033, 1072035, 1072036, 1072104, 1072105, 1072106, 1072107, 1072108, 1072109, 1072110, 1072128, 1072130, 1072129, 1072131]
            List matSet = [[4011000, 4000021, 4003000], [4011001, 4000021, 4003000], [4011004, 4000021, 4003000], [4011006, 4000021, 4003000], [4021000, 4021004, 4000021, 4000030, 4003000], [4021003, 4021004, 4000021, 4000030, 4003000],
                           [4021002, 4021004, 4000021, 4000030, 4003000], [4021000, 4000030, 4000113, 4003000], [4021003, 4000030, 4000095, 4003000], [4021006, 4000030, 4000096, 4003000], [4021005, 4000030, 4000097, 4003000], [4011007, 4021005, 4000030, 4000114, 4003000],
                           [4011007, 4021000, 4000030, 4000115, 4003000], [4011007, 4021003, 4000030, 4000109, 4003000], [4011007, 4021001, 4000030, 4000036, 4003000]]
            List matQtySet = [[3, 50, 15], [3, 50, 15], [2, 50, 15], [2, 50, 15], [3, 1, 30, 15, 20], [3, 1, 30, 15, 20], [3, 1, 30, 15, 20],
                              [5, 45, 100, 30], [4, 45, 100, 30], [4, 45, 100, 30], [4, 45, 100, 30], [2, 3, 50, 100, 35], [2, 3, 50, 100, 35], [2, 3, 50, 100, 35], [2, 3, 50, 80, 35]]
            int[] costSet = [19000, 19000, 19000, 21000, 20000, 20000, 20000, 40000, 32000, 35000, 35000, 50000, 50000, 50000, 50000]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
         }
         //Ludi fee is -10%, array not changed unlike 2040016 and 2040020
         cost *= 0.9
         String prompt = "You want me to make a #t" + item + "#? In that case, I'm going to need specific items from you in order to make it. Make sure you have room in your inventory, though!#b"
         if (stimulator) {
            prompt += "\r\n#i" + stimulatorId + "# 1 #t" + stimulatorId + "#"
         }
         if (mats instanceof ArrayList && matQty instanceof ArrayList) {
            for (int i = 0; i < mats.size(); i++) {
               prompt += "\r\n#i" + mats[i] + "# " + matQty[i] + " #t" + mats[i] + "#"
            }
         } else {
            prompt += "\r\n#i" + mats + "# " + matQty + " #t" + mats + "#"
         }
         if (cost > 0) {
            prompt += "\r\n#i4031138# " + cost + " meso"
         }
         cm.sendYesNo(prompt)
      } else if (status == 3) {
         boolean complete = true

         if (!cm.canHold(item, 1)) {
            cm.sendOk("2040021_NEED_FREE_SLOT")
            cm.dispose()
            return
         } else if (cm.getMeso() < cost) {
            cm.sendOk("2040021_ONLY_ACCEPT_MESO")
            cm.dispose()
            return
         } else {
            if (mats instanceof ArrayList && matQty instanceof ArrayList) {
               for (int i = 0; complete && i < mats.size(); i++) {
                  if (!cm.haveItem(mats[i] as Integer, matQty[i] as Integer)) {
                     complete = false
                  }
               }
            } else if (!cm.haveItem(mats as Integer, matQty as Integer)) {
               complete = false
            }
         }
         if (stimulator) { //check for stimulator
            if (!cm.haveItem(stimulatorId)) {
               complete = false
            }
         }
         if (!complete) {
            cm.sendOk("2040021_MISSING_ITEMS")
         } else {
            if (mats instanceof ArrayList && matQty instanceof ArrayList) {
               for (int i = 0; i < mats.size(); i++) {
                  cm.gainItem(mats[i] as Integer, (short) (-matQty[i] as Integer))
               }
            } else {
               cm.gainItem(mats as Integer, (short) (-matQty as Integer))
            }
            cm.gainMeso(-cost)
            if (stimulator) { //check for stimulator
               cm.gainItem(stimulatorId, (short) -1)
               int deleted = Math.floor(Math.random() * 10).intValue()
               if (deleted != 0) {
                  cm.gainItem(item, (short) 1, true, true)
                  cm.sendOk("2040021_SHOE_SUCCESS")
               } else {
                  cm.sendOk("2040021_STIMULATOR_ERROR")
               }
            } else {
               cm.gainItem(item, (short) 1)
               cm.sendOk("2040021_SHOE_SUCCESS")
            }
         }
         cm.dispose()
      }
   }
}

NPC2040021 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2040021(cm: cm))
   }
   return (NPC2040021) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
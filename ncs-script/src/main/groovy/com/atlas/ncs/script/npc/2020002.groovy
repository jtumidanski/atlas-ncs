package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2020002 {
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

   def start() {
      cm.getPlayer().setCS(true)
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == 1) {
         status++
      } else {
         cm.dispose()
      }
      if (status == 0 && mode == 1) {
         String selStr = "Hello there. El Nath winters are incredibly cold, you're going to need a warm pair of shoes to survive.#b"
         String[] options = ["Create Warrior shoes", "Create Bowman shoes", "Create Magician shoes", "Create Thief shoes"]
         for (int i = 0; i < options.length; i++) {
            selStr += "\r\n#L" + i + "# " + options[i] + "#l"
         }

         cm.sendSimple(selStr)
      } else if (status == 1 && mode == 1) {
         selectedType = selection
         String selStr = ""
         String[] shoes = []
         if (selectedType == 0) { //warrior shoes
            selStr = "Warrior shoes? Okay, then which set?#b"
            shoes = ["Sapphire Camel Boots#k - Warrior Lv. 60#b", "Orihalcon Camel Boots#k - Warrior Lv. 60#b", "Blood Camel Boots#k - Warrior Lv. 60#b",
                     "Blue Carzen Boots#k - Warrior Lv. 70#b", "Purple Carzen Boots#k - Warrior Lv. 70#b", "Dark Carzen Boots#k - Warrior Lv. 70#b",
                     "Red Rivers Boots#k - Warrior Lv. 80#b", "Blue Rivers Boots#k - Warrior Lv. 80#b", "Dark Rivers Boots#k - Warrior Lv. 80#b"]
         } else if (selectedType == 1) { //bowman shoes
            selStr = "Bowman shoes? Okay, then which set?#b"
            shoes = ["Red Gore Boots#k - Bowman Lv. 60#b", "Blue Gore Boots#k - Bowman Lv. 60#b", "Green Gore Boots#k - Bowman Lv. 60#b",
                     "Blue Elf Boots#k - Bowman Lv. 70#b", "Beige Elf Boots#k - Bowman Lv. 70#b", "Green Elf Boots#k - Bowman Lv. 70#b", "Dark Elf Boots#k - Bowman Lv. 70#b",
                     "Blue Wing Boots#k - Bowman Lv. 80#b", "Red Wing Boots#k - Bowman Lv. 80#b", "Green Wing Boots#k - Bowman Lv. 80#b", "Dark Wing Boots#k - Bowman Lv. 80#b"]
         } else if (selectedType == 2) { //magician shoes
            selStr = "Magician shoes? Okay, then which set?#b"
            shoes = ["Pink Goldrunners#k - Magician Lv. 60#b", "Green Goldrunners#k - Magician Lv. 60#b", "Orange Goldrunners#k - Magician Lv. 60#b", "Blue Goldrunners#k - Magician Lv. 60#b",
                     "Blue Lapiz Sandals#k - Magician Lv. 70#b", "Red Lapiz Sandals#k - Magician Lv. 70#b", "Brown Lapiz Sandals#k - Magician Lv. 70#b", "Gold Lapiz Sandals#k - Magician Lv. 70#b",
                     "Green Enigma Shoes#k - Magician Lv. 80#b", "Purple Enigma Shoes#k - Magician Lv. 80#b", "Dark Enigma Shoes#k - Magician Lv. 80#b"]
         } else if (selectedType == 3) { //thief shoes
            selStr = "Thief shoes? Okay, then which set?#b"
            shoes = ["Blood Moss Boots#k - Thief Lv. 60#b", "Gold Moss Boots#k - Thief Lv. 60#b", "Dark Moss Boots#k - Thief Lv. 60#b",
                     "Purple Mystique Shoes#k - Thief Lv. 70#b", "Blue Mystique Shoes#k - Thief Lv. 70#b", "Red Mystique Shoes#k - Thief Lv. 70#b",
                     "Green Pirate Boots#k - Thief Lv. 80#b", "Red Pirate Boots#k - Thief Lv. 80#b", "Dark Pirate Boots#k - Thief Lv. 80#b"]
         }
         for (int i = 0; i < shoes.length; i++) {
            selStr += "\r\n#L" + i + "# " + shoes[i] + "#l"
         }
         cm.sendSimple(selStr)
      } else if (status == 2 && mode == 1) {
         selectedItem = selection

         if (selectedType == 0) { //warrior shoes
            int[] itemSet = [1072147, 1072148, 1072149, 1072154, 1072155, 1072156, 1072210, 1072211, 1072212]
            List matSet = [[4021008, 4011007, 4021005, 4000030, 4003000], [4021008, 4011007, 4011005, 4000030, 4003000], [4021008, 4011007, 4021000, 4000030, 4003000],
                           [4005000, 4005002, 4011002, 4000048, 4003000], [4005000, 4005002, 4011005, 4000048, 4003000], [4005000, 4005002, 4021008, 4000048, 4003000],
                           [4005000, 4005002, 4021000, 4000030, 4003000], [4005000, 4005002, 4021002, 4000030, 4003000], [4005000, 4005002, 4021008, 4000030, 4003000]]
            List matQtySet = [[1, 1, 8, 80, 55], [1, 1, 8, 80, 55], [1, 1, 8, 80, 55], [1, 3, 5, 100, 55], [2, 2, 5, 100, 55], [3, 1, 1, 100, 55],
                              [2, 3, 7, 90, 65], [3, 2, 7, 90, 65], [4, 1, 2, 90, 65]]
            int[] costSet = [60000, 60000, 60000, 70000, 70000, 70000, 80000, 80000, 80000]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
         } else if (selectedType == 1) { //bowman shoes
            int[] itemSet = [1072144, 1072145, 1072146, 1072164, 1072165, 1072166, 1072167, 1072182, 1072183, 1072184, 1072185]
            List matSet = [[4011006, 4021000, 4021007, 4000030, 4003000], [4011006, 4021005, 4021007, 4000030, 4003000], [4011006, 4021003, 4021007, 4000030, 4003000],
                           [4005002, 4005000, 4021005, 4000055, 4003000], [4005002, 4005000, 4021004, 4000055, 4003000], [4005002, 4005000, 4021003, 4000055, 4003000], [4005002, 4005000, 4021008, 4000055, 4003000],
                           [4005002, 4005000, 4021002, 4000030, 4003000], [4005002, 4005000, 4021000, 4000030, 4003000], [4005002, 4005000, 4021003, 4000030, 4003000], [4005002, 4021008, 4000030, 4003000]]
            List matQtySet = [[5, 8, 1, 75, 50], [5, 8, 1, 75, 50], [5, 8, 1, 75, 50], [1, 3, 5, 100, 55], [2, 2, 5, 100, 55], [2, 2, 5, 100, 55], [3, 1, 1, 100, 55],
                              [2, 3, 7, 90, 60], [3, 2, 7, 90, 60], [4, 1, 7, 90, 60], [5, 2, 90, 60]]
            int[] costSet = [60000, 60000, 60000, 70000, 70000, 70000, 70000, 80000, 80000, 80000, 80000]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
         } else if (selectedType == 2) { //magician shoes
            int[] itemSet = [1072136, 1072137, 1072138, 1072139, 1072157, 1072158, 1072159, 1072160, 1072177, 1072178, 1072179]
            List matSet = [[4021009, 4011006, 4011005, 4000030, 4003000], [4021009, 4011006, 4021003, 4000030, 4003000], [4021009, 4011006, 4011003, 4000030, 4003000], [4021009, 4011006, 4021002, 4000030, 4003000],
                           [4005001, 4005003, 4021002, 4000051, 4003000], [4005001, 4005003, 4021000, 4000051, 4003000], [4005001, 4005003, 4011003, 4000051, 4003000], [4005001, 4005003, 4011006, 4000051, 4003000],
                           [4005001, 4005003, 4021003, 4000030, 4003000], [4005001, 4005003, 4021001, 4000030, 4003000], [4005001, 4005003, 4021008, 4000030, 4003000]]
            List matQtySet = [[1, 4, 5, 70, 50], [1, 4, 5, 70, 50], [1, 4, 5, 70, 50], [1, 4, 5, 70, 50],
                              [1, 3, 5, 100, 55], [2, 2, 5, 100, 55], [2, 2, 5, 100, 55], [3, 1, 3, 100, 55],
                              [2, 3, 7, 85, 60], [3, 2, 7, 85, 60], [4, 1, 2, 85, 60]]
            int[] costSet = [60000, 60000, 60000, 60000, 70000, 70000, 70000, 70000, 80000, 80000, 80000]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
         } else if (selectedType == 3) { //thief shoes
            int[] itemSet = [1072150, 1072151, 1072152, 1072161, 1072162, 1072163, 1072172, 1072173, 1072174]
            List matSet = [[4021007, 4011007, 4021000, 4000030, 4003000], [4021007, 4011007, 4011006, 4000030, 4003000], [4021007, 4011007, 4021008, 4000030, 4003000],
                           [4005003, 4005000, 4021001, 4000051, 4003000], [4005003, 4005002, 4021005, 4000051, 4003000], [4005002, 4005003, 4021000, 4000051, 4003000],
                           [4005000, 4005003, 4021003, 4000030, 4003000], [4005002, 4005003, 4021000, 4000030, 4003000], [4005003, 4005002, 4021008, 4000030, 4003000]]
            List matQtySet = [[1, 1, 8, 75, 50], [1, 1, 5, 75, 50], [1, 1, 1, 75, 50],
                              [1, 3, 5, 100, 55], [1, 3, 5, 100, 55], [1, 3, 5, 100, 55],
                              [3, 2, 7, 90, 60], [3, 2, 7, 90, 60], [3, 2, 7, 90, 60]]
            int[] costSet = [60000, 60000, 60000, 70000, 70000, 70000, 80000, 80000, 80000]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
         }
         String prompt = "You want me to make a #t" + item + "#? In that case, I'm going to need specific items from you in order to make it. Make sure you have room in your inventory, though!#b"
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
      } else if (status == 3 && mode == 1) {
         boolean complete = true

         if (!cm.canHold(item, 1)) {
            cm.sendOk("2020002_CHECK_INVENTORY_FOR_FREE_SLOT")
            cm.dispose()
            return
         } else if (cm.getMeso() < cost) {
            cm.sendOk("2020002_CANNOT_AFFORD_IT")
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
         if (!complete) {
            cm.sendOk("2020002_MISSING_MATERIALS")
         } else {
            if (mats instanceof ArrayList && matQty instanceof ArrayList) {
               for (int i = 0; i < mats.size(); i++) {
                  cm.gainItem(mats[i] as Integer, (short) (-matQty[i] as Integer))
               }
            } else {
               cm.gainItem(mats as Integer, (short) (-matQty as Integer))
            }
            cm.gainMeso(-cost)
            cm.gainItem(item, (short) 1)
            cm.sendOk("2020002_ALL_DONE")
         }
         cm.dispose()
      }
   }
}

NPC2020002 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2020002(cm: cm))
   }
   return (NPC2020002) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
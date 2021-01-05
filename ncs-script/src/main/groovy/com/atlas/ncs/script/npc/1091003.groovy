package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1091003 {
   NPCConversationManager cm
   int status = -1
   int sel = -1
   int selectedType
   int selectedItem
   int qty
   boolean equip
   int item
   Object mats
   Object matQty
   int cost
   int levelLimit
   String prompt

   def start() {
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
         String selStr = "What? You want to make your own weapons and gloves? Seriously... it's tough to do it by yourself if you don't have experience... I'll help you out. I've been a pirate for 20 years, and for 20 years I have made various items for the crew here. It's easy for me."
         String[] options = ["Make a Knuckler", "Make a Gun", "Make a pair of gloves"]
         for (int i = 0; i < options.length; i++) {
            selStr += "\r\n#b#L" + i + "# " + options[i] + "#l#k"
         }

         cm.sendSimple(selStr)
      } else if (status == 1 && mode == 1) {
         selectedType = selection
         if (selectedType == 0) { //Making a Knuckler
            String selStr = "As long as you bring in the materials required, I'll make you a fine Knuckler. Which Knuckler would you like to make?"
            String[] knucklers = ["Leather Arms (Level limit: 15, Pirate)", "Double Tail Knuckler (Level limit: 20, Pirate)", "Norman Grip (Level limit: 25, Pirate)", "Prime Hands (Level limit: 30, Pirate)", "Silver Maiden (Level limit: 35, Pirate)", "Neozard (Level limit: 40, Pirate)", "Fury Claw (Level limit: 50, Pirate)"]
            for (int i = 0; i < knucklers.length; i++) {
               selStr += "\r\n#b#L" + i + "# " + knucklers[i] + "#l#k"
            }
            equip = true
            cm.sendSimple(selStr)
         } else if (selectedType == 1) { //Making a Gun
            String selStr = "As long as you bring in the materials required, I'll make you a fine Gun. Which Gun would you like to make?"
            String[] guns = ["Dellinger Special (Level limit: 15, Pirate)", "The Negotiator (Level limit: 20, Pirate)", "Golden Hook (Level limit: 25, Pirate)", "Cold Mind (Level limit: 30, Pirate)", "Shooting Star (Level limit: 35, Pirate)", "Lunar Shooter (Level limit: 40, Pirate)", "Mr. Rasfelt (Level limit: 50, Pirate)"]
            for (int i = 0; i < guns.length; i++) {
               selStr += "\r\n#b#L" + i + "# " + guns[i] + "#l#k"
            }
            equip = true
            cm.sendSimple(selStr)
         } else if (selectedType == 2) { //Making a pair of pirate gloves
            String selStr = "As long as you bring in the materials required, I'll make you a fine glove. Which glove would you like to make?"
            String[] gloves = ["Green Lagger Halfglove", "Brown Leather Armour Glove", "Hard Leather Glove", "Yellow Tartis", "Brown Jewelled", "Brown Barbee", "Brown Royce", "Black Schult"]
            for (int i = 0; i < gloves.length; i++) {
               selStr += "\r\n#b#L" + i + "# " + gloves[i] + "#l#k"
            }
            equip = true
            cm.sendSimple(selStr)
         }
         if (equip) {
            status++
         }
      } else if (status == 3 && mode == 1) {
         if (equip) {
            selectedItem = selection
            qty = 1
         } else {
            qty = (selection > 0) ? selection : (selection < 0 ? -selection : 1)
         }

         if (selectedType == 0) { //Making a Knuckler
            int[] itemSet = [1482001, 1482002, 1482003, 1482004, 1482005, 1482006, 1482007]
            List matSet = [4000021, [4011001, 4011000, 4000021, 4003000], [4011000, 4011001, 4003000], [4011000, 4011001, 4000021, 4003000], [4011000, 4011001, 4000021, 4003000], [4011000, 4011001, 4021000, 4000021, 4003000], [4000039, 4011000, 4011001, 4000030, 4000021, 4003000]]
            List matQtySet = [20, [1, 1, 10, 5], [2, 1, 10], [1, 1, 30, 10], [2, 2, 30, 20], [1, 1, 2, 50, 20], [150, 1, 2, 20, 20, 20]]
            int[] costSet = [1000, 2000, 5000, 15000, 30000, 50000, 100000]
            int[] levelLimitSet = [15, 20, 25, 30, 35, 40, 50]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
            levelLimit = levelLimitSet[selectedItem]
         } else if (selectedType == 1) { //Making a Gun
            int[] itemSet = [1492001, 1492002, 1492003, 1492004, 1492005, 1492006, 1492007]
            List matSet = [[4011000, 4003000, 4003001], [4011000, 4003000, 4003001, 4000021], [4011000, 4003000], [4011001, 4000021, 4003000], [4011006, 4011001, 4000021, 4003000], [4011004, 4011001, 4000021, 4003000], [4011006, 4011004, 4011001, 4000030, 4003000]]
            List matQtySet = [[1, 5, 1], [1, 10, 5, 10], [2, 10], [2, 10, 10], [10, 2, 5, 10], [1, 2, 10, 20], [1, 2, 4, 30, 30]]
            int[] costSet = [1000, 2000, 5000, 15000, 30000, 50000, 100000]
            int[] levelLimitSet = [15, 20, 25, 30, 35, 40, 50]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
            levelLimit = levelLimitSet[selectedItem]
         } else if (selectedType == 2) { //Making a pair of pirate gloves
            int[] itemSet = [1082180, 1082183, 1082186, 1082189, 1082192, 1082195, 1082198, 1082201]
            List matSet = [[4000021, 4021003], 4000021, [4011000, 4000021], [4021006, 4000021, 4003000], [4011000, 4000021, 4003000], [4000021, 4011000, 4011001, 4003000], [4011000, 4000021, 4000030, 4003000], [4011007, 4021008, 4021007, 4000030, 4003000]]
            List matQtySet = [[15, 1], 35, [2, 20], [2, 50, 10], [3, 60, 15], [80, 3, 3, 25], [3, 20, 40, 30], [1, 1, 1, 50, 50]]
            int[] costSet = [1000, 8000, 15000, 25000, 30000, 40000, 50000, 70000]
            int[] levelLimitSet = [15, 20, 25, 30, 35, 40, 50, 60]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
            levelLimit = levelLimitSet[selectedItem]
         }

         prompt = "Making one #t" + item + "# requires the items listed below. The level limit for this item is " + levelLimit + ", so check and make sure you really need this item before getting it. What do you think? Do you really want one?\r\n"

         if (mats instanceof ArrayList && matQty instanceof ArrayList) {
            for (int i = 0; i < mats.size(); i++) {
               prompt += "\r\n#i" + mats[i] + "# " + ((matQty[i] as Integer) * qty) + " #t" + mats[i] + "#"
            }
         } else {
            prompt += "\r\n#i" + mats + "# " + ((matQty as Integer) * qty) + " #t" + mats + "#"
         }

         if (cost > 0) {
            prompt += "\r\n#i4031138# " + cost * qty + " meso"
         }

         cm.sendYesNo(prompt)
      } else if (status == 4 && mode == 1) {
         boolean pass = true

         if (!cm.canHold(item)) {
            cm.sendOk("1091003_CHECK_YOUR_INVENTORY")
            cm.dispose()
            return
         } else if (cm.getMeso() < cost * qty) {
            cm.sendNext("1091003_MISSING_ITEMS")
            cm.dispose()
            return
         } else {
            if (mats instanceof ArrayList && matQty instanceof ArrayList) {
               for (int i = 0; pass && i < mats.size(); i++) {
                  if (!cm.haveItem(mats[i] as Integer, ((matQty[i] as Integer) * qty))) {
                     pass = false
                  }
               }
            } else if (!cm.haveItem(mats as Integer, ((matQty as Integer) * qty))) {
               pass = false
            }
            /*if (mats instanceof Array) {
               for(var i = 0; pass && i < mats.length; i++)
               {
                  if (matQty[i] * qty == 1)	{
                     if (!cm.haveItem(mats[i]))
                     {
                        pass = false;
                     }
                  }
                  else {
                     var count = 0;
                     var iter = cm.getChar().getInventory(MapleInventoryType.ETC).listById(mats[i]).iterator();
                     while (iter.hasNext()) {
                        count += iter.next().getQuantity();
                     }
                     if (count < matQty[i] * qty)
                        pass = false;
                  }
               }
            }
            else {
               var count = 0;
               var iter = cm.getChar().getInventory(MapleInventoryType.ETC).listById(mats).iterator();
               while (iter.hasNext()) {
                  count += iter.next().getQuantity();
               }
               if (count < matQty * qty)
                  pass = false;
            }
                                */
         }

         if (!pass) {
            cm.sendNext("1091003_MISSING_ITEMS")
         } else {
            if (mats instanceof ArrayList && matQty instanceof ArrayList) {
               for (int i = 0; i < mats.size(); i++) {
                  cm.gainItem(mats[i] as Integer, (short) ((-matQty[i] as Integer) * qty))
               }
            } else {
               cm.gainItem(mats as Integer, (short) ((-matQty as Integer) * qty))
            }

            if (cost > 0) {
               cm.gainMeso(-cost * qty)
            }

            if (item == 4003000)//screws
            {
               cm.gainItem(4003000, (short) (15 * qty))
            } else {
               cm.gainItem(item, (short) qty)
            }
            cm.sendOk("1091003_ALL_DONE")
         }
         cm.dispose()
      }
   }
}

NPC1091003 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1091003(cm: cm))
   }
   return (NPC1091003) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
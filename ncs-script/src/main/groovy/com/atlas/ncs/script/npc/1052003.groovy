package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1052003 {
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
   boolean last_use

   def start() {
      cm.getPlayer().setCS(true)
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0 && type > 0) {    // hope types 2 & 3 works as well, as 1 and 4 END CHAT
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }

         if (status == 0) {
            String selStr = "Yes, I do own this forge. If you're willing to pay, I can offer you some of my services.#b"
            String[] options = ["Refine a mineral ore", "Refine a jewel ore", "I have Iron Hog's Metal Hoof...", "Upgrade a claw"]
            for (int i = 0; i < options.length; i++) {
               selStr += "\r\n#L" + i + "# " + options[i] + "#l"
            }

            cm.sendSimple(selStr)
         } else if (status == 1) {
            selectedType = selection
            if (selectedType == 0) { //mineral refine
               String selStr = "So, what kind of mineral ore would you like to refine?#b"
               String[] minerals = ["Bronze", "Steel", "Mithril", "Adamantium", "Silver", "Orihalcon", "Gold"]
               for (int i = 0; i < minerals.length; i++) {
                  selStr += "\r\n#L" + i + "# " + minerals[i] + "#l"
               }
               equip = false
               cm.sendSimple(selStr)
            } else if (selectedType == 1) { //jewel refine
               String selStr = "So, what kind of jewel ore would you like to refine?#b"
               String[] jewels = ["Garnet", "Amethyst", "Aquamarine", "Emerald", "Opal", "Sapphire", "Topaz", "Diamond", "Black Crystal"]
               for (int i = 0; i < jewels.length; i++) {
                  selStr += "\r\n#L" + i + "# " + jewels[i] + "#l"
               }
               equip = false
               cm.sendSimple(selStr)
            } else if (selectedType == 2) { //foot refine
               String selStr = "You know about that? Not many people realize the potential in the Iron Hog's Metal Hoof... I can make this into something special, if you want me to."
               equip = false
               cm.sendYesNo(selStr)
            } else if (selectedType == 3) { //claw refine
               String selStr = "Ah, you wish to upgrade a claw? Then tell me, which one?#b"
               String[] claws = ["Blood Gigantic#k - Thief Lv. 60#b", "Sapphire Gigantic#k - Thief Lv. 60#b", "Dark Gigantic#k - Thief Lv. 60#b"]
               for (int i = 0; i < claws.length; i++) {
                  selStr += "\r\n#L" + i + "# " + claws[i] + "#l"
               }
               equip = true
               cm.sendSimple(selStr)
            }
            if (equip) {
               status++
            }
         } else if (status == 2 && mode == 1) {
            selectedItem = selection
            if (selectedType == 0) { //mineral refine
               int[] itemSet = [4011000, 4011001, 4011002, 4011003, 4011004, 4011005, 4011006]
               List matSet = [4010000, 4010001, 4010002, 4010003, 4010004, 4010005, 4010006]
               List matQtySet = [10, 10, 10, 10, 10, 10, 10]
               int[] costSet = [300, 300, 300, 500, 500, 500, 800]
               item = itemSet[selectedItem]
               mats = matSet[selectedItem]
               matQty = matQtySet[selectedItem]
               cost = costSet[selectedItem]
            } else if (selectedType == 1) { //jewel refine
               int[] itemSet = [4021000, 4021001, 4021002, 4021003, 4021004, 4021005, 4021006, 4021007, 4021008]
               List matSet = [4020000, 4020001, 4020002, 4020003, 4020004, 4020005, 4020006, 4020007, 4020008]
               List matQtySet = [10, 10, 10, 10, 10, 10, 10, 10, 10]
               int[] costSet = [500, 500, 500, 500, 500, 500, 500, 1000, 3000]
               item = itemSet[selectedItem]
               mats = matSet[selectedItem]
               matQty = matQtySet[selectedItem]
               cost = costSet[selectedItem]
            } else if (selectedType == 2) { //special refine
               int[] itemSet = [4011001, 1]
               List matSet = [4000039, 1]
               List matQtySet = [100, 1]
               int[] costSet = [1000, 1]
               item = itemSet[0]
               mats = matSet[0]
               matQty = matQtySet[0]
               cost = costSet[0]
            }

            String prompt = "So, you want me to make some #t" + item + "#s? In that case, how many do you want me to make?"

            cm.sendGetNumber(prompt, 1, 1, 100)
         } else if (status == 3) {
            if (equip) {
               selectedItem = selection
               qty = 1
            } else {
               qty = (selection > 0) ? selection : (selection < 0 ? -selection : 1)
            }

            last_use = false

            if (selectedType == 3) { //claw refine
               int[] itemSet = [1472023, 1472024, 1472025]
               List matSet = [[1472022, 4011007, 4021000, 2012000], [1472022, 4011007, 4021005, 2012002], [1472022, 4011007, 4021008, 4000046]]
               List matQtySet = [[1, 1, 8, 10], [1, 1, 8, 10], [1, 1, 3, 5]]
               int[] costSet = [80000, 80000, 100000]
               item = itemSet[selectedItem]
               mats = matSet[selectedItem]
               matQty = matQtySet[selectedItem]
               cost = costSet[selectedItem]
               if (selectedItem != 2) {
                  last_use = true
               }
            }

            String prompt = "You want me to make "
            if (qty == 1) {
               prompt += "a #t" + item + "#?"
            } else {
               prompt += qty + " #t" + item + "#?"
            }

            prompt += " In that case, I'm going to need specific items from you in order to make it. Make sure you have room in your inventory, though!#b"

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
         } else if (status == 4) {
            boolean complete = true

            if (!cm.canHold(item, qty)) {
               cm.sendOk("1052003_CHECK_YOUR_INVENTORY_FOR_FREE_SLOT")
               cm.dispose()
               return
            } else if (cm.getMeso() < cost * qty) {
               cm.sendOk("1052003_NOT_ENOUGH_MESOS")
               cm.dispose()
               return
            } else {
               if (mats instanceof ArrayList && matQty instanceof ArrayList) {
                  for (int i = 0; complete && i < mats.size(); i++) {
                     if (!cm.haveItem(mats[i] as Integer, ((matQty[i] as Integer) * qty))) {
                        complete = false
                     }
                  }
               } else if (!cm.haveItem(mats as Integer, ((matQty as Integer) * qty))) {
                  complete = false
               }
            }

            if (!complete) {
               cm.sendOk("1052003_MATERIALS_MISSING")
            } else {
               if (mats instanceof ArrayList && matQty instanceof ArrayList) {
                  for (int i = 0; i < mats.size(); i++) {
                     cm.gainItem(mats[i] as Integer, (short) ((-matQty[i] as Integer) * qty))
                  }
               } else {
                  cm.gainItem(mats as Integer, (short) ((-matQty as Integer) * qty))
               }
               cm.gainMeso(-cost * qty)
               cm.gainItem(item, (short) qty)
               cm.sendNext("1052003_SUCCESS")
            }
            cm.dispose()
         }
      }
   }
}

NPC1052003 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1052003(cm: cm))
   }
   return (NPC1052003) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
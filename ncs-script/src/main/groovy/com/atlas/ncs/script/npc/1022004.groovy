package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1022004 {
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
         String selStr = cm.evaluateToken("1022004_HELLO")
         String[] options = ["1022004_MAKE_A_GLOVE", "1022004_UPGRADE_A_GLOVE", "1022004_CREATE_MATERIALS"]
         for (int i = 0; i < options.length; i++) {
            selStr += "\r\n#L" + i + "# " + cm.evaluateToken(options[i]) + "#l"
         }

         cm.sendSimple(selStr)
      } else if (status == 1 && mode == 1) {
         selectedType = selection
         if (selectedType == 0) { //glove refine
            String selStr = cm.evaluateToken("1022004_WHICH_GLOVE")
            String[] items = ["JUNO", "STEEL_FINGERLESS_GLOVES", "VENON", "WHITE_FINGERLESS_GLOVES", "BRONZE_MISSEL", "STEEL_BRIGGON", "IRON_KNUCKLE", "STEEL_BRIST", "BRONZE_CLENCH"]
            for (int i = 0; i < items.length; i++) {
               selStr += "\r\n#L" + i + "# " + cm.evaluateToken(items[i]) + "#l"
            }
            cm.sendSimple(selStr)
            equip = true
         } else if (selectedType == 1) { //glove upgrade
            String selStr = cm.evaluateToken("1022004_WHICH_GLOVE_TO_UPGRADE")
            String[] crystals = ["STEEL_MISSEL", "ORIHALCON_MISSEL", "YELLOW_BRIGGON", "DARK_BRIGGON", "ADAMANTIUM_KNUCKL", "DARK_KNUCKLE", "MITHRIL_BRIST", "GOLD_BRIST", "SAPPHIRE_CLENCH", "DARK_CLENCH"]
            for (int i = 0; i < crystals.length; i++) {
               selStr += "\r\n#L" + i + "# " + cm.evaluateToken(crystals[i]) + "#l"
            }
            cm.sendSimple(selStr)
            equip = true
         } else if (selectedType == 2) { //material refine
            String selStr = cm.evaluateToken("1022004_WHAT_MATERIALS")
            String[] materials = ["PROCESSED_WOOD_FROM_TREE_BRANCH", "PROCESSED_WOOD_FROM_FIREWOOD", "SCREWS"]
            for (int i = 0; i < materials.length; i++) {
               selStr += "\r\n#L" + i + "# " + cm.evaluateToken(materials[i]) + "#l"
            }
            cm.sendSimple(selStr)
            equip = false
         }
         if (equip) {
            status++
         }
      } else if (status == 2 && mode == 1) {
         selectedItem = selection
         if (selectedType == 2) { //material refine
            int[] itemSet = [4003001, 4003001, 4003000]
            List matSet = [4000003, 4000018, [4011000, 4011001]]
            List matQtySet = [10, 5, [1, 1]]
            int[] costSet = [0, 0, 0]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
         }

         cm.sendGetNumber("1022004_HOW_MANY", 1, 1, 100, item)
      } else if (status == 3 && mode == 1) {
         if (equip) {
            selectedItem = selection
            qty = 1
         } else {
            qty = (selection > 0) ? selection : (selection < 0 ? -selection : 1)
         }

         if (selectedType == 0) { //glove refine
            int[] itemSet = [1082003, 1082000, 1082004, 1082001, 1082007, 1082008, 1082023, 1082009, 1082059]
            List matSet = [[4000021, 4011001], 4011001, [4000021, 4011000], 4011001, [4011000, 4011001, 4003000], [4000021, 4011001, 4003000], [4000021, 4011001, 4003000],
                           [4011001, 4021007, 4000030, 4003000], [4011007, 4011000, 4011006, 4000030, 4003000]]
            List matQtySet = [[15, 1], 2, [40, 2], 2, [3, 2, 15], [30, 4, 15], [50, 5, 40], [3, 2, 30, 45], [1, 8, 2, 50, 50]]
            int[] costSet = [1000, 2000, 5000, 10000, 20000, 30000, 40000, 50000, 70000]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
         } else if (selectedType == 1) { //glove upgrade
            int[] itemSet = [1082005, 1082006, 1082035, 1082036, 1082024, 1082025, 1082010, 1082011, 1082060, 1082061]
            List matSet = [[1082007, 4011001], [1082007, 4011005], [1082008, 4021006], [1082008, 4021008], [1082023, 4011003], [1082023, 4021008],
                           [1082009, 4011002], [1082009, 4011006], [1082059, 4011002, 4021005], [1082059, 4021007, 4021008]]
            List matQtySet = [[1, 1], [1, 2], [1, 3], [1, 1], [1, 4], [1, 2], [1, 5], [1, 4], [1, 3, 5], [1, 2, 2]]
            int[] costSet = [20000, 25000, 30000, 40000, 45000, 50000, 55000, 60000, 70000, 80000]
            item = itemSet[selectedItem]
            mats = matSet[selectedItem]
            matQty = matQtySet[selectedItem]
            cost = costSet[selectedItem]
         }

         String itemPrompt = ""
         if (qty == 1) {
            itemPrompt += "a #t" + item + "#?"
         } else {
            itemPrompt += qty + " #t" + item + "#?"
         }

         String materialList = ""
         if (mats instanceof ArrayList && matQty instanceof ArrayList) {
            for (int i = 0; i < mats.size(); i++) {
               materialList += "\r\n#i" + mats[i] + "# " + ((matQty[i] as Integer) * qty) + " #t" + mats[i] + "#"
            }
         } else {
            materialList += "\r\n#i" + mats + "# " + ((matQty as Integer) * qty) + " #t" + mats + "#"
         }

         if (cost > 0) {
            materialList += "\r\n#i4031138# " + cost * qty + " meso"
         }

         cm.sendYesNo("1022004_CONFIRM", itemPrompt, materialList)
      } else if (status == 4 && mode == 1) {
         boolean complete = true
         int recvItem = item, recvQty

         if (item == 4003000)//screws
         {
            recvQty = 15 * qty
         } else {
            recvQty = qty
         }

         if (!cm.canHold(recvItem, recvQty)) {
            cm.sendOk("1022004_NEED_FREE_SLOT")
            cm.dispose()
            return
         } else if (cm.getMeso() < cost * qty) {
            cm.sendOk("1022004_CANNOT_AFFORD")
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
            cm.sendOk("1022004_MISSING_SOMETHING")
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

            cm.gainItem(recvItem, (short) recvQty)
            cm.sendOk("1022004_FINISHED")
         }
         cm.dispose()
      }
   }
}

NPC1022004 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1022004(cm: cm))
   }
   return (NPC1022004) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
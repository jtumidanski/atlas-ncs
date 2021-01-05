package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9000041 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   String[] options = ["EQUIP", "USE", "SET-UP", "ETC"]
   String name
   int selectedType = 0

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      status++
      if (mode != 1) {
         cm.dispose()
         return
      }

      if (status == 0) {
         if (!cm.getConfiguration().enableCustomNpcScript()) {
            cm.sendOk("9000041_MEDAL_RANKING_UNAVAILABLE")
            cm.dispose()
            return
         }

         String selStr = "Hello, I am the #bBazaar NPC#k! Sell to me any item on your inventory you don't need. #rWARNING#b: Make sure you have your items ready to sell at the slots #rAFTER#b the item you have selected to sell.#k Any items #bunder#k the item selected will be sold thoroughly."
         for (int i = 0; i < options.length; i++) {
            selStr += "\r\n#L" + i + "# " + options[i] + "#l"
         }
         cm.sendSimple(selStr)
      } else if (status == 1) {
         selectedType = selection
         cm.sendGetText("9000041_START", options[selectedType])

      } else if (status == 2) {
         name = cm.getText()
         int res = cm.getPlayer().sellAllItemsFromName((byte) (selectedType + 1), name)

         if (res > -1) {
            cm.sendOk("9000041_COMPLETE", cm.numberWithCommas(res))
         } else {
            cm.sendOk("9000041_NOT_IN_INVENTORY", name, options[selectedType])
         }

         cm.dispose()
      }
   }
}

NPC9000041 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9000041(cm: cm))
   }
   return (NPC9000041) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
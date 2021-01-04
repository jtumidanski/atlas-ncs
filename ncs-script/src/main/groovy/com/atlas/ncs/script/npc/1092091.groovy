package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1092091 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.getQuestProgressInt(2180, 1) == 2) {
         cm.sendNext("1092091_CHECK_ANOTHER_COW")
         cm.dispose()
         return
      }

      if (cm.canHold(4031848) && cm.haveItem(4031847)) {
         cm.sendNext("1092091_ONE_THIRD_FULL")
         cm.gainItem(4031847, (short) -1)
         cm.gainItem(4031848, (short) 1)

         cm.setQuestProgress(2180, 1, 2)
      } else if (cm.canHold(4031849) && cm.haveItem(4031848)) {
         cm.sendNext("1092091_TWO_THIRDS_FULL")
         cm.gainItem(4031848, (short) -1)
         cm.gainItem(4031849, (short) 1)

         cm.setQuestProgress(2180, 1, 2)
      } else if (cm.canHold(4031850) && cm.haveItem(4031849)) {
         cm.sendNext("1092091_FULL")
         cm.gainItem(4031849, (short) -1)
         cm.gainItem(4031850, (short) 1)

         cm.setQuestProgress(2180, 1, 2)
      } else {
         cm.sendNext("1092091_INVENTORY_FULL")
      }
      cm.dispose()
   }

   def action(Byte mode, Byte type, Integer selection) {

   }
}

NPC1092091 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1092091(cm: cm))
   }
   return (NPC1092091) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
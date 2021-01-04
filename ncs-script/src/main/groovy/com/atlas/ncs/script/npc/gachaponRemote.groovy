package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NpcGachaponRemote {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int ticketId = 5451000
   String[] mapName = ["Henesys", "Ellinia", "Perion", "Kerning City", "Sleepywood", "Mushroom Shrine", "Showa Spa (M)", "Showa Spa (F)", "New Leaf City", "Nautilus"]
   String curMapName = ""

   def start() {
      status = -1
      curMapName = mapName[(cm.getNpcId() != 9100117 && cm.getNpcId() != 9100109) ? (cm.getNpcId() - 9100100) : cm.getNpcId() == 9100109 ? 8 : 9]
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode < 0) {
         cm.dispose()
      } else {
         if (mode == 1) {
            status++
         } else {
            status--
         }
         if (status == 0 && cm.haveItem(ticketId)) {
            if (cm.canHold(1302000) && cm.canHold(2000000) && cm.canHold(3010001) && cm.canHold(4000000)) {
               // One free slot in every inventory.
               cm.gainItem(ticketId, (short) -1)
               cm.doGachapon()
            } else {
               cm.sendOk("gachaponRemote_NEED_INVENTORY_SPACE_FREE")

            }
         } else {
            cm.dispose()
         }
      }
   }
}

NpcGachaponRemote getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NpcGachaponRemote(cm: cm))
   }
   return (NpcGachaponRemote) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
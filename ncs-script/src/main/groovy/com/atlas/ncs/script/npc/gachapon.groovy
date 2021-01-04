package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NpcGachapon {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int ticketId = 5220000
   String[] mapName = ["Henesys", "Ellinia", "Perion", "Kerning City", "Sleepywood", "Mushroom Shrine", "Showa Spa (M)", "Showa Spa (F)", "Ludibrium", "New Leaf City", "El Nath", "Nautilus"]
   String curMapName = ""

   def start() {
      status = -1
      curMapName = mapName[(cm.getNpcId() != 9100117 && cm.getNpcId() != 9100109) ? (cm.getNpcId() - 9100100) : cm.getNpcId() == 9100109 ? 9 : 11]
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
         if (status == 0 && mode == 1) {
            if (cm.haveItem(ticketId)) {
               cm.sendYesNo("gachapon_USE_YOUR_TICKET", curMapName)
            } else {
               cm.sendSimple("gachapon_WELCOME", curMapName)
            }
         } else if (status == 1 && cm.haveItem(ticketId)) {
            if (cm.canHold(1302000) && cm.canHold(2000000) && cm.canHold(3010001) && cm.canHold(4000000)) {
               // One free slot in every inventory.
               cm.gainItem(ticketId, (short) -1)
               cm.doGachapon()
            } else {
               cm.sendOk("gachapon_NEED_INVENTORY_SPACE_FREE")
            }
            cm.dispose()
         } else if (status == 1) {
            if (selection == 0) {
               cm.sendNext("gachapon_GACHAPON_DETAIL")
            } else {
               cm.sendNext("gachapon_TICKET_DETAIL")
            }
         } else if (status == 2) {
            cm.sendNextPrev("gachapon_VARIETY_OF_ITEMS", curMapName, curMapName)
         } else {
            cm.dispose()
         }
      }
   }
}

NpcGachapon getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NpcGachapon(cm: cm))
   }
   return (NpcGachapon) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
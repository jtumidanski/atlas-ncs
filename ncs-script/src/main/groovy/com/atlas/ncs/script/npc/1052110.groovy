package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

/*
	NPC Name: 		Trash Can 2
	Map(s): 		
	Description: 	Kerning Subway
*/
class NPC1052110 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0 && type > 0) {
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }

         if (status == 0) {
            cm.sendOk("1052110_JUST_A_TRASH_CAN")
            cm.dispose()
         }
      }
   }
}

NPC1052110 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1052110(cm: cm))
   }
   return (NPC1052110) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
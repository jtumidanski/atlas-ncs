package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

/*
	NPC Name: 		Fourth Eos Rock
	Map(s): 		Ludibrium : Eos Tower 1st Floor
	Description: 	
*/
class NPC2040027 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.haveItem(4001020)) {
         cm.sendYesNo("2040027_CHOICES")
      } else {
         cm.sendOk("2040027_NEED_SCROLL")
         cm.dispose()
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (!(mode < 1)) {
         cm.gainItem(4001020, (short) -1)
         cm.warp(221021700, 3)
      }
      cm.dispose()
   }
}

NPC2040027 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2040027(cm: cm))
   }
   return (NPC2040027) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
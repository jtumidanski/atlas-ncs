package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

/*
	NPC Name: 		First Eos Rock
	Map(s): 		Ludibrium : Eos Tower 100th Floor
	Description: 	
*/
class NPC2040024 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.haveItem(4001020)) {
         cm.sendYesNo("2040024_TO_71")
      } else {
         cm.sendOk("2040024_NEED_SCROLL")
         cm.dispose()
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (!(mode < 1)) {
         cm.gainItem(4001020, (short) -1)
         cm.warp(221022900, 3)
      }
      cm.dispose()
   }
}

NPC2040024 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2040024(cm: cm))
   }
   return (NPC2040024) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
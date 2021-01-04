package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

/*
	NPC Name: 		Shane
	Map(s): 		
	Description: 	
*/
class NPC1032004 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      cm.sendYesNo("1032004_RETURN_TO_ELLINIA")
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode > 0) {
         cm.warp(101000000, 0)
      }
      cm.dispose()
   }
}

NPC1032004 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1032004(cm: cm))
   }
   return (NPC1032004) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
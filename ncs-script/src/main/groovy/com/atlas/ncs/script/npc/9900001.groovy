package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9900001 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.gmLevel() > 1) {
         cm.sendYesNo("9900001_LEVEL_UP")

      } else {
         cm.sendOk("9900001_WASSUP")

      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode > 0 && cm.gmLevel() > 1) {
         cm.levelUp(true)
      }
      cm.dispose()
   }
}

NPC9900001 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9900001(cm: cm))
   }
   return (NPC9900001) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
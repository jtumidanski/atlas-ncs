package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9201104 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.getMapId() == 610020000) {
         cm.sendOk("9201104_BRAVE_ADVENTURER")

      } else if (cm.getMapId() == 610020003) {
         cm.sendOk("9201104_YOU_SEEM_WORTHY")

      } else if (cm.getMapId() == 610020004) {
         cm.sendOk("9201104_YOU_SEEM_WORTHY_LONG")

      } else {
         cm.sendOk("9201104_PROGRESS_IS_SPLENDID")

      }
      cm.dispose()
   }

   def action(Byte mode, Byte type, Integer selection) {

   }
}

NPC9201104 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201104(cm: cm))
   }
   return (NPC9201104) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
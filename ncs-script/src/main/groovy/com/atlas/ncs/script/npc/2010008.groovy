package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2010008 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      cm.sendSimple("2010008_WHAT_WOULD_YOU_LIKE_TO_DO")
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode < 1) {
         cm.dispose()
      } else {
         status++
         if (status == 1) {
            sel = selection
            if (sel == 0) {
               if (cm.getPlayer().getGuildRank() == 1) {
                  cm.sendYesNo("2010008_EMBLEM_CHANGE_COST")
               } else {
                  cm.sendOk("2010008_MUST_BE_LEADER_TO_CHANGE_EMBLEM")
               }
            }
         } else if (status == 2 && sel == 0) {
            cm.getPlayer().genericGuildMessage(17)
            cm.dispose()
         } else {
            cm.dispose()
         }
      }
   }
}

NPC2010008 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2010008(cm: cm))
   }
   return (NPC2010008) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9040010 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      EventInstanceManager eim = cm.getEventInstance()
      if (eim != null) {
         if (cm.isEventLeader()) {
            if (cm.haveItem(4001024)) {
               cm.removeAll(4001024)
               Object prev = eim.setProperty("bossclear", "true", true)
               if (prev == null) {
                  int start = (eim.getProperty("entryTimestamp")).toInteger()
                  long diff = System.currentTimeMillis() - start

                  int points = 1000 - Math.floor(diff / (100 * 60)).intValue()
                  if (points < 100) {
                     points = 100
                  }

                  cm.gainGP(cm.getGuildId(), points)
               }

               eim.clearPQ()
            } else {
               cm.sendOk("9040010_FINAL_CHALLENGE")
            }
         } else {
            cm.sendOk("9040010_FINAL_CHALLENGE_LEADER")
         }
      } else {
         cm.warp(990001100)
      }

      cm.dispose()
   }

   def action(Byte mode, Byte type, Integer selection) {

   }
}

NPC9040010 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9040010(cm: cm))
   }
   return (NPC9040010) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
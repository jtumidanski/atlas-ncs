package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9040003 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   static def clearStage(int stage, EventInstanceManager eim) {
      eim.setProperty("stage" + stage + "clear", "true")
      eim.showClearEffect(true)
      eim.giveEventPlayersStageReward(stage)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 1) {
            status++
         } else {
            cm.dispose()
         }

         EventInstanceManager eim = cm.getEventInstance()
         if (eim.getProperty("stage4clear") != null && eim.getProperty("stage4clear") == "true") {
            cm.sendOk("9040003_IMMORTAL_SLEEP")

            cm.dispose()
            return
         }

         if (status == 0) {
            if (cm.isEventLeader()) {
               cm.sendNext("9040003_IMMORTAL_SLEEP_WILL_NOW")
               clearStage(4, eim)
               cm.gainGP(cm.getGuildId(), 30)
               cm.forceHitReactor("ghostgate", (byte) 1)
               cm.dispose()
            } else {
               cm.sendOk("9040003_LEADER_MUST_SPEAK")
               cm.dispose()
            }
         }
      }
   }
}

NPC9040003 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9040003(cm: cm))
   }
   return (NPC9040003) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
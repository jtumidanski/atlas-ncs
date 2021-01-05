package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC2040036 {
   NPCConversationManager cm
   int status = 0
   int sel = -1
   int curMap, stage

   def start() {
      curMap = cm.getMapId()
      stage = Math.floor((curMap - 922010100) / 100).intValue() + 1
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   static def clearStage(int stage, EventInstanceManager eim, int curMap) {
      eim.setProperty(stage + "stageclear", "true")
      eim.showClearEffect(true)
      eim.linkToNextStage(stage, "lpq", curMap)  //opens the portal to the next map
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else if (mode == 0) {
         cm.dispose()
      } else {
         if (mode == 1) {
            status++
         } else {
            status--
         }

         EventInstanceManager eim = cm.getEventInstance()

         if (eim.getProperty(stage.toString() + "stageclear") != null) {
            cm.sendNext("2040036_GO_TO_NEXT_STAGE")
         } else {
            if (eim.isEventLeader(cm.getCharacterId())) {
               int state = eim.getIntProperty("statusStg" + stage)

               if (state == -1) {           // preamble
                  cm.sendOk("2040036_WELCOME", stage)
                  eim.setProperty("statusStg" + stage, 0)
               } else {       // check stage completion
                  if (cm.haveItem(4001022, 25)) {
                     cm.sendOk("2040036_GOOD_JOB")
                     cm.gainItem(4001022, (short) -25)

                     eim.setProperty("statusStg" + stage, 1)
                     clearStage(stage, eim, curMap)
                  } else {
                     cm.sendNext("2040036_SORRY")
                  }
               }
            } else {
               cm.sendNext("2040036_PARTY_LEADER_MUST_TALK")
            }
         }

         cm.dispose()
      }
   }
}

NPC2040036 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2040036(cm: cm))
   }
   return (NPC2040036) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
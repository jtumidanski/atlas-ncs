package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9040009 {
   NPCConversationManager cm
   int status = -1
   int sel = -1
   int stage

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
         if (mode == 0 && status == 0) {
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }

         EventInstanceManager eim = cm.getEventInstance()
         if (eim == null) {
            cm.warp(990001100)
         } else {
            if (eim.getProperty("stage1clear") == "true") {
               cm.sendOk("9040009_EXCELLENT_WORK")
               cm.dispose()
               return
            }

            if (cm.isEventLeader()) {
               if (status == 0) {
                  if (eim.getProperty("stage1status") == null || eim.getProperty("stage1status") == "waiting") {
                     if (eim.getProperty("stage1phase") == null) {
                        stage = 1
                        eim.setProperty("stage1phase", stage)
                     } else {
                        stage = (eim.getProperty("stage1phase")).toInteger()
                     }

                     if (stage == 1) {
                        cm.sendOk("9040009_IN_THIS_CHALLENGE")
                     } else {
                        cm.sendOk("9040009_GOOD_LUCK")
                     }
                  } else if (eim.getProperty("stage1status") == "active") {
                     stage = (eim.getProperty("stage1phase")).toInteger()

                     if (eim.getProperty("stage1combo") == eim.getProperty("stage1guess")) {
                        if (stage == 3) {
                           cm.forceHitReactor("statuegate", (byte) 1)
                           clearStage(1, eim)
                           cm.gainGP(cm.getGuildId(), 15)
                           cm.sendOk("9040009_EXCELLENT_WORK")
                        } else {
                           cm.sendOk("9040009_VERY_GOOD")
                           eim.setProperty("stage1phase", stage + 1)
                           cm.sendPinkTextToMap("GATE_KEEPER_TEST_PART_COMPLETE", stage)
                        }

                     } else {
                        eim.showWrongEffect()
                        cm.sendOk("9040009_FAILED")
                        cm.sendPinkTextToMap("GATE_KEEPER_TEST_FAILED")
                        eim.setProperty("stage1phase", "1")
                     }
                     eim.setProperty("stage1status", "waiting")
                     cm.dispose()
                  } else {
                     cm.sendOk("9040009_PLEASE_WAIT")
                     cm.dispose()
                  }
               } else if (status == 1) {
                  int[] reactors = getReactors()
                  int[] combo = makeCombo(reactors)
                  cm.sendPinkTextToMap("GATE_KEEPER_TEST_COMBINATION_REVEALED")
                  int delay = 5000
                  for (int i = 0; i < combo.length; i++) {
                     cm.delayedHitReactor(combo[i], delay + 3500 * i)
                  }
                  eim.setProperty("stage1status", "display")
                  eim.setProperty("stage1combo", "")
                  cm.dispose()
               }
            } else {
               cm.sendOk("9040009_LEADER_MUST_SPEAK")
               cm.dispose()
            }
         }
      }
   }

   def getReactors() {
      int[] reactors = []

      Iterator<MapleReactor> iter = cm.getPlayer().getMap().getReactors().iterator() as Iterator<MapleReactor>
      while (iter.hasNext()) {
         MapleReactor mo = iter.next()
         if (mo.getName() != "statuegate") {
            reactors << mo.objectId()
         }
      }

      return reactors
   }

   def makeCombo(int[] reactors) {
      int[] combo = []
      while (combo.length < (stage + 3)) {
         int chosenReactor = reactors[Math.floor(Math.random() * reactors.length).intValue()]
         boolean repeat = false
         if (combo.length > 0) {
            for (int i = 0; i < combo.length; i++) {
               if (combo[i] == chosenReactor) {
                  repeat = true
                  break
               }
            }
         }
         if (!repeat) {
            combo << chosenReactor
         }
      }
      return combo
   }
}

NPC9040009 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9040009(cm: cm))
   }
   return (NPC9040009) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
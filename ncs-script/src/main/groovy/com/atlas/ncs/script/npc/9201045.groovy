package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.NPCConversationManager

import java.awt.Rectangle

class NPC9201045 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   boolean debug = false
   int curMap, stage

   def isAllGatesOpen() {
      MapleMap map = cm.getPlayer().getMap()

      for (int i = 0; i < 7; i++) {
         MapleReactor gate = map.getReactorByName("gate0" + i)
         if (gate.getState() != ((byte) 4)) {
            return false
         }
      }

      return true
   }

   static def clearStage(int stage, EventInstanceManager eim, int curMap) {
      eim.setProperty(stage + "stageclear", "true")
      eim.showClearEffect(true)
      eim.linkToNextStage(stage, "apq", curMap)  //opens the portal to the next map
   }

   def start() {
      curMap = cm.getMapId()
      stage = Math.floor((curMap - 670010200) / 100).intValue() + 1
      status = -1
      action((byte) 1, (byte) 0, 0)
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
            if (stage < 5) {
               cm.sendNext("9201045_ALREADY_OPEN")
            } else if (stage == 5) {
               eim.warpEventTeamToMapSpawnPoint(670010700, 0)
            } else {
               if (cm.isEventLeader()) {
                  if (eim.getIntProperty("marriedGroup") == 0) {
                     eim.restartEventTimer(1 * 60 * 1000)
                     eim.warpEventTeam(670010800)
                  } else {
                     eim.setIntProperty("marriedGroup", 0)
                     eim.restartEventTimer(2 * 60 * 1000)
                     eim.warpEventTeamToMapSpawnPoint(670010750, 1)
                  }
               } else {
                  cm.sendNext("9201045_LEADERS_COMMAND")
               }
            }
         } else {
            if (stage != 6) {
               if (eim.isEventLeader(cm.getCharacterId())) {
                  int state = eim.getIntProperty("statusStg" + stage)

                  if (state == -1) {           // preamble
                     if (stage == 4) {
                        cm.sendOk("9201045_WELCOME", stage)
                     } else if (stage == 5) {
                        cm.sendOk("9201045_WELCOME_2", stage)
                     }

                     int st = (debug) ? 2 : 0
                     eim.setProperty("statusStg" + stage, st)
                  } else {       // check stage completion
                     if (stage == 4) {
                        if (cm.haveItem(4031597, 50)) {
                           cm.gainItem(4031597, (short) -50)

                           long tl = eim.getTimeLeft()
                           if (tl >= 5 * 60 * 1000) {
                              eim.setProperty("timeLeft", tl.toString())
                              eim.restartEventTimer(4 * 60 * 1000)
                           }

                           cm.sendNext("9201045_WELL_DONE")
                           cm.sendPinkTextToMap("9201045_TIME_RUNS_SHORT")
                           clearStage(stage, eim, curMap)
                        } else {
                           cm.sendNext("9201045_DID_YOU_NOT_PAY_HEED")
                        }

                     } else if (stage == 5) {
                        boolean pass = true

                        if (eim.isEventTeamTogether()) {
                           List<Integer> party = cm.getEventInstance().getCharacterIds()
                           Rectangle area = cm.getMapArea(2)

                           for (int i = 0; i < party.size(); i++) {
                              int characterId = party[i]
                              if (cm.characterIsAlive(characterId) && !area.contains(cm.characterPosition(characterId))) {
                                 pass = false
                                 break
                              }
                           }
                        } else {
                           pass = false
                        }

                        if (pass) {
                           if (isAllGatesOpen()) {
                              String tl = eim.getProperty("timeLeft")
                              if (tl != null) {
                                 long tr = eim.getTimeLeft()

                                 Float tlf = (tl).toFloat()
                                 eim.restartEventTimer((long) tlf - (4 * 60 * 1000 - tr))
                              }

                              cm.sendNext("9201045_ALREADY_GATHERED")
                              cm.sendPinkTextToMap("9201045_BOSS_FIGHT")
                              clearStage(stage, eim, curMap)
                           } else {
                              cm.sendNext("9201045_I_CAN_TELL_IT")
                           }
                        } else {
                           cm.sendNext("9201045_TEAM_NOT_GATHERED")
                        }
                     }
                  }
               } else {
                  cm.sendNext("9201045_PARTY_LEADER_MUST_TALK")
               }
            } else {
               Rectangle area = cm.getMapArea(0)
               if (area.contains(cm.characterPosition(cm.getCharacterId()))) {
                  if (cm.characterIsAlive(cm.getCharacterId())) {
                     cm.warp(670010700, "st01")
                  } else {
                     cm.sendNext("9201045_STAND_BACK")

                  }
               } else {
                  if (cm.isEventLeader()) {
                     if (cm.haveItem(4031594, 1)) {
                        cm.gainItem(4031594, (short) -1)
                        cm.sendNext("9201045_CONGRATULATIONS")
                        clearStage(stage, eim, curMap)
                        eim.clearPQ()
                     } else {
                        cm.sendNext("9201045_HOW_IS_IT")
                     }
                  } else {
                     cm.sendNext("9201045_PARTY_LEADER_MUST_TALK")
                  }
               }
            }
         }

         cm.dispose()
      }
   }
}

NPC9201045 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201045(cm: cm))
   }
   return (NPC9201045) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
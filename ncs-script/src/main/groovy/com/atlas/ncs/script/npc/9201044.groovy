package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9201044 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   boolean debug = false
   boolean autoPass = false

   def spawnMobs(maxSpawn) {
      int[] spawnPosX
      int[] spawnPosY

      MapleMap mapObj = cm.getMap()
      if (stage == 2) {
         spawnPosX = [619, 299, 47, -140, -471]
         spawnPosY = [-840, -840, -840, -840, -840]

         for (int i = 0; i < 5; i++) {
            for (int j = 0; j < 2; j++) {
               MapleLifeFactory.getMonster(9400515).ifPresent({ mobObj1 -> mapObj.spawnMonsterOnGroundBelow(mobObj1, new Point(spawnPosX[i], spawnPosY[i])) })
               MapleLifeFactory.getMonster(9400516).ifPresent({ mobObj2 -> mapObj.spawnMonsterOnGroundBelow(mobObj2, new Point(spawnPosX[i], spawnPosY[i])) })
               MapleLifeFactory.getMonster(9400517).ifPresent({ mobObj3 -> mapObj.spawnMonsterOnGroundBelow(mobObj3, new Point(spawnPosX[i], spawnPosY[i])) })
            }
         }
      } else {
         spawnPosX = [2303, 1832, 1656, 1379, 1171]
         spawnPosY = [240, 150, 300, 150, 240]

         for (int i = 0; i < maxSpawn; i++) {
            int rndMob = 9400519 + Math.floor(Math.random() * 4).intValue()
            int rndPos = Math.floor(Math.random() * 5).intValue()

            MapleLifeFactory.getMonster(rndMob).ifPresent({ mobObj -> mapObj.spawnMonsterOnGroundBelow(mobObj, new Point(spawnPosX[rndPos], spawnPosY[rndPos])) })
         }
      }
   }

   static def generateCombo1() {
      int[] positions = [0, 0, 0, 0, 0, 0, 0, 0, 0]
      int rndPicked = Math.floor(Math.random() * Math.pow(3, 5)).intValue()

      while (rndPicked > 0) {
         (positions[rndPicked % 3])++

         rndPicked = Math.floor(rndPicked / 3).intValue()
      }

      String returnString = ""
      for (int i = 0; i < positions.length; i++) {
         returnString += positions[i]
         if (i != positions.length - 1) {
            returnString += ","
         }
      }

      return returnString
   }

   static def generateCombo2() {
      int toPick = 5, rndPicked
      int[] positions = [0, 0, 0, 0, 0, 0, 0, 0, 0]
      while (toPick > 0) {
         rndPicked = Math.floor(Math.random() * 9).intValue()

         if (positions[rndPicked] == 0) {
            positions[rndPicked] = 1
            toPick--
         }
      }

      String returnString = ""
      for (int i = 0; i < positions.length; i++) {
         returnString += positions[i]
         if (i != positions.length - 1) {
            returnString += ","
         }
      }

      return returnString
   }

   int curMap, stage

   def clearStage(stage, eim, curMap) {
      eim.setProperty(stage + "stageclear", "true")
      if (stage > 1) {
         eim.showClearEffect(true)
         eim.linkToNextStage(stage, "apq", curMap)  //opens the portal to the next map
      } else {
         cm.getMap().getPortal("go01").setPortalState(false)

         int val = Math.floor(Math.random() * 3).intValue()
         eim.showClearEffect(670010200, "gate" + val, 2)

         cm.getMap().getPortal("go0" + val).setPortalState(true)
         eim.linkPortalToScript(stage, "go0" + val, "apq0" + val, curMap)
      }
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

         EventInstanceManager eim = cm.getPlayer().getEventInstance()

         if (eim.getProperty(stage.toString() + "stageclear") != null) {
            cm.sendNext("9201044_PORTAL_IS_ALREADY_OPOEN")

         } else {
            if (eim.isEventLeader(cm.getPlayer())) {
               int state = eim.getIntProperty("statusStg" + stage)

               if (state == -1) {           // preamble
                  if (stage == 1) {
                     cm.sendOk("9201044_WELCOME", stage)

                  } else if (stage == 2) {
                     cm.sendOk("9201044_WELCOME_CLIMB", stage)

                  } else if (stage == 3) {
                     cm.sendOk("9201044_WELCOME_CLIMB_2", stage)

                  }

                  int st = (autoPass) ? 2 : 0
                  eim.setProperty("statusStg" + stage, st)
               } else {       // check stage completion
                  if (state == 2) {
                     eim.setProperty("statusStg" + stage, 1)
                     clearStage(stage, eim, curMap)
                     cm.dispose()
                     return
                  }

                  MapleMap map = cm.getPlayer().getMap()
                  if (stage == 1) {
                     if (eim.getIntProperty("statusStg" + stage) == 1) {
                        clearStage(stage, eim, curMap)
                     } else {
                        cm.sendOk("9201044_TALK_WITH")

                     }
                  } else if (stage == 2 || stage == 3) {
                     if (map.countMonsters() == 0) {
                        int[] objectSet = [0, 0, 0, 0, 0, 0, 0, 0, 0]
                        int playersOnCombo = 0
                        MapleCharacter[] party = cm.getEventInstance().getPlayers()
                        for (int i = 0; i < party.size(); i++) {
                           for (int y = 0; y < map.getAreas().size(); y++) {
                              if (map.getArea(y).contains(party[i].position())) {
                                 playersOnCombo++
                                 objectSet[y] += 1
                                 break
                              }
                           }
                        }

                        if (playersOnCombo == 5/* || cm.getPlayer().gmLevel() > 1*/ || debug) {
                           String comboStr = eim.getProperty("stage" + stage + "combo")
                           if (comboStr == null || comboStr == "") {
                              if (stage == 2) {
                                 comboStr = generateCombo1()
                              } else {
                                 comboStr = generateCombo2()
                              }

                              eim.setProperty("stage" + stage + "combo", comboStr)
                              if (debug) {
                                 print("generated " + comboStr + " for stg" + stage + "\n")
                              }
                           }

                           String[] combo = comboStr.split(',')
                           boolean correctCombo = true
                           int guessedRight = objectSet.length
                           int playersRight = 0

                           if (!debug) {
                              for (int i = 0; i < objectSet.length; i++) {
                                 if ((combo[i]).toInteger() != objectSet[i]) {
                                    correctCombo = false
                                    guessedRight--
                                 } else {
                                    if (objectSet[i] > 0) {
                                       playersRight++
                                    }
                                 }
                              }
                           } else {
                              for (int i = 0; i < objectSet.length; i++) {
                                 int ci = cm.getPlayer().countItem(4000000 + i)

                                 if (ci != (combo[i]).toInteger()) {
                                    correctCombo = false
                                    guessedRight--
                                 } else {
                                    if (ci > 0) {
                                       playersRight++
                                    }
                                 }
                              }
                           }


                           if (correctCombo/* || cm.getPlayer().gmLevel() > 1*/) {
                              eim.setProperty("statusStg" + stage, 1)
                              clearStage(stage, eim, curMap)
                              cm.dispose()
                           } else {
                              int miss = eim.getIntProperty("missCount") + 1
                              int maxMiss = (stage == 2) ? 7 : 1

                              if (miss < maxMiss) {   //already implies stage 2
                                 eim.setIntProperty("missCount", miss)

                                 if (guessedRight == 6) { //6 unused slots on this stage
                                    cm.sendNext("9201044_WEIGH_DIFFERENTLY")

                                    MessageBroadcaster.getInstance().sendMapServerNotice(cm.getPlayer().getMap(), ServerNoticeType.PINK_TEXT, I18nMessage.from("9201044_ALL_WEIGH_DIFFERENTLY"))
                                 } else {
                                    cm.sendNext("9201044_THINK_YOUR_NEXT_COURSE")

                                    MessageBroadcaster.getInstance().sendMapServerNotice(cm.getPlayer().getMap(), ServerNoticeType.PINK_TEXT, I18nMessage.from("9201044_ONE_WEIGH_SAME"))
                                 }
                              } else {
                                 spawnMobs(playersRight)
                                 eim.setIntProperty("missCount", 0)
                                 if (stage == 2) {
                                    eim.setProperty("stage2combo", "")

                                    cm.sendNext("9201044_FAILED_TO_DISCOVER")

                                    MessageBroadcaster.getInstance().sendMapServerNotice(cm.getPlayer().getMap(), ServerNoticeType.PINK_TEXT, I18nMessage.from("9201044_WRONG_COMBINATION"))
                                 }
                              }

                              eim.showWrongEffect()
                              cm.dispose()
                           }
                        } else {
                           if (stage == 2) {
                              cm.sendNext("9201044_HOW_TO")

                           } else {
                              cm.sendNext("9201044_HOW_TO_2")

                           }

                           cm.dispose()
                        }
                     } else {
                        cm.sendNext("9201044_DEFEAT_ALL_MOBS")

                     }
                  }
               }
            } else {
               cm.sendNext("9201044_PARTY_LEADER_MUST_TALK")

            }
         }

         cm.dispose()
      }
   }
}

NPC9201044 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201044(cm: cm))
   }
   return (NPC9201044) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
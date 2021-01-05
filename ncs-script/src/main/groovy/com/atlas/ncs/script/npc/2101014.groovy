package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.Expedition
import com.atlas.ncs.processor.ExpeditionType
import com.atlas.ncs.processor.NPCConversationManager

class NPC2101014 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int arenaType
   int map
   ExpeditionType expeditionType = ExpeditionType.ARIANT
   ExpeditionType expeditionType1 = ExpeditionType.ARIANT1
   ExpeditionType expeditionType2 = ExpeditionType.ARIANT2

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
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
         if (cm.getMapId() == 980010000) {
            if (cm.getLevel() > 30) {
               cm.sendOk("2101014_MAXIMUM_LEVEL_30")

               cm.dispose()
               return
            }

            if (status == 0) {
               Expedition expedition = cm.getExpedition(expeditionType)
               Expedition expedition1 = cm.getExpedition(expeditionType1)
               Expedition expedition2 = cm.getExpedition(expeditionType2)
               String startSnd = "What would you like to do? \r\n\r\n\t#e#r(Choose a Battle Arena)#n#k\r\n#b"
               String toSnd = startSnd

               if (expedition == null) {
                  toSnd += "#L0#Battle Arena (1) (Empty)#l\r\n"
               } else if (cm.hasCharactersInMap(980010101)) {
                  toSnd += "#L0#Join Battle Arena (1)  Owner (" + expedition.getLeaderName() + ")" + " Current Member: " + cm.getExpeditionMemberNames(expeditionType) + "\r\n"
               }
               if (expedition1 == null) {
                  toSnd += "#L1#Battle Arena (2) (Empty)#l\r\n"
               } else if (cm.hasCharactersInMap(980010201)) {
                  toSnd += "#L1#Join Battle Arena (2)  Owner (" + expedition1.getLeaderName() + ")" + " Current Member: " + cm.getExpeditionMemberNames(expeditionType1) + "\r\n"
               }
               if (expedition2 == null) {
                  toSnd += "#L2#Battle Arena (3) (Empty)#l\r\n"
               } else if (cm.hasCharactersInMap(980010301)) {
                  toSnd += "#L2#Join Battle Arena (3)  Owner (" + expedition2.getLeaderName() + ")" + " Current Member: " + cm.getExpeditionMemberNames(expeditionType2) + "\r\n"
               }
               if (toSnd == startSnd) {
                  cm.sendOk("2101014_ARENA_CURRENTLY_OCCUPIED")

                  cm.dispose()
               } else {
                  cm.sendSimple(toSnd)
               }
            } else if (status == 1) {
               arenaType = selection
               Expedition expedition = fetchArenaType()
               if (expedition == null) {
                  cm.dispose()
                  return
               }

               if (expedition != null) {
                  enterArena(-1)
               } else {
                  cm.sendGetText("2101014_HOW_MANY_CAN_JOIN")
               }
            } else if (status == 2) {
               Integer players = (cm.getText()).toInteger()
               if (players == null) {
                  cm.sendNext("2101014_ENTER_NUMERIC_LIMIT")
                  status = 0
               } else if (players < 2) {
                  cm.sendNext("2101014_MINIMUM_2_PLAYERS")
                  status = 0
               } else {
                  enterArena(players)
               }
            }
         }
      }
   }


   def fetchArenaType() {
      Expedition expedition
      switch (arenaType) {
         case 0:
            expeditionType = ExpeditionType.ARIANT
            expedition = cm.getExpedition(expeditionType)
            map = 980010100
            break
         case 1:
            expeditionType = ExpeditionType.ARIANT1
            expedition = cm.getExpedition(expeditionType)
            map = 980010200
            break
         case 2:
            expeditionType = ExpeditionType.ARIANT2
            expedition = cm.getExpedition(expeditionType)
            map = 980010300
            break
         default:
            expeditionType = null
            map = 0
            expedition = null
      }

      return expedition
   }

   def enterArena(int arenaPlayers) {
      Expedition expedition = fetchArenaType()
      if (expedition == null) {
         cm.dispose()
      } else if (expedition == null) {
         if (arenaPlayers != -1) {
            int res = cm.createExpedition(expeditionType, true, 0, arenaPlayers)
            if (res == 0) {
               cm.warp(map, 0)
               cm.sendNotice("ARENA_CREATED_WAIT_FOR_PEOPLE_TO_JOIN")
            } else if (res > 0) {
               cm.sendOk("2101014_QUOTA_LIMIT")
            } else {
               cm.sendOk("2101014_UNEXPECTED_ERROR")
            }
         } else {
            cm.sendOk("2101014_UNEXPECTED_ERROR")
         }

         cm.dispose()
      } else {
         if (playerAlreadyInLobby(cm.getCharacterId())) {
            cm.sendOk("2101014_YOU_ARE_ALREADY_INSIDE")
            cm.dispose()
            return
         }

         int playerAdd = expedition.addMemberInt(cm.getCharacterId())
         if (playerAdd == 3) {
            cm.sendOk("2101014_LOBBY_IS_FULL")
            cm.dispose()
         } else {
            if (playerAdd == 0) {
               cm.warp(map, 0)
               cm.dispose()
            } else if (playerAdd == 2) {
               cm.sendOk("2101014_LEADER_DID_NOT_ALLOW_YOU")
               cm.dispose()
            } else {
               cm.sendOk("2101014_ERROR")
               cm.dispose()
            }
         }
      }
   }

   def playerAlreadyInLobby(int characterId) {
      return cm.getExpedition(ExpeditionType.ARIANT) != null && cm.getExpedition(ExpeditionType.ARIANT).contains(characterId) ||
            cm.getExpedition(ExpeditionType.ARIANT1) != null && cm.getExpedition(ExpeditionType.ARIANT1).contains(characterId) ||
            cm.getExpedition(ExpeditionType.ARIANT2) != null && cm.getExpedition(ExpeditionType.ARIANT2).contains(characterId)
   }
}

NPC2101014 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2101014(cm: cm))
   }
   return (NPC2101014) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
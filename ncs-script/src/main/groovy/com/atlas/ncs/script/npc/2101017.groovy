package com.atlas.ncs.script.npc

import com.atlas.ncs.model.ExpeditionCharacter
import com.atlas.ncs.processor.Expedition
import com.atlas.ncs.processor.ExpeditionType
import com.atlas.ncs.processor.NPCConversationManager

class NPC2101017 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   ExpeditionType expeditionType
   Expedition expedition
   List<ExpeditionCharacter> expeditionMembers

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {

      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0) {
            cm.dispose()
            return
         }

         if (cm.getMapId() == 980010100 || cm.getMapId() == 980010200 || cm.getMapId() == 980010300) {
            if (cm.getMapId() == 980010100) {
               expeditionType = ExpeditionType.ARIANT
               expedition = cm.getExpedition(expeditionType)

            } else if (cm.getMapId() == 980010200) {
               expeditionType = ExpeditionType.ARIANT1
               expedition = cm.getExpedition(expeditionType)
            } else {
               expeditionType = ExpeditionType.ARIANT2
               expedition = cm.getExpedition(expeditionType)
            }

            if (expedition == null) {
               cm.dispose()
               return
            }

            expeditionMembers = expedition.getMemberList()
            if (status == 0) {
               if (cm.isLeaderExpedition(expeditionType)) {
                  cm.sendSimple("2101017_WHAT_WOULD_YOU_LIKE_TO_DO")

                  status = 1
               } else {
                  String toSend = "Current members inside this arena:\r\n#b"
                  toSend += cm.getExpeditionMemberNames(expeditionType)
                  cm.sendOk(toSend)
                  cm.dispose()
               }
            } else if (status == 1) {
               if (selection == 1) {
                  String toSend = "Current members inside this arena:\r\n#b"
                  toSend += cm.getExpeditionMemberNames(expeditionType)
                  cm.sendOk(toSend)
                  cm.dispose()
               } else if (selection == 2) {
                  int size = expeditionMembers.size()
                  if (size == 1) {
                     cm.sendOk("2101017_ONLY_MEMBER_OF_EXPEDITION")

                     cm.dispose()
                     return
                  }
                  String text = "The following members make up your expedition (Click on them to expel them):\r\n"
                  text += "\r\n\t\t1." + expedition.getLeaderName()
                  for (int i = 1; i < size; i++) {
                     text += "\r\n#b#L" + (i + 1) + "#" + (i + 1) + ". " + expeditionMembers.get(i).name() + "#l\n"
                  }
                  cm.sendSimple(text)
                  status = 6
               } else if (selection == 3) {
                  if (expedition.getMemberList().size() < 1) {
                     cm.sendOk("2101017_NEED_MORE_PLAYERS")
                     cm.dispose()
                  } else {
                     if (cm.getParty() != null) {
                        cm.sendOk("2101017_CANNOT_ENTER_AS_PARTY")
                        cm.dispose()
                        return
                     }

                     String errorMsg = cm.startAriantBattle(expeditionType, cm.getMapId())
                     if (errorMsg != "") {
                        cm.sendOk(errorMsg)
                     }

                     cm.dispose()
                  }
               } else if (selection == 4) {
                  cm.sendPinkTextToMap("ARENA_LEADER_LEFT")
                  expedition.warpTeam(980010000)
                  cm.endExpedition(expedition)
                  cm.dispose()
               }
            } else if (status == 6) {
               if (selection > 0) {
                  ExpeditionCharacter banned = expeditionMembers.get(selection - 1)
                  expedition.ban(banned)
                  cm.sendOk("2101017_YOU_HAVE_BANNED", banned.name())
                  cm.dispose()
               } else {
                  cm.sendSimple(list)
                  status = 2
               }
            }
         } else if (GameConstants.isAriantColiseumArena(cm.getMapId())) {
            if (cm.getMapId() == 980010101) {
               expeditionType = ExpeditionType.ARIANT
               expedition = cm.getExpedition(expeditionType)
            } else if (cm.getMapId() == 980010201) {
               expeditionType = ExpeditionType.ARIANT1
               expedition = cm.getExpedition(expeditionType)
            } else {
               expeditionType = ExpeditionType.ARIANT2
               expedition = cm.getExpedition(expeditionType)
            }
            if (status == 0) {
               String gotTheBombs = expedition.getProperty("gotBomb" + cm.getCharacterId())
               if (gotTheBombs != null) {
                  cm.sendOk("2101017_ALREADY_GAVE_BOMB")
                  cm.dispose()
               } else if (cm.canHoldAll([2270002, 2100067], [50, 5])) {
                  cm.sendOk("2101017_I_HAVE_GIVEN_YOU")
                  expedition.setProperty("gotBomb" + cm.getCharacterId(), "1")
                  cm.gainItem(2270002, (short) 50)
                  cm.gainItem(2100067, (short) 5)
                  cm.dispose()
               } else {
                  cm.sendOk("2101017_INVENTORY_IS_FULL")
                  cm.dispose()
               }
            }
         } else {
            cm.sendOk("2101017_HELLO")
            cm.dispose()
         }
      }
   }
}

NPC2101017 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2101017(cm: cm))
   }
   return (NPC2101017) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
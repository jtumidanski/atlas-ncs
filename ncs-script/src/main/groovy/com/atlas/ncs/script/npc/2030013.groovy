package com.atlas.ncs.script.npc

import com.atlas.ncs.model.ExpeditionCharacter
import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.Expedition
import com.atlas.ncs.processor.ExpeditionType
import com.atlas.ncs.processor.NPCConversationManager

class NPC2030013 {
   NPCConversationManager cm
   int status = 0
   int sel = -1
   Expedition expedition
   List<ExpeditionCharacter> expeditionMembers
   EventManager em
   ExpeditionType expeditionType = ExpeditionType.ZAKUM
   String expeditionName = "Zakum"
   String expeditionBoss = "Zakum"
   String expeditionMap = "Zakum's Altar"
   int expeditionItem = 4001017

   String list = "What would you like to do?#b\r\n\r\n#L1#View current Expedition members#l\r\n#L2#Start the fight!#l\r\n#L3#Stop the expedition.#l"

   def start() {
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      expedition = cm.getExpedition(expeditionType)
      em = cm.getEventManager("ZakumBattle").orElseThrow()

      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0) {
            cm.dispose()
            return
         }

         if (status == 0) {
            if (cm.getLevel() < expeditionType.getMinLevel() || cm.getLevel() > expeditionType.getMaxLevel()) {
               cm.sendOk("2030013_DO_NOT_MEET_CRITERIA", expeditionBoss)
               cm.dispose()
            } else if (expedition == null) { //Start an expedition
               cm.sendSimple("2030013_LIKE_TO_ASSEMBLE_A_TEAM", expeditionName, em.getProperty("party"), expeditionBoss)
               status = 1
            } else if (expedition.isLeader(cm.getCharacterId())) { //If you're the leader, manage the expedition
               if (expedition.isInProgress()) {
                  cm.sendOk("2030013_EXPEDITION_IN_PROGRESS_ALREADY")
                  cm.dispose()
               } else {
                  cm.sendSimple(list)
                  status = 2
               }
            } else if (expedition.isRegistering()) { //If the expedition is registering
               if (expedition.contains(cm.getCharacterId())) { //If you're in it but it hasn't started, be patient
                  cm.sendOk("2030013_ALREADY_REGISTERED", expedition.getLeaderName())
                  cm.dispose()
               } else { //If you aren't in it, you're going to get added
                  cm.sendOk(expedition.addMember(cm.getCharacterId()))
                  cm.dispose()
               }
            } else if (expedition.isInProgress()) { //Only if the expedition is in progress
               if (expedition.contains(cm.getCharacterId())) { //If you're registered, warp you in
                  EventInstanceManager eim = em.getInstance(expeditionName + cm.getChannelId())
                  if (eim.getIntProperty("canJoin") == 1) {
                     eim.registerPlayer(cm.getCharacterId())
                  } else {
                     cm.sendOk("2030013_ALREADY_STARTED", expeditionBoss)
                  }

                  cm.dispose()
               } else { //If you're not in by now, tough luck
                  cm.sendOk("2030013_ANOTHER_ALREADY_STARTED", expeditionBoss)
                  cm.dispose()
               }
            }
         } else if (status == 1) {
            if (selection == 1) {
               if (!cm.haveItem(expeditionItem)) {
                  cm.sendOk("2030013_INVENTORY_IS_FULL", expeditionItem, expeditionBoss)
                  cm.dispose()
                  return
               }

               expedition = cm.getExpedition(expeditionType)
               if (expedition != null) {
                  cm.sendOk("2030013_SOMEONE_ALREADY_LEADER")
                  cm.dispose()
                  return
               }

               int res = cm.createExpedition(expeditionType)
               if (res == 0) {
                  cm.sendOk("2030013_EXPEDITION_CREATED", expeditionBoss)
               } else if (res > 0) {
                  cm.sendOk("2030013_REACHED_QUOTA")
               } else {
                  cm.sendOk("2030013_UNEXPECTED_ERROR")
               }

               cm.dispose()
            } else if (selection == 2) {
               cm.sendOk("2030013_NOT_EVERYONE_IS_UP_TO_CHALLENGING", expeditionBoss)
               cm.dispose()
            }
         } else if (status == 2) {
            if (selection == 1) {
               if (expedition == null) {
                  cm.sendOk("2030013_COULD_NOT_BE_LOADED")
                  cm.dispose()
                  return
               }
               expeditionMembers = expedition.getMemberList()
               int size = expeditionMembers.size()
               if (size == 1) {
                  cm.sendOk("2030013_NEED_MORE_MEMBERS")
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
            } else if (selection == 2) {
               int min = expeditionType.getMinSize()

               int size = expedition.getMemberList().size()
               if (size < min) {
                  cm.sendOk("2030013_MINIMUM_MEMBERS", min)
                  cm.dispose()
                  return
               }

               cm.sendOk("2030013_EXPEDITION_WILL_BEGIN", expeditionMap)
               status = 4
            } else if (selection == 3) {
               cm.sendBlueTextToMap("EXPEDITION_ENDED_BY", expedition.getLeaderName())
               cm.endExpedition(expedition)
               cm.sendOk("2030013_EXPEDITION_ENDED")
               cm.dispose()
            }
         } else if (status == 4) {
            if (em == null) {
               cm.sendOk("2030013_COULD_NOT_BE_INITIALIZED")
               cm.dispose()
               return
            }

            em.setProperty("leader", cm.getCharacterName())
            em.setProperty("channel", cm.getChannelId())
            if (!em.startInstance(expedition)) {
               cm.sendOk("2030013_ANOTHER_ALREADY_STARTED", expeditionBoss)
               cm.dispose()
               return
            }

            cm.dispose()
         } else if (status == 6) {
            if (selection > 0) {
               ExpeditionCharacter banned = expeditionMembers.get(selection - 1)
               expedition.ban(banned)
               cm.sendOk("2030013_YOU_HAVE_BANNED", banned.name())
               cm.dispose()
            } else {
               cm.sendSimple(list)
               status = 2
            }
         }
      }
   }
}

NPC2030013 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2030013(cm: cm))
   }
   return (NPC2030013) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
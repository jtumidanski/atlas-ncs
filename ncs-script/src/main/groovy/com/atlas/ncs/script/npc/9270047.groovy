package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9270047 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   MapleExpedition expedition
   List<Map.Entry<Integer, String>> expeditionMembers
   MapleCharacter player
   EventManager em
   MapleExpeditionType expeditionType = MapleExpeditionType.SCARGA
   String expeditionName = "Scarlion"
   String expeditionBoss = "Scarlion and Targa"
   String expeditionMap = "Spooky World"
   int expeditionItem = 4032246

   String list = "What would you like to do?#b\r\n\r\n#L1#View current Expedition members#l\r\n#L2#Start the fight!#l\r\n#L3#Stop the expedition.#l"

   def start() {
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      player = cm.getPlayer()
      expedition = cm.getExpedition(expeditionType)
      em = cm.getEventManager("ScargaBattle")

      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0) {
            cm.dispose()
            return
         }

         if (status == 0) {
            if (player.getLevel() < expeditionType.getMinLevel() || player.getLevel() > expeditionType.getMaxLevel()) {
               cm.sendOk("9270047_DO_NOT_MEET_CRITERIA", expeditionBoss)

               cm.dispose()
            } else if (expedition == null) { //Start an expedition
               cm.sendSimple("9270047_EXPEDITION_INFO", expeditionName, em.getProperty("party"), expeditionBoss)

               status = 1
            } else if (expedition.isLeader(player)) { //If you're the leader, manage the expedition
               if (expedition.isInProgress()) {
                  cm.sendOk("9270047_ALREADY_IN_PROGRESS")

                  cm.dispose()
               } else {
                  cm.sendSimple(list)
                  status = 2
               }
            } else if (expedition.isRegistering()) { //If the expedition is registering
               if (expedition.contains(player)) { //If you're in it but it hasn't started, be patient
                  cm.sendOk("9270047_ALREADY_REGISTERED", expedition.getLeader().getName())
                  cm.dispose()
               } else { //If you aren't in it, you're going to get added
                  cm.sendOk(expedition.addMember(cm.getPlayer()))
                  cm.dispose()
               }
            } else if (expedition.isInProgress()) { //Only if the expedition is in progress
               if (expedition.contains(player)) { //If you're registered, warp you in
                  EventInstanceManager eim = em.getInstance(expeditionName + player.getClient().getChannel())
                  if (eim.getIntProperty("canJoin") == 1) {
                     eim.registerPlayer(player)
                  } else {
                     cm.sendOk("9270047_ALREADY_STARTED", expeditionBoss)
                  }

                  cm.dispose()
               } else { //If you're not in by now, tough luck
                  cm.sendOk("9270047_ANOTHER_EXPEDITION", expeditionBoss)
                  cm.dispose()
               }
            }
         } else if (status == 1) {
            if (selection == 1) {
               if (!cm.haveItem(expeditionItem)) {
                  cm.sendOk("9270047_MUST_HAVE_ITEM", expeditionItem, expeditionBoss)
                  cm.dispose()
                  return
               }

               expedition = cm.getExpedition(expeditionType)
               if (expedition != null) {
                  cm.sendOk("9270047_ANOTHER_LEADER")

                  cm.dispose()
                  return
               }

               int res = cm.createExpedition(expeditionType)
               if (res == 0) {
                  cm.sendOk("9270047_EXPEDITION_CREATED", expeditionBoss)
               } else if (res > 0) {
                  cm.sendOk("9270047_QUOTA_LIMIT")

               } else {
                  cm.sendOk("9270047_UNEXPECTED_ERROR")

               }

               cm.dispose()
            } else if (selection == 2) {
               cm.sendOk("9270047_NOT_EVERYONE_IS_UP_FOR_IT", expeditionBoss)
               cm.dispose()
            }
         } else if (status == 2) {
            if (selection == 1) {
               if (expedition == null) {
                  cm.sendOk("9270047_COULD_NOT_BE_LOADED")

                  cm.dispose()
                  return
               }
               expeditionMembers = expedition.getMemberList()
               int size = expeditionMembers.size()
               if (size == 1) {
                  cm.sendOk("9270047_NEED_MORE_MEMBERS")

                  cm.dispose()
                  return
               }
               String text = "The following members make up your expedition (Click on them to expel them):\r\n"
               text += "\r\n\t\t1." + expedition.getLeader().getName()
               for (int i = 1; i < size; i++) {
                  text += "\r\n#b#L" + (i + 1) + "#" + (i + 1) + ". " + expeditionMembers.get(i).getValue() + "#l\n"
               }
               cm.sendSimple(text)
               status = 6
            } else if (selection == 2) {
               int min = expeditionType.getMinSize()
               int size = expedition.getMemberList().size()
               if (size < min) {
                  cm.sendOk("9270047_MINIMUM_PLAYERS", min)
                  cm.dispose()
                  return
               }

               cm.sendOk("9270047_EXPEDITION_WILL_BEGIN", expeditionMap)
               status = 4
            } else if (selection == 3) {
               MessageBroadcaster.getInstance().sendMapServerNotice(player.getMap(), ServerNoticeType.LIGHT_BLUE, I18nMessage.from("EXPEDITION_ENDED_BY").with(expedition.getLeader().getName()))
               cm.endExpedition(expedition)
               cm.sendOk("9270047_EXPEDITION_ENDED")

               cm.dispose()
            }
         } else if (status == 4) {
            if (em == null) {
               cm.sendOk("9270047_EVENT_COULD_NOT_BE_INITIALIZED")

               cm.dispose()
               return
            }

            em.setProperty("leader", player.getName())
            em.setProperty("channel", player.getClient().getChannel())
            if (!em.startInstance(expedition)) {
               cm.sendOk("9270047_ANOTHER_EXPEDITION", expeditionBoss)
               cm.dispose()
               return
            }

            cm.dispose()
         } else if (status == 6) {
            if (selection > 0) {
               Map.Entry<Integer, String> banned = expeditionMembers.get(selection - 1)
               expedition.ban(banned)
               cm.sendOk("9270047_YOU_HAVE_BANNED", banned.getValue())
               cm.dispose()
            } else {
               cm.sendSimple(list)
               status = 2
            }
         }
      }
   }
}

NPC9270047 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9270047(cm: cm))
   }
   return (NPC9270047) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
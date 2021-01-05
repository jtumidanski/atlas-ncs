package com.atlas.ncs.script.npc

import com.atlas.ncs.model.PartyCharacter
import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC2133000 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   EventManager em

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

         if (status == 0) {
            em = cm.getEventManager("EllinPQ").orElseThrow()
            if (em == null) {
               cm.sendOk("2133000_PQ_ENCOUNTERED_ERROR")

               cm.dispose()
               return
            } else if (cm.isUsingOldPqNpcStyle()) {
               action((byte) 1, (byte) 0, 0)
               return
            }

            cm.sendSimple("2133000_PARTY_QUEST_INFO", em.getProperty("party"), cm.isRecvPartySearchInviteEnabled() ? "disable" : "enable")

         } else if (status == 1) {
            if (selection == 0) {
               if (cm.getParty().isEmpty()) {
                  cm.sendOk("2133000_MUST_BE_IN_PARTY")
                  cm.dispose()
               } else if (!cm.isPartyLeader()) {
                  cm.sendOk("2133000_PARTY_LEADER_MUST_START")
                  cm.dispose()
               } else {
                  PartyCharacter[] eli = em.getEligibleParty(cm.getParty().orElseThrow())
                  if (eli.size() > 0) {
                     if (!em.startInstance(cm.getParty().orElseThrow(), cm.getMapId(), 1)) {
                        cm.sendOk("2133000_ANOTHER_PARTY_HAS_ENTERED")
                     }
                  } else {
                     cm.sendOk("2133000_PARTY_REQUIREMENTS")
                  }

                  cm.dispose()
               }
            } else if (selection == 1) {
               boolean psState = cm.toggleRecvPartySearchInvite()
               cm.sendOk("2133000_PARTY_SEARCH_STATUS", (psState ? "enabled" : "disabled"))
               cm.dispose()
            } else if (selection == 2) {
               cm.sendOk("2133000_PARTY_QUEST_INFO_2")
               cm.dispose()
            } else {
               cm.sendSimple("2133000_PRIZES")
            }
         } else if (status == 2) {
            if (selection == 0) {
               if (!cm.haveItem(1032060) && cm.haveItem(4001198, 10)) {
                  cm.gainItem(1032060, (short) 1)
                  cm.gainItem(4001198, (short) -10)
                  cm.dispose()
               } else {
                  cm.sendOk("2133000_ALREADY_HAVE_EARRINGS_OR_NEED_MORE_FRAGMENTS")
                  cm.dispose()
               }
            } else if (selection == 1) {
               if (cm.haveItem(1032060) && !cm.haveItem(1032061) && cm.haveItem(4001198, 10)) {
                  cm.gainItem(1032060, (short) -1)
                  cm.gainItem(1032061, (short) 1)
                  cm.gainItem(4001198, (short) -10)
                  cm.dispose()
               } else {
                  cm.sendOk("2133000_NO_EARRINGS_OR_NEED_MORE_FRAGMENTS")
                  cm.dispose()
               }
            } else if (selection == 2) {
               if (cm.haveItem(1032061) && !cm.haveItem(1032072) && cm.haveItem(4001198, 10)) {
                  cm.gainItem(1032061, (short) -1)
                  cm.gainItem(1032072, (short) 1)
                  cm.gainItem(4001198, (short) -10)
                  cm.dispose()
               } else {
                  cm.sendOk("2133000_NO_EARRINGS_OR_NEED_MORE_FRAGMENTS_GLITTERING")
                  cm.dispose()
               }
            }
         }
      }
   }
}

NPC2133000 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2133000(cm: cm))
   }
   return (NPC2133000) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
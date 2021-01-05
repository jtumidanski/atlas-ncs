package com.atlas.ncs.script.npc

import com.atlas.ncs.model.PartyCharacter
import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC2041023 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   EventManager em = null

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
            if (!(cm.isQuestCompleted(6316) && (cm.isQuestStarted(6225) || cm.isQuestStarted(6315)))) {
               cm.sendOk("2041023_NO_REASON_TO_MEET_THANATOS")

               cm.dispose()
               return
            }

            em = cm.getEventManager("ElementalBattle").orElseThrow()
            if (em == null) {
               cm.sendOk("2041023_ENCOUNTERED_ERROR")

               cm.dispose()
               return
            } else if (cm.isUsingOldPqNpcStyle()) {
               action((byte) 1, (byte) 0, 0)
               return
            }

            cm.sendSimple("2041023_HELLO", em.getProperty("party"), (cm.isRecvPartySearchInviteEnabled() ? "disable" : "enable"))

         } else if (status == 1) {
            if (selection == 0) {
               if (cm.getParty().isEmpty()) {
                  cm.sendOk("2041023_MUST_BE_IN_PARTY")

                  cm.dispose()
               } else if (!cm.isPartyLeader()) {
                  cm.sendOk("2041023_MUST_BE_LEADER")

                  cm.dispose()
               } else {
                  PartyCharacter[] eli = em.getEligibleParty(cm.getParty().orElseThrow())
                  if (eli.size() > 0) {
                     if (!em.startInstance(cm.getParty().orElseThrow(), cm.getMapId(), 1)) {
                        cm.sendOk("2041023_ANOTHER_PARTY_CHALLENGING")

                     }
                  } else {
                     cm.sendOk("2041023_PARTY_REQUIREMENT_ISSUE")

                  }

                  cm.dispose()
               }
            } else if (selection == 1) {
               boolean psState = cm.toggleRecvPartySearchInvite()
               cm.sendOk("2041023_PARTY_SEARCH_STATUS", psState ? "enabled" : "disabled")
               cm.dispose()
            } else {
               cm.sendOk("2041023_HELLO_SHORT")

               cm.dispose()
            }
         }
      }
   }
}

NPC2041023 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2041023(cm: cm))
   }
   return (NPC2041023) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
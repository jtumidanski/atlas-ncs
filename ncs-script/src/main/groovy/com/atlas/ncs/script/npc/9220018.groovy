package com.atlas.ncs.script.npc

import com.atlas.ncs.model.PartyCharacter
import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9220018 {
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
            em = cm.getEventManager("TreasurePQ").orElseThrow()
            if (em == null) {
               cm.sendOk("9220018_ENCOUNTERED_ERROR")

               cm.dispose()
               return
            } else if (cm.isUsingOldPqNpcStyle()) {
               action((byte) 1, (byte) 0, 0)
               return
            }

            cm.sendSimple("9220018_PARTY_QUEST_INFO", em.getProperty("party"), cm.isRecvPartySearchInviteEnabled() ? "disable" : "enable")

         } else if (status == 1) {
            if (selection == 0) {
               if (cm.getParty().isEmpty()) {
                  cm.sendOk("9220018_MUST_BE_IN_PARTY")
                  cm.dispose()
               } else if (!cm.isPartyLeader()) {
                  cm.sendOk("9220018_PARTY_LEADER_MUST_START")
                  cm.dispose()
               } else {
                  PartyCharacter[] eli = em.getEligibleParty(cm.getParty().orElseThrow())
                  if (eli.size() > 0) {
                     if (!em.startInstance(cm.getParty().orElseThrow(), cm.getMapId(), 1)) {
                        cm.sendOk("9220018_ANOTHER_PARTY")
                     }
                  } else {
                     cm.sendOk("9220018_PARTY_REQUIREMENT")
                  }

                  cm.dispose()
               }
            } else if (selection == 1) {
               boolean psState = cm.toggleRecvPartySearchInvite()
               cm.sendOk("9220018_PARTY_SEARCH_STATUS", psState ? "enabled" : "disabled")
               cm.dispose()
            } else {
               cm.sendOk("9220018_PARTY_QUEST_INFO_2")
               cm.dispose()
            }
         }
      }
   }
}

NPC9220018 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9220018(cm: cm))
   }
   return (NPC9220018) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
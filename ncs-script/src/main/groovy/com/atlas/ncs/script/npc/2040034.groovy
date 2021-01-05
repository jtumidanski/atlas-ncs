package com.atlas.ncs.script.npc

import com.atlas.ncs.model.PartyCharacter
import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

/*
	NPC Name: 		Red Sign
	Map(s): 		101st Floor Eos Tower
	Description: 	
*/
class NPC2040034 {
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
            em = cm.getEventManager("LudiPQ").orElseThrow()
            if (em == null) {
               cm.sendOk("2040034_PQ_ENCOUNTERED_ERROR")
               cm.dispose()
               return
            } else if (cm.isUsingOldPqNpcStyle()) {
               action((byte) 1, (byte) 0, 0)
               return
            }

            cm.sendSimple("2040034_INFO", em.getProperty("party"), (cm.isRecvPartySearchInviteEnabled() ? "disable" : "enable"))
         } else if (status == 1) {
            if (selection == 0) {
               if (cm.getParty().isEmpty()) {
                  cm.sendOk("2040034_NEED_PARTY")
                  cm.dispose()
               } else if (!cm.isPartyLeader()) {
                  cm.sendOk("2040034_PARTY_LEADER_MUST_START")
                  cm.dispose()
               } else {
                  PartyCharacter[] eli = em.getEligibleParty(cm.getParty().orElseThrow())
                  if (eli.size() > 0) {
                     if (!em.startInstance(cm.getParty().orElseThrow(), cm.getMapId(), 1)) {
                        cm.sendOk("2040034_ANOTHER_PARTY_HAS_STARTED")
                     }
                  } else {
                     cm.sendOk("2040034_PARTY_REQUIREMENT_ISSUE")
                  }

                  cm.dispose()
               }
            } else if (selection == 1) {
               boolean psState = cm.toggleRecvPartySearchInvite()
               cm.sendOk("2040034_PARTY_SEARCH_STATUS", psState ? "enabled" : "disabled")
               cm.dispose()
            } else {
               cm.sendOk("2040034_DIMENSIONAL_SCHISM")
               cm.dispose()
            }
         }
      }
   }
}

NPC2040034 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2040034(cm: cm))
   }
   return (NPC2040034) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
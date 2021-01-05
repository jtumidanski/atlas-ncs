package com.atlas.ncs.script.npc

import com.atlas.ncs.model.PartyCharacter
import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC2095000 {
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
         if (mode == 0 && type > 0) {
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }

         if (status == 0) {
            if (cm.getMapId() != 925010400) {
               em = cm.getEventManager("DelliBattle").orElseThrow()
               if (em == null) {
                  cm.sendOk("2095000_BATTLE_ENCOUNTERED_ERROR")
                  cm.dispose()
                  return
               } else if (cm.isUsingOldPqNpcStyle()) {
                  action((byte) 1, (byte) 0, 0)
                  return
               }

               cm.sendSimple("2095000_PARTY_QUEST_INFO", em.getProperty("party"), cm.isRecvPartySearchInviteEnabled() ? "disable" : "enable")

            } else {
               cm.sendYesNo("2095000_ARE_YOU_READY")
            }
         } else if (status == 1) {
            if (cm.getMapId() != 925010400) {
               if (selection == 0) {
                  if (cm.getParty().isEmpty()) {
                     cm.sendOk("2095000_MUST_BE_IN_PARTY")
                     cm.dispose()
                  } else if (!cm.isPartyLeader()) {
                     cm.sendOk("2095000_PARTY_LEADER_MUST_START")
                     cm.dispose()
                  } else {
                     PartyCharacter[] eli = em.getEligibleParty(cm.getParty().orElseThrow())
                     if (eli.size() > 0) {
                        if (!em.startInstance(cm.getParty().orElseThrow(), cm.getMapId(), 1)) {
                           cm.sendOk("2095000_ANOTHER_PARTY_ENTERED")
                        }
                     } else {
                        cm.sendOk("2095000_PARTY_REQUIREMENTS")
                     }

                     cm.dispose()
                  }
               } else if (selection == 1) {
                  boolean psState = cm.toggleRecvPartySearchInvite()
                  cm.sendOk("2095000_PARTY_SEARCH_STATUS", (psState ? "enabled" : "disabled"))
                  cm.dispose()
               } else {
                  cm.sendOk("2095000_PARTY_QUEST_INFO_SHORT")
                  cm.dispose()
               }
            } else {
               cm.warp(120000104)
               cm.dispose()
            }
         }
      }
   }
}

NPC2095000 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2095000(cm: cm))
   }
   return (NPC2095000) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
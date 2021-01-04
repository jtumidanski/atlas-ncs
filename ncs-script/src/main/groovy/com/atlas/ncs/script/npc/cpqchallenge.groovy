package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NpcCpqChallenge {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   MapleCharacter[] party

   def start(MapleCharacter[] characters) {
      status = -1
      party = characters
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.answerCPQChallenge(false)
         cm.getChar().setChallenged(false)
         cm.dispose()
      } else {
         if (mode == 0) {
            cm.answerCPQChallenge(false)
            cm.getChar().setChallenged(false)
            cm.dispose()
            return
         }
      }
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 1) {
            status++
         } else {
            status--
         }

         if (status == 0) {
            if (cm.getParty().orElseThrow().getMembers().size() == party.size()) {
               cm.getPlayer().setChallenged(true)
               String snd = ""
               for (int i = 0; i < party.size(); i++ )
               snd += "#bName: " + party[i].getName() + " / (Level: " + party[i].getLevel() + ") / " + GameConstants.getJobName(party[i].getJob().getId()) + "#k\r\n\r\n"
               cm.sendAcceptDecline(snd + "Would you like to fight this party at the Monster Carnival?")
            } else {
               cm.answerCPQChallenge(false)
               cm.getChar().setChallenged(false)
               cm.dispose()
            }
         } else if (status == 1) {
            if (party.size() == cm.getParty().orElseThrow().getMembers().size()) {
               cm.answerCPQChallenge(true)
            } else {
               cm.answerCPQChallenge(false)
               cm.getChar().setChallenged(false)
               cm.sendOk("cpqchallenge_NOT_THE_SAME")

            }
            cm.dispose()
         }
      }
   }
}

NpcCpqChallenge getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NpcCpqChallenge(cm: cm))
   }
   return (NpcCpqChallenge) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
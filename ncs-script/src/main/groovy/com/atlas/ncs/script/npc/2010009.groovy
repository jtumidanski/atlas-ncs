package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2010009 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int choice
   String allianceName

   int allianceCost = 2000000
   int increaseCost = 1000000
   int allianceLimit = 5

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == 1) {
         status++
      } else {
         cm.dispose()
         return
      }
      if (status == 0) {
         if (cm.getGuildId() < 1 || cm.getGuildRank() != 1) {
            cm.sendNext("2010009_HELLO_NON_GUILD_LEADER")
            cm.dispose()
            return
         }

         cm.sendSimple("2010009_HELLO_GUILD_LEADER")
      } else if (status == 1) {
         choice = selection
         if (selection == 0) {
            cm.sendNext("2010009_UNION_DEFINITION")
            cm.dispose()
         } else if (selection == 1) {
            cm.sendNext("2010009_HOW_TO")
            cm.dispose()
         } else if (selection == 2) {
            if (!cm.isPartyLeader()) {
               cm.sendNext("2010009_PARTY_LEADER_MUST_TALK")
               cm.dispose()
               return
            }
            if (cm.getAllianceId() > 0) {
               cm.sendOk("2010009_GUILD_CANNOT_BE_PART_OF_ANOTHER_UNION")
               cm.dispose()
               return
            }

            cm.sendYesNo("2010009_UNION_CREATION_FEE", allianceCost)
         } else if (selection == 3) {
            if (cm.getGuildCharacter().isEmpty()) {
               cm.sendOk("2010009_MUST_OWN_A_UNION_TO_EXPAND_ONE")
               cm.dispose()
               return
            }

            int rank = cm.getAllianceRank()
            if (rank == 1) {
               cm.sendYesNo("2010009_UNION_EXPANSION_COST", increaseCost)
            } else {
               cm.sendNext("2010009_MUST_BE_UNION_LEADER_TO_EXPAND")
               cm.dispose()
            }
         } else if (selection == 4) {
            if (cm.getGuildCharacter().isEmpty()) {
               cm.sendOk("2010009_MUST_OWN_A_UNION_TO_DISBAND_ONE")
               cm.dispose()
               return
            }

            int rank = cm.getAllianceRank()
            if (rank == 1) {
               cm.sendYesNo("2010009_UNION_DISBAND_CONFIRMATION")
            } else {
               cm.sendNext("2010009_MUST_BE_UNION_LEADER_TO_DISBAND")
               cm.dispose()
            }
         }
      } else if (status == 2) {
         if (choice == 2) {
            if (cm.getMeso() < allianceCost) {
               cm.sendOk("2010009_NOT_ENOUGH_MESOS")
               cm.dispose()
               return
            }
            cm.sendGetText("2010009_UNION_NAME_INPUT")
         } else if (choice == 3) {
            if (cm.getAllianceCapacity() == allianceLimit) {
               cm.sendOk("2010009_UNION_AT_CAPACITY")
               cm.dispose()
               return
            }
            if (cm.getMeso() < increaseCost) {
               cm.sendOk("2010009_NOT_ENOUGH_MESOS")
               cm.dispose()
               return
            }

            cm.upgradeAlliance()
            cm.gainMeso(-increaseCost)
            cm.sendOk("2010009_UNION_EXPANSION_SUCCESS")
            cm.dispose()
         } else if (choice == 4) {
            if (cm.hasGuild() || cm.getAllianceId() <= 0) {
               cm.sendNext("2010009_CANNOT_DISBAND_NON_EXISTENT_UNION")
               cm.dispose()
            } else {
               cm.disbandAlliance(cm.getAllianceId())
               cm.sendOk("2010009_UNION_DISBAND_SUCCESS")
               cm.dispose()
            }
         }
      } else if (status == 3) {
         allianceName = cm.getText()
         cm.sendYesNo("2010009_UNION_NAME_CONFIRMATION", allianceName)
      } else if (status == 4) {
         if (!cm.canBeUsedAllianceName(allianceName)) {
            cm.sendNext("2010009_UNION_NAME_UNAVAILABLE") //Not real text
            status = 1
            choice = 2
         } else {
            if (cm.createAlliance(allianceName).isEmpty()) {
               cm.sendOk("2010009_PARTY_COMPOSITION_ISSUE")
            } else {
               cm.gainMeso(-allianceCost)
               cm.sendOk("2010009_UNION_CREATION_SUCCESS")
            }
            cm.dispose()
         }
      }
   }
}

NPC2010009 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2010009(cm: cm))
   }
   return (NPC2010009) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
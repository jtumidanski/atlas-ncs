package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1032001 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def action = ["1stJob": false, "2ndjob": false, "3thJobI": false, "3thJobC": false]
   int job = 210
   boolean spawnPlayerNpc = false
   int spawnPlayerNpcFee = 7000000
   int jobType = 2

   def start() {
      if ((cm.getJobId() / 100).intValue() == jobType && cm.canSpawnPlayerNpc(cm.getHallOfFameMapId(cm.getJobId()))) {
         spawnPlayerNpc = true

         String sendStr = cm.evaluateToken("1032001_WALKED_A_LONG_WAY")
         if (spawnPlayerNpcFee > 0) {
            sendStr += cm.evaluateToken("1032001_FEE", spawnPlayerNpcFee)
         }

         cm.sendYesNo(sendStr)
      } else {
         if (cm.getJobId() == 0) {
            action["1stJob"] = true
            cm.sendNext("1032001_HELLO")
         } else if (cm.getLevel() >= 30 && cm.getJobId() == 200) {
            action["2ndJob"] = true
            if (cm.haveItem(4031012)) {
               cm.sendNext("1032001_NEXT_STEP")
            } else if (cm.haveItem(4031009)) {
               cm.sendOk("1032001_GO_AND_SEE")
               cm.dispose()
            } else {
               cm.sendNext("1032001_ASTONISHING")
            }
         } else if (action["3thJobI"] || (cm.gotPartyQuestItem("JB3") && cm.getLevel() >= 70 && cm.getJobId() % 10 == 0 && (cm.getJobId() / 100).intValue() == 2 && !cm.gotPartyQuestItem("JBP"))) {
            action["3thJobI"] = true
            cm.sendNext("1032001_THERE_YOU_ARE")
         } else if (cm.gotPartyQuestItem("JBP") && !cm.haveItem(4031059)) {
            cm.sendNext("1032001_PLEASE_BRING_ME")
            cm.dispose()
         } else if (cm.haveItem(4031059) && cm.gotPartyQuestItem("JBP")) {
            action["3thJobC"] = true
            cm.sendNext("1032001_NICE_WORK")
         } else {
            cm.sendOk("1032001_CHOSEN_WISELY")
            cm.dispose()
         }
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      status++
      if (mode == -1 && selection == -1) {
         cm.dispose()
         return
      } else if (mode == 0 && type == 0) {
         status -= 2
      }

      if (status == -1) {
         start()
         return
      } else {
         if (spawnPlayerNpc) {
            if (mode > 0) {
               if (cm.getMeso() < spawnPlayerNpcFee) {
                  cm.sendOk("1032001_NOT_ENOUGH_MESOS")
                  cm.dispose()
                  return
               }

               if (cm.spawnPlayerNPC(cm.getHallOfFameMapId(cm.getJobId()))) {
                  cm.sendOk("1032001_THERE_YOU_GO")
                  cm.gainMeso(-spawnPlayerNpcFee)
               } else {
                  cm.sendOk("1032001_CURRENTLY_FULL")
               }
            }

            cm.dispose()
            return
         } else {
            if (mode != 1 || status == 7 || (action["1stJob"] && status == 4) || (cm.haveItem(4031008) && status == 2) || (action["3thJobI"] && status == 1)) {
               if (mode == 0 && status == 2 && type == 1) {
                  cm.sendOk("1032001_NO_OTHER_CHOICE")
               }
               if (!(mode == 0 && type == 0)) {
                  cm.dispose()
                  return
               }
            }
         }
      }

      if (action["1stJob"]) {
         if (status == 0) {
            if (cm.getLevel() >= 8 && cm.canGetFirstJob(jobType)) {
               cm.sendYesNo("1032001_WANNA_BE_A_MAGICIAN")
            } else {
               cm.sendOk("1032001_TRAIN_A_BIT_MORE")
               cm.dispose()
            }
         } else if (status == 1) {
            if (cm.canHold(1372043)) {
               if (cm.getJobId() == 0) {
                  cm.changeJob(200)
                  cm.gainItem(1372043, (short) 1)
                  cm.resetStats()
               }
               cm.sendNext("1032001_GIVE_YOU_SOME_OF_MY_ABILITIES")
            } else {
               cm.sendNext("1032001_MAKE_SOME_INVENTORY_ROOM")
               cm.dispose()
            }
         } else if (status == 2) {
            cm.sendNextPrev("1032001_YOU_ARE_MUCH_STRONGER_NOW")
         } else if (status == 3) {
            cm.sendNextPrev("1032001_STATS_SHOULD_SUPPORT_YOUR_SKILLS")
         } else if (status == 4) {
            cm.sendNextPrev("1032001_IF_YOU_DIE")
         } else if (status == 5) {
            cm.sendNextPrev("1032001_ALL_I_CAN_TEACH_YOU")
         } else {
            cm.dispose()
         }
      } else if (action["2ndJob"]) {
         if (status == 0) {
            if (cm.haveItem(4031012)) {
               cm.sendSimple("1032001_PATH_INFO")
            } else {
               cm.sendNext("1032001_NOT_A_DIFFICULT_TEST")
               if (!cm.isQuestStarted(100006)) {
                  cm.startQuest(100006)
               }
            }
         } else if (status == 1) {
            if (!cm.haveItem(4031012)) {
               if (cm.canHold(4031009)) {
                  if (!cm.haveItem(4031009)) {
                     cm.gainItem(4031009, (short) 1)
                  }
                  cm.sendNextPrev("1032001_PLEASE_GET_THIS_LETTER_TO")
               } else {
                  cm.sendNext("1032001_MAKE_SOME_SPACE")
                  cm.dispose()
               }
            } else {
               if (selection < 3) {
                  if (selection == 0) {
                     cm.sendNext("1032001_FIRE_POISON_INFO")
                     //f/p magician
                  } else if (selection == 1) {
                     cm.sendNext("1032001_ICE_LIGHTNING_INFO")
                     //i/l magician
                  } else {
                     cm.sendNext("1032001_CLERIC_INFO")
                     //cleric
                  }

                  status -= 2
               } else {
                  cm.sendSimple("1032001_CHOOSE_THE_JOB")
               }
            }
         } else if (status == 2) {
            if (cm.haveItem(4031009)) {
               cm.dispose()
               return
            }
            job += selection * 10
            cm.sendYesNo("1032001_CONFIRM", job == 210 ? "#bWizard (Fire / Poison)#k" : job == 220 ? "#bWizard (Ice / Lighting)#k" : "#bCleric#k")
         } else if (status == 3) {
            if (cm.haveItem(4031012)) {
               cm.gainItem(4031012, (short) -1)
            }
            cm.completeQuest(100008)
            cm.sendNext("1032001_2ND_JOB_SUCCESS", job == 210 ? "#bWizard (Fire / Poison)#k" : job == 220 ? "#bWizard (Ice / Lighting)#k" : "#bCleric#k")
            if (cm.getJobId() != job) {
               cm.changeJob(job)
            }
         } else if (status == 4) {
            cm.sendNextPrev("1032001_GIVEN_YOU_A_BOOK", job == 210 ? "#bWizard (Fire / Poison)#k" : job == 220 ? "#bWizard (Ice / Lighting)#k" : "#bCleric#k")
         } else if (status == 5) {
            cm.sendNextPrev("1032001_GIVEN_SP")
         } else if (status == 6) {
            cm.sendNextPrev("1032001_NEED_TO_BE_STRONG", job == 210 ? "Wizard (Fire / Poison)" : job == 220 ? "Wizard (Ice / Lighting)" : "Cleric")
         }
      } else if (action["3thJobI"]) {
         if (status == 0) {
            if (cm.gotPartyQuestItem("JB3")) {
               cm.removePartyQuestItem("JB3")
               cm.removePartyQuestItem("JB3")
               cm.setPartyQuestItemObtained("JBP")
            }
            cm.sendNextPrev("1032001_DEFEAT_CLONE")
         }
      } else if (action["3thJobC"]) {
         cm.removePartyQuestItem("JBP")
         cm.gainItem(4031059, (short) -1)
         cm.gainItem(4031057, (short) 1)
         cm.dispose()
      }
   }
}

NPC1032001 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1032001(cm: cm))
   }
   return (NPC1032001) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
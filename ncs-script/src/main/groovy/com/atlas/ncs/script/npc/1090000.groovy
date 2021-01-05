package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC1090000 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def action = ["1stJob": false, "2ndjob": false, "2ndjobT": false, "3thJobI": false, "3thJobC": false]
   int job = 510

   boolean spawnPlayerNpc = false
   int spawnPlayerNpcFee = 7000000
   int jobType = 5

   int advQuest = 0

   def start() {
      if (cm.isQuestStarted(6330)) {
         if (cm.getEventInstance() != null) {
            advQuest = 5
            cm.sendNext("1090000_NOT_BAD")
         } else if (cm.getQuestProgressInt(6330, 6331) == 0) {
            advQuest = 1
            cm.sendNext("1090000_READY_RIGHT")
         } else {
            advQuest = 3
            cm.teachSkill(5121003, (byte) 0, (byte) 10, -1)
            cm.forceCompleteQuest(6330)
            cm.sendNext("1090000_CONGRATULATIONS_TRANSFORMATION")
         }
      } else if (cm.isQuestStarted(6370)) {
         if (cm.getEventInstance() != null) {
            advQuest = 6
            cm.sendNext("1090000_NOT_BAD")
         } else if (cm.getQuestProgressInt(6370, 6371) == 0) {
            advQuest = 2
            cm.sendNext("1090000_READY_RIGHT")
         } else {
            advQuest = 4
            cm.teachSkill(5221006, (byte) 0, (byte) 10, -1)
            cm.forceCompleteQuest(6370)
            cm.sendNext("1090000_CONGRATULATIONS_BATTLE_SHIP")
         }
      } else if ((cm.getJobId() / 100).intValue() == jobType && cm.canSpawnPlayerNpc(cm.getHallOfFameMapId(cm.getJobId()))) {
         spawnPlayerNpc = true

         String sendStr = "You have walked a long way to reach the power, wisdom and courage you hold today, haven't you? What do you say about having right now #ra NPC on the Hall of Fame holding the current image of your character#k? Do you like it?"
         if (spawnPlayerNpcFee > 0) {
            sendStr += " I can do it for you, for the fee of #b " + cm.numberWithCommas(spawnPlayerNpcFee) + " mesos.#k"
         }

         cm.sendYesNo(sendStr)
      } else {
         if (cm.getJobId() == 0) {
            action["1stJob"] = true
            cm.sendNext("1090000_WANT_TO_BE_A_PIRATE", cm.getFirstJobStatRequirement(jobType))
         } else if (cm.getLevel() >= 30 && cm.getJobId() == 500) {
            action["2ndJob"] = true
            if (cm.isQuestCompleted(2191) || cm.isQuestCompleted(2192)) {
               cm.sendNext("1090000_ALLOW_YOU")
            } else {
               cm.sendNext("1090000_PROGRESS")
            }
         } else if (action["3thJobI"] || (cm.gotPartyQuestItem("JB3") && cm.getLevel() >= 70 && cm.getJobId() % 10 == 0 && (cm.getJobId() / 100).intValue() == 5 && !cm.gotPartyQuestItem("JBP"))) {
            action["3thJobI"] = true
            cm.sendNext("1090000_THERE_YOU_ARE")
         } else if (cm.gotPartyQuestItem("JBP") && !cm.haveItem(4031059)) {
            cm.sendNext("1090000_BRING_ME")
            cm.dispose()
         } else if (cm.haveItem(4031059) && cm.gotPartyQuestItem("JBP")) {
            action["3thJobC"] = true
            cm.sendNext("1090000_DEFEATED_CLONE")
         } else {
            cm.sendOk("1090000_CHOSEN_WISELY")
            cm.dispose()
         }
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      status++
      if (mode == -1 && selection == -1) {
         cm.dispose()
         return
      } else if (mode == 0 && type != 1) {
         status -= 2
      }

      if (status == -1) {
         start()
         return
      } else {
         if (advQuest > 0) {
            if (advQuest < 3) {
               EventManager em = cm.getEventManager(advQuest == 1 ? "4jship" : "4jsuper").orElseThrow()
               if (!em.startInstance(cm.getCharacterId())) {
                  cm.sendOk("1090000_SOMEONE_IS_ALREADY_CHALLENGING")
               }
            } else if (advQuest < 5) {
               if (advQuest == 3) {
                  cm.sendOk("1090000_SIMILAR_TO_TRANSFORMATION")
               } else {
                  cm.sendOk("1090000_DEFINITELY_DIFFERENT")
               }
            } else {
               if (advQuest < 6) {
                  cm.setQuestProgress(6330, 6331, 2)
               } else {
                  cm.setQuestProgress(6370, 6371, 2)
               }
               cm.warp(120000101)
            }

            cm.dispose()
         } else if (spawnPlayerNpc) {
            if (mode > 0) {
               if (cm.getMeso() < spawnPlayerNpcFee) {
                  cm.sendOk("1090000_NOT_ENOUGH_MESOS")
                  cm.dispose()
                  return
               }

               if (cm.spawnPlayerNPC(cm.getHallOfFameMapId(cm.getJobId()))) {
                  cm.sendOk("1090000_THERE_YOU_GO")
                  cm.gainMeso(-spawnPlayerNpcFee)
               } else {
                  cm.sendOk("1090000_FULL")
               }
            }

            cm.dispose()
            return
         } else {
            if (mode != 1 || status == 7 && type != 1 || (action["1stJob"] && status == 4) || (cm.haveItem(4031008) && status == 2) || (action["3thJobI"] && status == 1)) {
               if (mode == 0 && status == 2 && type == 1) {
                  cm.sendOk("1090000_NO_OTHER_CHOICE")
               }
               if (!(mode == 0 && type != 1)) {
                  cm.dispose()
                  return
               }
            }
         }
      }

      if (action["1stJob"]) {
         if (status == 0) {
            if (cm.getLevel() >= 10 && cm.canGetFirstJob(jobType)) {
               cm.sendYesNo("1090000_WANNA_BE_A_PIRATE")
            } else {
               cm.sendOk("1090000_TRAIN_A_BIT_MORE")
               cm.dispose()
            }
         } else if (status == 1) {
            if (cm.canHold(2070000) && cm.canHoldAll([1482000, 1492000])) {
               if (cm.getJobId() == 0) {
                  cm.changeJob(500)
                  cm.gainItem(1492000, (short) 1)
                  cm.gainItem(1482000, (short) 1)
                  cm.gainItem(2330000, (short) 1000)
                  cm.resetStats()
               }
               cm.sendNext("1090000_YOU_ARE_PART_OF_US")
            } else {
               cm.sendNext("1090000_MAKE_SOME_INVENTORY_ROOM")
               cm.dispose()
            }
         } else if (status == 2) {
            cm.sendNextPrev("1090000_MUCH_STRONGER_NOW")
         } else if (status == 3) {
            cm.sendNextPrev("1090000_CANNOT_CHANGE_YOUR_MIND")
         } else {
            cm.dispose()
         }
      } else if (action["2ndJob"]) {
         if (status == 0) {
            if (cm.isQuestCompleted(2191) || cm.isQuestCompleted(2192)) {
               cm.sendSimple("1090000_INFO")
            } else {
               cm.sendNext("1090000_GOOD_DECISION")
            }
         } else if (status == 1) {
            if (!cm.isQuestCompleted(2191) && !cm.isQuestCompleted(2192)) {
               // Pirate works differently from the other jobs. It warps you directly in.
               action["2ndJobT"] = true
               cm.sendYesNo("1090000_WOULD_YOU_LIKE_TO_TAKE_THE_TEST")
            } else {
               if (selection < 3) {
                  if (selection == 0) {    //brawler
                     cm.sendNext("1090000_BRAWLER_INFO")
                  } else if (selection == 1) {    //gunslinger
                     cm.sendNext("1090000_GUNSLINGER_INFO")
                  }

                  status -= 2
               } else {
                  cm.sendNextPrev("1090000_LONG_ROAD")
               }
            }
         } else if (status == 2) {
            if (action["2ndJobT"]) {
               int map
               if (cm.isQuestStarted(2191)) {
                  map = 108000502
               } else {
                  map = 108000501
               }
               if (cm.countCharactersInMap(map) > 0) {
                  cm.sendOk("1090000_TRAINING_MAPS_IN_USE")
                  cm.dispose()
               } else {
                  cm.warp(map, 0)
                  cm.dispose()
               }
            } else {
               if (cm.isQuestCompleted(2191) && cm.isQuestCompleted(2192)) {
                  job = (Math.random() < 0.5) ? 510 : 520
               } else if (cm.isQuestCompleted(2191)) {
                  job = 510
               } else if (cm.isQuestCompleted(2192)) {
                  job = 520
               }

               cm.sendYesNo("1090000_2ND_JOB_CONFIRMATION", job == 510 ? "#bBrawler#k" : "#bGunslinger#k")
            }
         } else if (status == 3) {
            if (cm.haveItem(4031012)) {
               cm.gainItem(4031012, (short) -1)
            }

            if (job == 510) {
               cm.sendNext("1090000_BRAWLER_SUCCESS")
            } else {
               cm.sendNext("1090000_GUNSLINGER_SUCCESS")
            }

            if (cm.getJobId() != job) {
               cm.changeJob(job)
            }
         } else if (status == 4) {
            cm.sendNextPrev("1090000_GIVEN_BOOK", job == 510 ? "brawler" : "gunslinger")
         } else if (status == 5) {
            cm.sendNextPrev("1090000_GIVEN_SP")
         } else if (status == 6) {
            cm.sendNextPrev((job == 510 ? "Brawlers" : "Gunslingers") + " need to be strong. But remember that you can't abuse that power and use it on a weakling. Please use your enormous power the right way, because... for you to use that the right way, that is much harden than just getting stronger. Please find me after you have advanced much further. I'll be waiting for you.")
         }
      } else if (action["3thJobI"]) {
         if (status == 0) {
            if (cm.gotPartyQuestItem("JB3")) {
               cm.removePartyQuestItem("JB3")
               cm.setPartyQuestItemObtained("JBP")
            }
            cm.sendNextPrev("1090000_CLONE_INFO")
         }
      } else if (action["3thJobC"]) {
         cm.removePartyQuestItem("JBP")
         cm.gainItem(4031059, (short) -1)
         cm.gainItem(4031057, (short) 1)
         cm.dispose()
      }
   }
}

NPC1090000 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1090000(cm: cm))
   }
   return (NPC1090000) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9201047 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int curMap, stage

   def start() {
      curMap = cm.getMapId()
      stage = Math.floor((curMap - 670010200) / 100).intValue() + 1

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
            if (cm.getMapId() != 670010200) {
               cm.sendYesNo("9201047_GOING_TO_LEAVE")

            } else {
               if (cm.isEventLeader()) {
                  EventInstanceManager eim = cm.getEventInstance()
                  int st = eim.getIntProperty("statusStg" + stage)

                  if (cm.haveItem(4031595, 1)) {
                     cm.gainItem(4031595, (short) -1)
                     eim.setIntProperty("statusStg" + stage, 1)
                     cm.sendOk("9201047_SPLENDID")
                  } else if (st < 1 && cm.getMapMonsterCount() == 0) {
                     eim.setIntProperty("statusStg" + stage, 1)
                     cm.toggleDrops()
                     cm.getMonster(9400518).ifPresent({ mobObj ->
                        cm.spawnMonsterOnGroundBelow(mobObj, -245, 810)
                        cm.sendOk("9201047_DEFEAT_IT")
                     })
                  } else {
                     if (st < 1) {
                        cm.sendOk("9201047_YOUR_TASK")
                     } else {
                        cm.sendOk("9201047_YOUR_TASK_SHORT")
                     }
                  }
               } else {
                  cm.sendOk("9201047_YOUR_TASK_LONG_NON_LEADER")
               }

               cm.dispose()
            }
         } else if (status == 1) {
            cm.warp(670010000, "st00")
            cm.dispose()
         }
      }
   }
}

NPC9201047 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201047(cm: cm))
   }
   return (NPC9201047) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
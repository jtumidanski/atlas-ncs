package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

/*
	NPC Name: 		Hotel Receptionist
	Map(s): 		
	Description: 	Sleepywood Hotel
*/
class NPC1061100 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int regularCost = 499
   int vipCost = 999
   int isRegular = 0

   def start() {
      cm.sendNext("1061100_WELCOME")
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1 || (mode == 0 && status == 1)) {
         cm.dispose()
      } else {
         if (mode == 0 && status == 2) {
            cm.sendNext("1061100_THINK_CAREFULLY")
            cm.dispose()
            return
         }
         status++
         if (status == 1) {
            cm.sendSimple("1061100_VIP_COST", regularCost, vipCost)
            isRegular = 1
         } else if (status == 2) {
            if (selection == 0) {
               cm.sendYesNo("1061100_REGULAR_CONFIRMATION")
            } else if (selection == 1) {
               cm.sendYesNo("1061100_VIP_CONFIRMATION")
               isRegular = 0
            }
         } else if (status == 3) {
            if (isRegular == 1) {
               if (cm.getMeso() >= regularCost) {
                  cm.warp(105040401)
                  cm.gainMeso(-regularCost)
               } else {
                  cm.sendNext("1061100_NOT_ENOUGH_MESO", regularCost)
               }
            } else {
               if (cm.getMeso() >= vipCost) {
                  cm.warp(105040402)
                  cm.gainMeso(-vipCost)
               } else {
                  cm.sendNext("1061100_NOT_ENOUGH_MESO", vipCost)
               }
            }
            cm.dispose()
         }
      }
   }
}

NPC1061100 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1061100(cm: cm))
   }
   return (NPC1061100) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }
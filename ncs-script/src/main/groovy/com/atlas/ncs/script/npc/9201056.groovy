package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9201056 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int fee = 15000

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode != 1) {
         if (mode == 0) {
            cm.sendOk("9201056_SEE_YOU_NEXT_TIME")

         }
         cm.dispose()
      } else {
         status++
         if (cm.getMapId() == 682000000) {
            if (status == 0) {
               if (selection == 0) {
                  cm.sendYesNo("9201056_RETURN_BACK", fee)

               }
            } else if (status == 1) {
               if (cm.getMeso() >= fee) {
                  cm.gainMeso(-fee)
                  cm.warp(600000000)
               } else {
                  cm.sendOk("9201056_NOT_ENOUGH_MESO")
               }
               cm.dispose()
            }
         } else {
            if (status == 0) {
               cm.sendYesNo("9201056_WOULD_YOU_LIKE_TO", fee)
            } else if (status == 1) {
               if (cm.getMeso() >= fee) {
                  cm.gainMeso(-fee)
                  cm.warp(682000000, 0)
               } else {
                  cm.sendOk("9201056_NOT_ENOUGH_MESO")
               }
               cm.dispose()
            }
         }
      }
   }
}

NPC9201056 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201056(cm: cm))
   }
   return (NPC9201056) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }